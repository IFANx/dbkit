package gen

import (
	"dbkit/internal/common"
	"dbkit/internal/common/ast"
	"dbkit/internal/randomly"
)

func GenerateExprWithAggregate(columns []*common.Column, depthLimit int) ast.AstNode {
	generator := exprGenerator{columns: columns, depLimit: depthLimit, allowAggregate: true}
	return generator.genExpression(0)
}

func GenerateExpr(columns []*common.Column, depthLimit int) ast.AstNode {
	generator := exprGenerator{columns: columns, depLimit: depthLimit, allowAggregate: false}
	return generator.genExpression(0)
}

type exprGenerator struct {
	columns        []*common.Column
	depLimit       int
	allowAggregate bool
	// provider 提供各种getConstant、getUnaryPreOp、getBinaryOp、getFunc等方法
}

func (generator *exprGenerator) genExpression(depth int) ast.AstNode {
	if depth >= generator.depLimit {
		return generator.genLeafNode()
	}
	var nodeType ast.NodeType
	if generator.allowAggregate {
		nodeType = randomly.RandIntGap(ast.NodeTypeConst, ast.NodeTypeAggregateOp)
	} else {
		nodeType = randomly.RandIntGap(ast.NodeTypeConst, ast.NodeTypeFuncOp)
	}
	switch nodeType {
	case ast.NodeTypeColRef:
		return generator.genColumn()
	case ast.NodeTypeConst:
		return generator.genConstant()
	case ast.NodeTypeUnaryPreOp:
		return generator.genUnaryPreExpr(depth)
	case ast.NodeTypeUnaryPostOp:
		return generator.genUnaryPostExpr(depth)
	case ast.NodeTypeBinaryOp:
		return generator.genBinaryExpr(depth)
	case ast.NodeTypeBetweenOp:
		return generator.genBetweenExpr(depth)
	case ast.NodeTypeInOp:
		return generator.genInExpr(depth)
	case ast.NodeTypeCastOp:
		return generator.genCastExpr(depth)
	case ast.NodeTypeFuncOp:
		return generator.genFuncExpr(depth)
	case ast.NodeTypeAggregateOp:
		return generator.genAggregateExpr(depth)
	}
	panic("未知的节点类型")
}

func (generator *exprGenerator) genLeafNode() ast.AstNode {
	if randomly.RandBool() && len(generator.columns) > 0 {
		return generator.genColumn()
	} else {
		return generator.genConstant()
	}
}

func (generator *exprGenerator) genColumn() ast.AstNode {
	selColumn := RandPickOneCol(generator.columns)
	return &ast.ColRefNode{Column: selColumn}
}

func (generator *exprGenerator) genConstant() ast.AstNode {
	dataType := GetRandomMySQLDataType()
	return &ast.ConstNode{
		ConstType: GetRandomMySQLDataType(),
		Value:     dataType.GenRandomVal(),
	}
}

func (generator *exprGenerator) genUnaryPreExpr(depth int) ast.AstNode {
	subExpr := generator.genExpression(depth + 1)
	op := randomly.RandPickOneStr([]string{"+", "-", "NOT", "!"})
	return &ast.UnaryPreOpNode{
		OpName:  op,
		Operand: subExpr,
	}
}

func (generator *exprGenerator) genUnaryPostExpr(depth int) ast.AstNode {
	subExpr := generator.genExpression(depth + 1)
	op := randomly.RandPickOneStr([]string{"IS NULL", "IS TRUE", "IS FALSE", "!"})
	return &ast.UnaryPostOpNode{
		OpName:  op,
		Operand: subExpr,
	}
}

func (generator *exprGenerator) genBinaryExpr(depth int) ast.AstNode {
	leftExpr := generator.genExpression(depth + 1)
	rightExpr := generator.genExpression(depth + 1)
	op := randomly.RandPickOneStr([]string{"AND", "OR", "XOR", "=", "!=", "<", "<=", ">", ">=", "LIKE"})
	return &ast.BinaryOpNode{
		OpName: op,
		Left:   leftExpr,
		Right:  rightExpr,
	}
}

func (generator *exprGenerator) genBetweenExpr(depth int) ast.AstNode {
	expr := generator.genExpression(depth + 1)
	leftExpr := generator.genExpression(depth + 1)
	rightExpr := generator.genExpression(depth + 1)
	return &ast.BetweenOpNode{
		Expr:  expr,
		Left:  leftExpr,
		Right: rightExpr,
	}
}

func (generator *exprGenerator) genInExpr(depth int) ast.AstNode {
	count := randomly.RandIntGap(1, 5)
	exprList := make([]ast.AstNode, count)
	for i := 0; i < count; i++ {
		exprList[i] = generator.genExpression(depth + 1)
	}
	expr := generator.genBetweenExpr(depth + 1)
	return &ast.InOpNode{
		Expr:       expr,
		InExprList: exprList,
	}
}

func (generator *exprGenerator) genCastExpr(depth int) ast.AstNode {
	expr := generator.genExpression(depth + 1)
	return &ast.CastOpNode{
		Expr:       expr,
		TargetType: GetRandomMySQLDataType(),
	}
}

func (generator *exprGenerator) genFuncExpr(depth int) ast.AstNode {
	myFunc := GetRandomMySQLFunc()
	exprList := make([]ast.AstNode, myFunc.argCnt)
	for i := 0; i < myFunc.argCnt; i++ {
		exprList[i] = generator.genExpression(depth + 1)
	}
	return &ast.FuncOpNode{
		FuncName: myFunc.name,
		ArgCount: myFunc.argCnt,
		ExprList: exprList,
	}
}

func (generator *exprGenerator) genAggregateExpr(depth int) ast.AstNode {
	myAggregate := GetRandomMySQLAggregate()
	col := generator.genColumn().(*ast.ColRefNode)
	return &ast.AggregateNode{
		FuncName: myAggregate.name,
		Column:   col,
	}
}

func RandPickOneCol(candidates []*common.Column) *common.Column {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	return candidates[randomly.RandIntGap(0, len(candidates))]
}
