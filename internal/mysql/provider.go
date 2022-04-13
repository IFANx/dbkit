package mysql

import (
	"dbkit/internal"
	"dbkit/internal/common"
	"dbkit/internal/common/stmt"
	"dbkit/internal/mysql/gen"
)

type MySQLProvider struct{}

func (provider *MySQLProvider) GetDBMS() common.DBMS {
	return common.MYSQL
}

func (provider *MySQLProvider) ParseDataType(name string) common.DataType {
	return gen.ParseDataType(name)
}

func (provider *MySQLProvider) GenCreateTableStmt(table *internal.Table) stmt.CreateTableStmt {
	return gen.GenCreateTableStmt(table)
}

func (provider *MySQLProvider) GenCreateIndexStmt(table *internal.Table) stmt.CreateIndexStmt {
	return gen.GenCreateIndexStmt(table)
}

func (provider *MySQLProvider) GenInsertStmt(table *internal.Table) stmt.InsertStmt {
	return gen.GenInsertStmt(table)
}

func (provider *MySQLProvider) GenUpdateStmt(table *internal.Table) stmt.UpdateStmt {
	return gen.GenUpdateStmt(table)
}

func (provider *MySQLProvider) GenSelectStmt(table *internal.Table) stmt.SelectStmt {
	return gen.GenSelectStmt(table)
}

func (provider *MySQLProvider) GenDeleteStmt(table *internal.Table) stmt.DeleteStmt {
	return gen.GenDeleteStmt(table)
}
