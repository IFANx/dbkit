package internal

import (
	"dbkit/internal/common"
	"dbkit/internal/common/stmt"
)

type Provider interface {
	GetDBMS() common.DBMS
	ParseDataType(string) common.DataType
	GenCreateTableStmt(table *Table) stmt.CreateTableStmt
	GenCreateIndexStmt(table *Table) stmt.CreateIndexStmt
	GenInsertStmt(table *Table) stmt.InsertStmt
	GenUpdateStmt(table *Table) stmt.UpdateStmt
	GenSelectStmt(table *Table) stmt.SelectStmt
	GenDeleteStmt(table *Table) stmt.DeleteStmt
}
