package ast

import "strings"

type FuncOpNode struct {
	FuncName string
	ArgCount int
	ExprList []AstNode
}

func (node *FuncOpNode) Name() string {
	return node.FuncName
}

func (node *FuncOpNode) Type() NodeType {
	return NodeTypeFuncOp
}

func (node *FuncOpNode) String() string {
	exprList := make([]string, 0)
	for _, expr := range node.ExprList {
		exprList = append(exprList, expr.String())
	}
	return node.FuncName + "(" + strings.Join(exprList, ", ") + ")"
}
