package gen

import (
	"dbkit/internal"
	"dbkit/internal/common"
	"dbkit/internal/common/stmt"
	"dbkit/internal/randomly"
)

func GenCreateIndexStmt(idxName string, table common.Table) stmt.CreateIndexStmt {
	candidateColumns := randomly.RandomPickNotEmptyStr(table.ColumnNames)
	indexedColumns := make([]string, 0)
	for _, colName := range candidateColumns {
		column := table.Columns[colName]
		if column.Type.IsNumeric() {
			indexedColumns = append(indexedColumns, colName)
		} else if column.Type.IsString() {
			switch table.DBMS {
			case internal.MYSQL:
			case internal.MARIADB:
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
