package gen

import (
	"dbkit/internal"
	"dbkit/internal/common/stmt"
)

func GenDeleteStmt(table *internal.Table) stmt.DeleteStmt {
	predicate := GenPredicate(table)
	return stmt.DeleteStmt{
		TableName: table.Name,
		Predicate: predicate,
	}
}
