package ast

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
	return "(" + node.Operand.String() + ")" + node.OpName
}
