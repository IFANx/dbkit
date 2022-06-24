package common

import (
	"dbkit/internal/common/dbms"
	"dbkit/internal/randomly"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Table struct {
	TestCtx       SqlExecutor
	DBMS          dbms.DBMS
	DBProvider    Provider
	Name          string
	DBName        string
	ColumnNames   []string
	Columns       map[string]*Column
	Indexes       map[string]*Index
	IndexCount    int
	HasPrimaryKey bool
	// partition & storage engine
}

func (table *Table) Build() {
	table.DropTable()
	stmt := table.DBProvider.GenCreateTableStmt(table)
	log.Infof("Create table statement: %s", stmt.String())
	err := table.TestCtx.ExecSQL(stmt.String())
	if err != nil {
		log.Infof("Fail to create table: %s", err)
	}
	table.UpdateSchema()
	for i := 0; i < randomly.RandIntGap(1, 3); i++ {
		stmt := table.DBProvider.GenCreateIndexStmt(table)
		log.Infof("Create index statement: %s", stmt.String())
		err := table.TestCtx.ExecSQL(stmt.String())
		if err != nil {
			log.Infof("Fail to create index: %s", err)
		}
	}
	for i := 0; i < randomly.RandIntGap(5, 10); i++ {
		stmt := table.DBProvider.GenInsertStmt(table)
		log.Infof("Insert statement: %s", stmt.String())
		err := table.TestCtx.ExecSQL(stmt.String())
		if err != nil {
			log.Infof("Fail to insert data: %s", err)
		}
	}
	table.UpdateSchema()
}

func (table *Table) DropTable() {
	table.TestCtx.ExecSQLIgnoreRes("DROP TABLE IF EXISTS " + table.Name)
}

func (table *Table) UpdateSchema() {
	switch table.DBMS {
	case dbms.MYSQL, dbms.MARIADB, dbms.TIDB:
		rows, err := table.TestCtx.Queryx("desc " + table.Name)
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
				Table:      nil,
				Name:       colName,
				Type:       table.DBProvider.ParseDataType(colType),
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
			"where table_schema = '%s' and table_name = '%s'", table.DBName, table.Name)
		rows, err = table.TestCtx.Queryx(sql)
		if err != nil {
			log.Warnf("Fail to get index: %s", err)
		}
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
