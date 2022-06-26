package gen

import (
	"dbkit/internal/common"
	"dbkit/internal/common/statement"
	"dbkit/internal/randomly"
)

func GenerateCreateIndexStmt(indexName string, table *common.Table) *statement.CreateIndexStmt {

	columns := make([]*common.Column, 0)
	if randomly.RandIntGap(1, 10) == 1 { //multiple column
		for _, val := range table.Columns {
			if randomly.RandBool() {
				columns = append(columns, val)
			}
		}
	} else { // single column
		for _, val := range table.Columns {
			if randomly.RandBool() {
				columns = append(columns, val)
				break
			}
		}
	}

	return &statement.CreateIndexStmt{
		OptionCreate:    randomly.RandIntGap(statement.CreOptFullText-1, statement.CreOptUnique),
		IndexName:       indexName,
		TypeIndex:       randomly.RandIntGap(statement.IndexBtree-1, statement.IndexHash),
		TableName:       table.Name,
		Columns:         columns,
		OptionAlgorithm: randomly.RandIntGap(statement.AlgorCopy-1, statement.AlgorInplace),
		OptionLock:      randomly.RandIntGap(statement.LockDefault, statement.LockShared),
		Where:           GenerateExpr(columns, 3),
	}
}
