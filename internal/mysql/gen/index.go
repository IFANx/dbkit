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

	if randomly.RandIntGap(1, 10) == 1 { //partial index
		predicate := GenPredicate(table)
		indexedColumns = append(indexedColumns, " WHERE ")
		indexedColumns = append(indexedColumns, predicate)
	}

	return stmt.CreateIndexStmt{
		IndexName:    idxName,
		TableName:    table.Name,
		Columns:      indexedColumns,
		OptionCreate: randomly.RandIntGap(stmt.UNIQUE, stmt.SPATIAL),
	}
}
