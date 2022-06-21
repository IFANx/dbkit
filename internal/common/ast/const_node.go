package ast

import "dbkit/internal/common"

type ConstNode struct {
	ConstType common.DataType
	Value     string
}

func (node *ConstNode) Name() string {
	return "Constant"
}

func (node *ConstNode) Type() NodeType {
	return NodeTypeConst
}

func (node *ConstNode) String() string {
	return node.Value
}
