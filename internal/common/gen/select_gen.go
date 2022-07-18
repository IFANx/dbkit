package gen

import (
	"dbkit/internal/common"
	"dbkit/internal/common/ast"
	"dbkit/internal/common/statement"
	"dbkit/internal/randomly"
	"math/rand"
)

func GenerateSelectStmt(tables []*common.Table) *statement.SelectStmt {
	selOptList := make([]statement.SelectOption, 0)
	selOptIdx := randomly.RandIntGap(-1, statement.SelOptSqlCalcFoundRows+1)
	if selOptIdx > -1 {
		selOptList = append(selOptList, selOptIdx)
	}
	neededTables := make([]*common.Table, 0)
	if randomly.RandBool() && len(tables) > 1 {
		neededTables = RandPickNotEmptyTab(tables)
	} else {
		neededTables = append(neededTables, tables[0])
	}
	neededColumns := make([]*common.Column, 0)
	for _, tab := range neededTables {
		for _, col := range tab.Columns {
			neededColumns = append(neededColumns, col)
		}
	}
	selExprList := make([]ast.AstNode, 0)
	selColumns := make([]*common.Column, 0)
	genOnlyColumn := false
	genAggregate := false
	if randomly.RandBool() {
		if randomly.RandBool() { // 只生成列名
			randNum := randomly.RandIntGap(1, len(neededColumns))
			selColumns = RandPickColumns(neededColumns, randNum)
			for _, col := range selColumns {
				selExprList = append(selExprList, &ast.ColRefNode{Column: col})
			}
			genOnlyColumn = true
		} else { // 只生成聚合函数
			selExprList = append(selExprList, RandGenAggregate(neededColumns))
			genAggregate = true
		}
	} else {
		for i := 0; i < randomly.RandIntGap(1, 3); i++ {
			selExprList = append(selExprList, GenerateExpr(neededColumns, 1))
			//selExprList = append(selExprList, GenerateExprWithAggregate(neededColumns, 3))
		}
	}

	var joinAst, joinOnAst ast.AstNode
	if len(neededTables) > 1 && randomly.RandBool() {
		joinAst = GenerateJoinAst(neededTables)
		joinOnAst = GenerateExpr(neededColumns, 3)
	}

	var groupAst ast.AstNode
	if genAggregate {
		if randomly.RandBool() {
			groupAst = GenerateExpr(neededColumns, 1)
		} else {
			groupAst = nil
		}
	}

	var havingAst ast.AstNode
	havingAst = nil
	if randomly.RandBool() {
		if groupAst != nil {
			havingAst = RandGenAggregate(neededColumns)
		} else {
			if genOnlyColumn {
				havingAst = GenerateExpr(selColumns, 1)
			}
		}
	}

	var orderByExpr ast.AstNode
	var orderByOpt statement.OrderOption
	if randomly.RandBool() {
		if groupAst != nil {
			orderByExpr = groupAst
		} else {
			if randomly.RandBool() {
				randOrderByColumn := RandPickOneCol(neededColumns)
				orderByExpr = &ast.ColRefNode{Column: randOrderByColumn}
			} else {
				orderByExpr = GenerateExpr(neededColumns, 3)
			}
		}
	} else {
		orderByExpr = nil
	}

	if randomly.RandBool() && orderByExpr != nil {
		orderByOpt = randomly.RandIntGap(0, 1)
	} else {
		orderByOpt = -1
	}
	var forOpt statement.ForOption
	forOpt = randomly.RandIntGap(statement.ForOptShare-1, statement.ForOptUpdate)
	return &statement.SelectStmt{
		Options:    selOptList,
		SelectExpr: selExprList,
		Tables:     neededTables,
		Join:       joinAst,
		JoinOn:     joinOnAst,
		Partitions: nil, // 需要表结构信息
		Where:      GenerateExpr(neededColumns, 3),
		GroupBy:    groupAst,
		Having:     havingAst,
		OrderBy:    orderByExpr,
		OrderOpt:   orderByOpt,
		Limit:      -1, // 未确定好的生成策略
		Offset:     -1,
		ForOpt:     forOpt,
	}
}

func GenerateJoinAst(tables []*common.Table) ast.AstNode {
	var preNode, curNode ast.AstNode
	preNode = &statement.JoinNode{
		JoinType: randomly.RandIntGap(statement.JoinTypeInner, statement.JoinTypeNatural),
		Left:     &ast.TabRefNode{Table: tables[0]},
		Right:    nil,
	}
	for _, tab := range tables[1:] {
		curNode = &statement.JoinNode{
			JoinType: randomly.RandIntGap(statement.JoinTypeInner, statement.JoinTypeNatural),
			Left:     preNode,
			Right:    &ast.TabRefNode{Table: tab},
		}
		preNode = curNode
	}
	return curNode
}

func RandPickNotEmptyTab(candidates []*common.Table) []*common.Table {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	n := len(candidates)
	x := randomly.RandIntGap(1, n)
	rand.Shuffle(n, func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})
	return candidates[:x]
}
