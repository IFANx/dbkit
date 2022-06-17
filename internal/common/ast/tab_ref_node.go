package ast

import "dbkit/internal"

type TabRefNode struct {
	Table *internal.Table
}

func (node *TabRefNode) Name() string {
	return "TableRef"
}

func (node *TabRefNode) Type() NodeType {
	return NodeTypeTabRef
}

func (node *TabRefNode) String() string {
	return node.Table.Name
}
