package dameng

import (
	"dbkit/internal/DAMENG/gen"
	"dbkit/internal/common"
	"dbkit/internal/common/dbms"
	"dbkit/internal/common/stmt"
)

type DAMENGProvider struct{}

func (provider *DAMENGProvider) GetDBMS() dbms.DBMS {
	return dbms.DAMENG
}

func (provider *DAMENGProvider) ParseDataType(name string) common.DataType {
	return gen.ParseDataType(name)
}

func (provider *DAMENGProvider) GenCreateTableStmt(table *common.Table) *stmt.CreateTableStmt {
	return gen.GenCreateTableStmt(table.Name)
}

func (provider *DAMENGProvider) GenCreateIndexStmt(table *common.Table) *stmt.CreateIndexStmt {
	return gen.GenCreateIndexStmt(table)
}

func (provider *DAMENGProvider) GenInsertStmt(table *common.Table) *stmt.InsertStmt {
	return gen.GenInsertStmt(table)
}

func (provider *DAMENGProvider) GenUpdateStmt(table *common.Table) *stmt.UpdateStmt {
	return gen.GenUpdateStmt(table)
}

func (provider *DAMENGProvider) GenSelectStmt(table *common.Table) *stmt.SelectStmt {
	return gen.GenSelectStmt(table)
}

func (provider *DAMENGProvider) GenDeleteStmt(table *common.Table) *stmt.DeleteStmt {
	return gen.GenDeleteStmt(table)
}
