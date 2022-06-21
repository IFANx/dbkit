package ast

import (
	"dbkit/internal/common"
)

type TabRefNode struct {
	Table *common.Table
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
