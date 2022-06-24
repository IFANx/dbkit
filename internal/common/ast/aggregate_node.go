package ast

type AggregateNode struct {
	FuncName string
	Column   *ColRefNode
}

func (node *AggregateNode) Name() string {
	return node.FuncName
}

func (node *AggregateNode) Type() NodeType {
	return NodeTypeAggregateOp
}

func (node *AggregateNode) String() string {
	return node.FuncName + "(" + node.Column.Name() + ")"
}
