package common

import (
	"dbkit/internal/common/dbms"
	"dbkit/internal/randomly"
	"fmt"

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
	table.DropTable()
	stmt := table.DB.DBProvider.GenCreateTableStmt(table)
	log.Infof("Create table statement: %s", stmt.String())
	err := table.DB.ExecSQL(stmt.String())
	if err != nil {
		log.Infof("Fail to create table: %s", err)
	}
	table.UpdateSchema()
	for i := 0; i < randomly.RandIntGap(1, 3); i++ {
		stmt := table.DB.DBProvider.GenCreateIndexStmt(table)
		log.Infof("Create index statement: %s", stmt.String())
		err := table.DB.ExecSQL(stmt.String())
		if err != nil {
			log.Infof("Fail to create index: %s", err)
		}
	}
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
	case dbms.MYSQL, dbms.MARIADB, dbms.TIDB:
		rows, err := table.DB.Queryx("desc " + table.Name)
		if err != nil {
			log.Warnf("Fail to get table structure: %s", err)
		}
		defer rows.Close()
		res := make(map[string]interface{})
		colNames := make([]string, 0)
		columns := make(map[string]*Column)
		for rows.Next() {
			_ = rows.MapScan(res)
			colName := string(res["Field"].([]byte))
			colType := string(res["Type"].([]byte))
			notNull := string(res["Null"].([]byte)) == "NO"
			primary := string(res["Key"].([]byte)) == "PRI"
			unique := string(res["Key"].([]byte)) == "UNI"
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
		sql := fmt.Sprintf("select * from information_schema.statistics "+
			"where table_schema = '%s' and table_name = '%s'", table.DB.DBName, table.Name)
		rows, err = table.DB.Queryx(sql)
		if err != nil {
			log.Warnf("Fail to get index: %s", err)
		}
		indexNames := make([]string, 0)
		indexes := make(map[string]*Index)
		for rows.Next() {
			_ = rows.MapScan(res)
			indexName := string(res["INDEX_NAME"].([]byte))
			indexCol := string(res["COLUMN_NAME"].([]byte))
			isPrimary := indexName == "PRIMARY"
			isUnique := string(res["NON_UNIQUE"].([]byte)) == "0"
			if indexes[indexName] == nil {
				indexCols := make([]string, 0)
				indexCols = append(indexCols, indexCol)
				indexes[indexName] = &Index{
					Name:        indexName,
					IndexedCols: indexCols,
					IsPrimary:   isPrimary,
					IsUnique:    isUnique,
				}
			} else {
				indexes[indexName].IndexedCols = append(indexes[indexName].IndexedCols, indexCol)
			}
		}
		table.IndexNames = indexNames
		table.Indexes = indexes
	}
}

func (table *Table) showSchema() {
	fmt.Println("============================")
	for idx, colName := range table.ColumnNames {
		col := table.Columns[colName]
		fmt.Printf("%d: %s %s\n", idx, col.Name, col.Type.Name())
	}
	fmt.Println("============================")
}
