package ast

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
