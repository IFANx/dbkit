package gen

import (
	"dbkit/internal"
	"dbkit/internal/common/stmt"
	"dbkit/internal/randomly"
)

func GenUpdateStmt(table *internal.Table) stmt.UpdateStmt {
	predicate := GenPredicate(table)
	updatedColumns := randomly.RandomPickNotEmptyStr(table.ColumnNames)
	colValMap := make(map[string]string)
	for _, colName := range updatedColumns {
		colType := table.Columns[colName].Type
		colValMap[colName] = colType.GenRandomVal()
	}
	return stmt.UpdateStmt{
		TableName:   table.Name,
		ColValPairs: colValMap,
		Predicate:   predicate,
	}
}
