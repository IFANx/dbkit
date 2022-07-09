package ast

import "strings"

type AggregateNode struct {
	FuncName string
	Columns  []*ColRefNode
	Distinct bool
}

func (node *AggregateNode) Name() string {
	return node.FuncName
}

func (node *AggregateNode) Type() NodeType {
	return NodeTypeAggregateOp
}

func (node *AggregateNode) String() string {
	columns := make([]string, 0)
	for _, col := range node.Columns {
		columns = append(columns, col.String())
	}
	distinct := ""
	if node.Distinct {
		distinct = "DISTINCT "
	}
	return node.FuncName + "(" + distinct + strings.Join(columns, ", ") + ")"
}
