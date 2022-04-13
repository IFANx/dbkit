package gen

import (
	"dbkit/internal"
	"dbkit/internal/common"
	"dbkit/internal/common/stmt"
	"dbkit/internal/randomly"
	"fmt"
)

func GenCreateIndexStmt(table *internal.Table) stmt.CreateIndexStmt {
	table.IndexCount++
	idxName := fmt.Sprintf("i%d", table.IndexCount)
	candidateColumns := randomly.RandomPickNotEmptyStr(table.ColumnNames)
	indexedColumns := make([]string, 0)
	for _, colName := range candidateColumns {
		column := table.Columns[colName]
		if column.Type.IsNumeric() {
			indexedColumns = append(indexedColumns, colName)
		} else if column.Type.IsString() {
			switch table.DBMS {
			case common.MYSQL:
			case common.MARIADB:
				indexedColumns = append(indexedColumns, colName+"(5)")
			default:
				indexedColumns = append(indexedColumns, colName)
			}
		}
	}

	return stmt.CreateIndexStmt{
		IndexName: idxName,
		TableName: table.Name,
		Columns:   indexedColumns,
		Unique:    randomly.RandBool(),
	}
}
