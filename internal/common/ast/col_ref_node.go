package ast

import (
	"dbkit/internal/common"
)

type ColRefNode struct {
	Column *common.Column
}

func (node *ColRefNode) Name() string {
	return "ColumnRef"
}

func (node *ColRefNode) Type() NodeType {
	return NodeTypeColRef
}

func (node *ColRefNode) String() string {
	return node.Column.Table.Name + "." + node.Column.Name
}
