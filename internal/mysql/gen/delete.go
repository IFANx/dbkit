package gen

import (
	"dbkit/internal/common"
	"dbkit/internal/common/stmt"
)

func GenDeleteStmt(table *common.Table) stmt.DeleteStmt {
	predicate := GenPredicate(table)
	return stmt.DeleteStmt{
		TableName: table.Name,
		Predicate: predicate,
	}
}
