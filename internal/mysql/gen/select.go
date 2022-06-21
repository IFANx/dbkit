package gen

import (
	"dbkit/internal/common"
	"dbkit/internal/common/stmt"
	"dbkit/internal/randomly"
)

func GenSelectStmt(table *common.Table) stmt.SelectStmt {
	predicate := GenPredicate(table)
	selectedColumns := randomly.RandPickNotEmptyStr(table.ColumnNames)
	postFix := randomly.RandIntGap(0, 5)
	return stmt.SelectStmt{
		TableName: table.Name,
		Targets:   selectedColumns,
		Predicate: predicate,
		ForShare:  postFix == 0,
		ForUpdate: postFix == 1,
	}
}
