package gen

import (
	"dbkit/internal/common"
	"dbkit/internal/common/statement"
	"dbkit/internal/randomly"
)

func GenerateCreateIndexStmt(indexName string, table *common.Table) *statement.CreateIndexStmt {

	return &statement.CreateIndexStmt{
		OptionCreate:    randomly.RandIntGap(statement.CreOptFullText-1, statement.CreOptUnique),
		IndexName:       indexName,
		TypeIndex:       randomly.RandIntGap(statement.IndexBtree-1, statement.IndexHash),
		TableName:       table.Name,
		ColumnName:      randomly.RandPickOneStr(table.ColumnNames),
		OptionAlgorithm: randomly.RandIntGap(statement.AlgorCopy-1, statement.AlgorInplace),
		OptionLock:      randomly.RandIntGap(statement.LockDefault, statement.LockShared),
	}
}
