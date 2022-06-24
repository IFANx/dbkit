package gen

import (
	"dbkit/internal/common"
	"dbkit/internal/common/dbms"
	"dbkit/internal/common/stmt"
	"dbkit/internal/randomly"
	"fmt"
)

func GenCreateIndexStmt(table *common.Table) stmt.CreateIndexStmt {
	table.IndexCount++
	idxName := fmt.Sprintf("i%d", table.IndexCount)
	candidateColumns := randomly.RandPickNotEmptyStr(table.ColumnNames)
	indexedColumns := make([]string, 0)
	for _, colName := range candidateColumns {
		column := table.Columns[colName]
		if column.Type.IsNumeric() {
			indexedColumns = append(indexedColumns, colName)
		} else if column.Type.IsString() {
			switch table.DBMS {
			case dbms.MYSQL:
			case dbms.MARIADB:
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
		IndexName: idxName,
		TableName: table.Name,
		Columns:   indexedColumns,
		Unique:    randomly.RandBool(),
	}
}
