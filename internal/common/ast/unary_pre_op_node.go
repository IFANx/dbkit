package ast

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
