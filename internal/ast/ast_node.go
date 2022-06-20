package ast

import (
	"dbkit/internal"
	"strings"
)

// DataType TODO
type DataType interface {
	Name() string
	DBMS() string
	String() string
}

type NodeType = int

const (
	NodeTypeConst = iota
	NodeTypeTabRef
	NodeTypeColRef
	NodeTypeUnaryPreOp
	NodeTypeUnaryPostOp
	NodeTypeBinaryOp
	NodeTypeBetweenOp
	NodeTypeInOp
	NodeTypeCastOp
	NodeTypeFuncOp
	NodeTypeJoin
	NodeTypePartition
)

type AstNode interface {
	Type() NodeType
	Name() string
	String() string
}

type ConstNode struct {
	ConstType DataType
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

type ColRefNode struct {
	Column *internal.Column
}

func (node *ColRefNode) Name() string {
	return "ColumnRef"
}

func (node *ColRefNode) Type() NodeType {
	return NodeTypeColRef
}

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

func (node *ColRefNode) String() string {
	return node.Column.Table.Name + "." + node.Column.Name
}

type UnaryPreOpNode struct {
	OpName  string
	Operand AstNode
}

func (node *UnaryPreOpNode) Name() string {
	return node.OpName
}

func (node *UnaryPreOpNode) Type() NodeType {
	return NodeTypeUnaryPreOp
}

func (node *UnaryPreOpNode) String() string {
	return node.OpName + "(" + node.Operand.String() + ")"
}

type UnaryPostOpNode struct {
	OpName  string
	Operand AstNode
}

func (node *UnaryPostOpNode) Name() string {
	return node.OpName
}

func (node *UnaryPostOpNode) Type() NodeType {
	return NodeTypeUnaryPostOp
}

func (node *UnaryPostOpNode) String() string {
	return node.OpName + "(" + node.Operand.String() + ")"
}

type BinaryOpNode struct {
	OpName string
	Left   AstNode
	Right  AstNode
}

func (node *BinaryOpNode) Name() string {
	return node.OpName
}

func (node *BinaryOpNode) Type() NodeType {
	return NodeTypeBinaryOp
}

func (node *BinaryOpNode) String() string {
	return "(" + node.Left.String() + ")" + node.OpName + "(" + node.Right.String() + ")"
}

type BetweenOpNode struct {
	Expr  AstNode
	Left  AstNode
	Right AstNode
}

func (node *BetweenOpNode) Name() string {
	return "BetweenAndOp"
}

func (node *BetweenOpNode) Type() NodeType {
	return NodeTypeBetweenOp
}

func (node *BetweenOpNode) String() string {
	return "(" + node.Expr.String() + ") BETWEEN (" + node.Left.String() + ") AND (" + node.Right.String() + ")"
}

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

type CastOpNode struct {
	Expr       AstNode
	TargetType DataType
}

func (node *CastOpNode) Name() string {
	return "CastOp"
}

func (node *CastOpNode) Type() NodeType {
	return NodeTypeCastOp
}

func (node *CastOpNode) String() string {
	return "CAST((" + node.Expr.String() + ") AS " + node.TargetType.String() + ")"
}

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
