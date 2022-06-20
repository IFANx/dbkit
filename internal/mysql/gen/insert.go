package gen

import (
	"dbkit/internal"
	"dbkit/internal/common/stmt"
	"dbkit/internal/randomly"
	"dbkit/internal/util"
)

func GenInsertStmt(table *internal.Table) stmt.InsertStmt {
	insertedColumns := randomly.RandPickNotEmptyStr(table.ColumnNames)
	for colName, column := range table.Columns {
		if (column.NotNull || column.Primary) &&
			!util.CheckStrExists(colName, insertedColumns) {
			insertedColumns = append(insertedColumns, colName)
		}
	}
	insertedValues := make([]string, 0)
	for _, colName := range insertedColumns {
		colType := table.Columns[colName].Type
		insertedValues = append(insertedValues, colType.GenRandomVal())
	}
	return stmt.InsertStmt{
		TableName: table.Name,
		Columns:   insertedColumns,
		Values:    insertedValues,
	}
}
