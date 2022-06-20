package gen

import (
	"dbkit/internal"
	"dbkit/internal/common/ast"
	"dbkit/internal/randomly"
)

func GenerateExpr(columns []*internal.Column, depthLimit int) ast.AstNode {
	generator := exprGenerator{columns: columns, depLimit: depthLimit}
	return generator.genExpression(0)
}

type exprGenerator struct {
	columns  []*internal.Column
	depLimit int
	// provider 提供各种getConstant、getUnaryPreOp、getBinaryOp、getFunc等方法
}

func (generator *exprGenerator) genExpression(depth int) ast.AstNode {
	if depth >= generator.depLimit {
		return generator.genLeafNode()
	}
	nodeType := randomly.RandIntGap(ast.NodeTypeConst, ast.NodeTypeFuncOp)
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
	selColumn := randomly.RandPickOneCol(generator.columns)
	return &ast.ColRefNode{Column: selColumn}
}

func (generator *exprGenerator) genConstant() ast.AstNode {
	return nil
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
	return &ast.CastOpNode{
		Expr:       nil,
		TargetType: ast.DataType{},
	}
}

func (generator *exprGenerator) genFuncExpr(depth int) ast.AstNode {
	return &ast.FuncOpNode{
		FuncName: "",
		ArgCount: 0,
		ExprList: nil,
	}
}