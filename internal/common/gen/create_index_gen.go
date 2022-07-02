package gen

import (
	"dbkit/internal/common"
	"dbkit/internal/common/ast"
	"dbkit/internal/common/statement"
	"dbkit/internal/randomly"
)

func GenerateCreateIndexStmt(indexName string, table *common.Table) *statement.CreateIndexStmt {

	columns := make([]*common.Column, 0)
	for _, val := range table.Columns {
		if randomly.RandBool() {
			columns = append(columns, val)
			break
		}
	}

	var keyPart ast.AstNode
	if randomly.RandBool() {
		keyPart = GenerateExpr(columns, 3)
	}

	var where ast.AstNode
	if randomly.RandBool() {
		where = GenerateExpr(columns, 3)
	}

	return &statement.CreateIndexStmt{
		OptionCreate:    randomly.RandIntGap(statement.CreOptFullText-1, statement.CreOptUnique),
		IndexName:       indexName,
		TypeIndex:       randomly.RandIntGap(statement.IndexBtree-1, statement.IndexHash),
		TableName:       table.Name,
		Columns:         columns,
		KeyPart:         keyPart,
		OptionAlgorithm: randomly.RandIntGap(statement.AlgorCopy-1, statement.AlgorInplace),
		OptionLock:      randomly.RandIntGap(statement.LockDefault-1, statement.LockShared),
		Where:           where,
	}
}
