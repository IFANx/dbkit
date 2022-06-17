package ast

import "strings"

type InOpNode struct {
	Expr       AstNode
	InExprList []AstNode
}

func (node *InOpNode) Name() string {
	return "InOp"
}

func (node *InOpNode) Type() NodeType {
	return NodeTypeInOp
}

func (node *InOpNode) String() string {
	exprList := make([]string, 0)
	for _, expr := range node.InExprList {
		exprList = append(exprList, expr.String())
	}
	return "(" + node.Expr.String() + ") IN (" + strings.Join(exprList, ", ") + ")"
}
