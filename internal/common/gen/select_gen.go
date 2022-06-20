package gen

import (
	"dbkit/internal"
	"dbkit/internal/ast"
	"dbkit/internal/randomly"
)

func GenerateSelectStmt(tables []*internal.Table) *ast.SelectStmt {
	selOptList := make([]ast.SelectOption, 0)
	selOptIdx := randomly.RandIntGap(-1, ast.SelOptSqlCalcFoundRows+1)
	if selOptIdx > -1 {
		selOptList = append(selOptList, selOptIdx)
	}
	neededTables := make([]*internal.Table, 0)
	if randomly.RandBool() && len(tables) > 1 {
		neededTables = randomly.RandPickNotEmptyTab(tables)
	} else {
		neededTables = append(neededTables, tables[0])
	}
	neededColumns := make([]*internal.Column, 0)
	for _, tab := range neededTables {
		for _, col := range tab.Columns {
			neededColumns = append(neededColumns, col)
		}
	}
	selExprList := make([]ast.AstNode, 0)
	for i := 0; i < randomly.RandIntGap(1, 5); i++ {
		selExprList = append(selExprList, GenerateExpr(neededColumns, 3))
	}
	var joinAst, joinOnAst ast.AstNode
	if len(neededTables) > 1 && randomly.RandBool() {
		joinAst = GenerateJoinAst(neededTables)
		joinOnAst = GenerateExpr(neededColumns, 3)
	}

	var orderByExpr ast.AstNode
	var orderByOpt ast.OrderOption
	if randomly.RandBool() {
		orderByExpr = GenerateExpr(neededColumns, 3)
		orderByOpt = randomly.RandIntGap(0, 2)
	}
	var forOpt ast.ForOption
	forOpt = randomly.RandIntGap(ast.ForOptShare-1, ast.ForOptUpdate+1)
	return &ast.SelectStmt{
		Options:    selOptList,
		SelectExpr: selExprList,
		Tables:     neededTables,
		Join:       joinAst,
		JoinOn:     joinOnAst,
		Partitions: nil, // 需要表结构信息
		Where:      GenerateExpr(neededColumns, 5),
		GroupBy:    GenerateExpr(neededColumns, 5),
		Having:     GenerateExpr(neededColumns, 5),
		OrderBy:    orderByExpr,
		OrderOpt:   orderByOpt,
		Limit:      -1, // 未确定好的生成策略
		Offset:     -1,
		ForOpt:     forOpt,
	}
}

func GenerateJoinAst(tables []*internal.Table) ast.AstNode {
	var preNode, curNode ast.AstNode
	preNode = &ast.JoinNode{
		JoinType: randomly.RandIntGap(ast.JoinTypeInner, ast.JoinTypeNatural+1),
		Left:     &ast.TabRefNode{Table: tables[0]},
		Right:    nil,
	}
	for _, tab := range tables[1:] {
		curNode = &ast.JoinNode{
			JoinType: randomly.RandIntGap(ast.JoinTypeInner, ast.JoinTypeNatural+1),
			Left:     preNode,
			Right:    &ast.TabRefNode{Table: tab},
		}
		preNode = curNode
	}
	return curNode
}
