package internal

import (
	"dbkit/internal/common"
	"dbkit/internal/randomly"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Table struct {
	TestCtx       *TestContext
	DBMS          common.DBMS
	DBProvider    Provider
	Name          string
	DBName        string
	ColumnNames   []string
	Columns       map[string]*Column
	Indexes       map[string]*common.Index
	IndexCount    int
	HasPrimaryKey bool
}

func (table *Table) Build() {
	table.DropTable()
	stmt := table.DBProvider.GenCreateTableStmt(table)
	log.Infof("生成数据表创建语句：%s", stmt.String())
	err := table.TestCtx.ExecSQL(stmt.String())
	if err != nil {
		log.Infof("创建数据表失败：%s", err)
	}
	table.UpdateSchema()
	for i := 0; i < randomly.RandIntGap(1, 3); i++ {
		stmt := table.DBProvider.GenCreateIndexStmt(table)
		log.Infof("生成索引创建语句：%s", stmt.String())
		err := table.TestCtx.ExecSQL(stmt.String())
		if err != nil {
			log.Infof("创建索引失败：%s", err)
		}
	}
	for i := 0; i < randomly.RandIntGap(5, 10); i++ {
		stmt := table.DBProvider.GenInsertStmt(table)
		log.Infof("生成数据插入语句：%s", stmt.String())
		err := table.TestCtx.ExecSQL(stmt.String())
		if err != nil {
			log.Infof("插入记录失败：%s", err)
		}
	}
	table.UpdateSchema()
}

func (table *Table) DropTable() {
	table.TestCtx.ExecSQLIgnoreRes("DROP TABLE IF EXISTS " + table.Name)
}

func (table *Table) UpdateSchema() {
	switch table.DBMS {
	case common.MYSQL, common.MARIADB, common.TIDB:
		rows, err := table.TestCtx.Queryx("desc " + table.Name)
		if err != nil {
			log.Warnf("获取表格结构信息失败")
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
			log.Warnf("获取表格索引信息失败：[%s]", err)
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
