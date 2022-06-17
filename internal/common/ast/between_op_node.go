package ast

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
