package mysql

import (
	"dbkit/internal/common"
	"dbkit/internal/common/dbms"
	"dbkit/internal/common/stmt"
	"dbkit/internal/mysql/gen"
)

type MySQLProvider struct{}

func (provider *MySQLProvider) GetDBMS() dbms.DBMS {
	return dbms.MYSQL
}

func (provider *MySQLProvider) ParseDataType(name string) common.DataType {
	return gen.ParseDataType(name)
}

func (provider *MySQLProvider) GenCreateTableStmt(table *common.Table) stmt.CreateTableStmt {
	return gen.GenCreateTableStmt(table)
}

func (provider *MySQLProvider) GenCreateIndexStmt(table *common.Table) stmt.CreateIndexStmt {
	return gen.GenCreateIndexStmt(table)
}

func (provider *MySQLProvider) GenInsertStmt(table *common.Table) stmt.InsertStmt {
	return gen.GenInsertStmt(table)
}

func (provider *MySQLProvider) GenUpdateStmt(table *common.Table) stmt.UpdateStmt {
	return gen.GenUpdateStmt(table)
}

func (provider *MySQLProvider) GenSelectStmt(table *common.Table) stmt.SelectStmt {
	return gen.GenSelectStmt(table)
}

func (provider *MySQLProvider) GenDeleteStmt(table *common.Table) stmt.DeleteStmt {
	return gen.GenDeleteStmt(table)
}
