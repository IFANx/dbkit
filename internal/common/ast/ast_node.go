package ast

type DataType struct {
	Name string
}

func (dt DataType) String() string {
	return dt.Name
}

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
	NodeTypeJoin
	NodeTypePartition
)

type AstNode interface {
	Type() NodeType
	Name() string
	String() string
}
