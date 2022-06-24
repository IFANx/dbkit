package ast

type NodeType = int

const (
	NodeTypeTabRef = iota
	NodeTypeConst
	NodeTypeColRef
	NodeTypeUnaryPreOp
	NodeTypeUnaryPostOp
	NodeTypeBinaryOp
	NodeTypeBetweenOp
	NodeTypeInOp
	NodeTypeCastOp
	NodeTypeFuncOp
	NodeTypeAggregateOp
	NodeTypeJoin
	NodeTypePartition
)

type AstNode interface {
	Type() NodeType
	Name() string
	String() string
}
