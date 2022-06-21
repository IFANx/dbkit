package common

import (
	"dbkit/internal/common/stmt"
)

type Provider interface {
	GetDBMS() DBMS
	ParseDataType(string) DataType
	GenCreateTableStmt(table *Table) stmt.CreateTableStmt
	GenCreateIndexStmt(table *Table) stmt.CreateIndexStmt
	GenInsertStmt(table *Table) stmt.InsertStmt
	GenUpdateStmt(table *Table) stmt.UpdateStmt
	GenSelectStmt(table *Table) stmt.SelectStmt
	GenDeleteStmt(table *Table) stmt.DeleteStmt
}
