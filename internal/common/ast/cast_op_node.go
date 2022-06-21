package ast

import "dbkit/internal/common"

type CastOpNode struct {
	Expr       AstNode
	TargetType common.DataType
}

func (node *CastOpNode) Name() string {
	return "CastOp"
}

func (node *CastOpNode) Type() NodeType {
	return NodeTypeCastOp
}

func (node *CastOpNode) String() string {
	return "CAST((" + node.Expr.String() + ") AS " + node.TargetType.Name() + ")"
}
