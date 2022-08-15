package common

import (
	"dbkit/internal/common/dbms"
	"dbkit/internal/randomly"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Table struct {
	DB            *Database
	Name          string
	ColumnNames   []string
	Columns       map[string]*Column
	IndexNames    []string
	Indexes       map[string]*Index
	IndexCount    int
	HasPrimaryKey bool
	// partition & storage engine
}

func (table *Table) Build() {
	table.clean()
	stmt := table.DB.DBProvider.GenCreateTableStmt(table)
	log.Infof("Create table statement: %s", stmt.String())
	err := table.DB.ExecSQL(stmt.String())
	if err != nil {
		log.Infof("Fail to create table: %s", err)
	}
	table.UpdateSchema()
	//for i := 0; i < randomly.RandIntGap(1, 3); i++ {
	//	stmt := table.DB.DBProvider.GenCreateIndexStmt(table)
	//	if stmt.IndexName != "" {
	//		log.Infof("Create index statement: %s", stmt.String())
	//		err := table.DB.ExecSQL(stmt.String())
	//		if err != nil {
	//			log.Infof("Fail to create index: %s", err)
	//		} else {
	//			table.IndexCount++
	//		}
	//	}
	//}
	for i := 0; i < randomly.RandIntGap(5, 10); i++ {
		stmt := table.DB.DBProvider.GenInsertStmt(table)
		log.Infof("Insert statement: %s", stmt.String())
		err := table.DB.ExecSQL(stmt.String())
		if err != nil {
			log.Infof("Fail to insert data: %s", err)
		}
	}
	table.UpdateSchema()
}

func (table *Table) DropTable() {
	table.DB.ExecSQLIgnoreError("DROP TABLE IF EXISTS " + table.Name)
}

func (table *Table) UpdateSchema() {
	switch table.DB.DBMS {
	case dbms.MYSQL, dbms.MARIADB, dbms.TIDB, dbms.DAMENG:
		//mysql与dm8的查询数据表列信息的语句不同
		//rows, err := table.DB.Queryx("desc " + table.Name)
		sql := fmt.Sprintf("select * from DBA_TAB_COLUMNS where OWNER='%s' AND TABLE_NAME='%s'", strings.ToUpper(table.DB.DBName), strings.ToUpper(table.Name))
		rows, err := table.DB.Queryx(sql)
		if err != nil {
			log.Warnf("Fail to get table structure: %s", err)
		}
		//DM8查询列的constraint信息
		//sql1 := fmt.Sprintf("SELECT * from DBA_CONSTRAINTS a, ALL_CONS_COLUMNS b where a.CONSTRAINT_NAME=b.CONSTRAINT_NAME and CONSTRAINT_TYPE in ('R','P','U') and a.OWNER='%s' and a.TABLE_NAME='%s'", table.DB.DBName, table.Name)
		//rows1, err := table.DB.Queryx(sql1)
		//if err != nil {
		//	log.Warnf("Fail to get table structure: %s", err)
		//}
		defer rows.Close()
		//defer rows1.Close()
		res := make(map[string]interface{})
		colNames := make([]string, 0)
		columns := make(map[string]*Column)
		for rows.Next() {
			_ = rows.MapScan(res)
			//mysql的列信息
			//colName := string(res["Field"].([]byte))
			//colType := string(res["Type"].([]byte))
			//notNull := string(res["Null"].([]byte)) == "NO"
			//primary := string(res["Key"].([]byte)) == "PRI"
			//unique := string(res["Key"].([]byte)) == "UNI"

			//DM8的列信息
			colName := fmt.Sprintf("%s", res["COLUMN_NAME"])
			colType := strings.ToLower(fmt.Sprintf("%s", res["DATA_TYPE"]))
			notNullval := fmt.Sprintf("%s", res["NULLABLE"])
			notNull := notNullval == "N"
			//colName := string(res["COLUMN_NAME"].([]byte))
			//colType := string(res["DATA_TYPE"].([]byte))
			//notNull := string(res["NULLABLE"].([]byte)) == "N"
			primary := false
			unique := false

			columns[colName] = &Column{
				Table:      table,
				Name:       colName,
				Type:       table.DB.DBProvider.ParseDataType(colType),
				NotNull:    notNull,
				Unique:     unique,
				Primary:    primary,
				Length:     0,
				ValueCache: make([]string, 0),
			}
			colNames = append(colNames, colName)
		}
		table.Columns = columns
		table.ColumnNames = colNames
		table.showSchema()
		//sql = fmt.Sprintf("select * from information_schema.statistics "+
		//	"where table_schema = '%s' and table_name = '%s'", strings.ToUpper(table.DB.DBName), strings.ToUpper(table.Name))
		//rows, err = table.DB.Queryx(sql)
		//if err != nil {
		//	log.Warnf("Fail to get index: %s", err)
		//}
		//indexNames := make([]string, 0)
		//indexes := make(map[string]*Index)
		//for rows.Next() {
		//	_ = rows.MapScan(res)
		//	indexName := string(res["INDEX_NAME"].([]byte))
		//	indexCol := string(res["COLUMN_NAME"].([]byte))
		//	isPrimary := indexName == "PRIMARY"
		//	isUnique := string(res["NON_UNIQUE"].([]byte)) == "0"
		//	if indexes[indexName] == nil {
		//		indexNames = append(indexNames, indexName)
		//		indexCols := make([]string, 0)
		//		indexCols = append(indexCols, indexCol)
		//		indexes[indexName] = &Index{
		//			Name:        indexName,
		//			IndexedCols: indexCols,
		//			IsPrimary:   isPrimary,
		//			IsUnique:    isUnique,
		//		}
		//	} else {
		//		indexes[indexName].IndexedCols = append(indexes[indexName].IndexedCols, indexCol)
		//	}
		//}
		//table.IndexNames = indexNames
		//table.Indexes = indexes
		//table.showSchema()
	}
}

func (table *Table) showSchema() {
	fmt.Println("============================")
	for i, colName := range table.ColumnNames {
		col := table.Columns[colName]
		fmt.Printf("%d: %s %s\n", i, col.Name, col.Type.Name())
	}
	for i, idxName := range table.IndexNames {
		idx := table.Indexes[idxName]
		fmt.Printf("%d: %s %v\n", i, idx.Name, idx.IndexedCols)
	}
	fmt.Println("============================")
}

func (table *Table) clean() {
	table.DropTable()
	table.Columns = make(map[string]*Column)
	table.ColumnNames = make([]string, 0)
	table.Indexes = make(map[string]*Index)
	table.IndexNames = make([]string, 0)
	table.IndexCount = 0
	table.HasPrimaryKey = false
}
