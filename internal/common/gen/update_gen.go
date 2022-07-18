package gen

import (
	"dbkit/internal/common"
	"dbkit/internal/common/ast"
	"dbkit/internal/common/statement"
	"dbkit/internal/randomly"
	"math/rand"
	"time"
)

func GenerateUpdateStmt(tables []*common.Table, partitions []string) *statement.UpdateStmt {
	rand.Seed(time.Now().UnixNano())
	updOptList := make([]statement.UpdateOption, 0)
	// 需要添加控制选项的开关
	if true { // 可以生成参数
		// 需要引擎参数
		if true { // innodb引擎
			if randomly.RandBool() {
				updOptList = RandPickOptions(updOptList, statement.UpdOptIgnore)
			} else {
				updOptList = nil
			}
		} else {
			if randomly.RandBool() {
				updOptList = RandPickOptions(updOptList, statement.UpdOptLowPriority)
			} else {
				updOptList = nil
			}
		}
	}

	parList := make([]string, 0)
	// 需要添加控制选项的开关
	if true { // 可以生成partition
		if randomly.RandBool() && len(partitions) > 0 {
			parList = RandPickStrings(partitions)
		}
	}

	neededTables := make([]*common.Table, 0)
	// 需要单多表参数
	if true { // 只支持单表的方法
		neededTables = append(neededTables, tables[0])
	} else {
		if randomly.RandBool() && len(tables) > 1 { // 多表
			neededTables = RandPickNotEmptyTab(tables)
		} else {
			neededTables = append(neededTables, tables[0])
		}
	}

	neededColumns := make([]*common.Column, 0)
	for _, tab := range neededTables {
		for _, col := range tab.Columns {
			neededColumns = append(neededColumns, col)
		}
	}

	var joinAst, joinOnAst ast.AstNode
	if len(neededTables) > 1 && randomly.RandBool() {
		joinAst = GenerateJoinAst(neededTables)
		joinOnAst = GenerateExpr(neededColumns, 3)
	}

	updColumns := make([]ast.ColRefNode, 0)
	updColNum := randomly.RandIntGap(1, len(neededColumns))
	randColumns := RandPickColumns(neededColumns, updColNum)
	for _, col := range randColumns {
		updColumns = append(updColumns, ast.ColRefNode{Column: col})
	}

	updExprList := make([]string, 0)
	for i := 0; i < updColNum; i++ {
		if true { // 待修改
			updExprList = append(updExprList, randColumns[i].Type.GenRandomVal())
		} else {
			// 待修改
			//updExprList = append(updExprList, GenerateExpr(updColumns, 3))
		}
	}

	randOrderByColumns := make([]*common.Column, 0)
	orderByColumns := make([]ast.ColRefNode, 0)
	if len(tables) > 0 {
		// 需要添加控制选项的开关
		if true { // 可以生成ORDER BY
			if randomly.RandBool() {
				randOrderByColumns = RandPickOrderColumns(neededColumns)
				for _, col := range randOrderByColumns {
					orderByColumns = append(orderByColumns, ast.ColRefNode{Column: col})
				}
			}
		}
	}

	orderByOpt := -1
	if randomly.RandBool() && len(orderByColumns) > 0 {
		orderByOpt = randomly.RandIntGap(0, 1)
	}

	// 待修改
	// 需要添加控制选项的开关
	if randomly.RandBool() && orderByOpt == -1 {

	}

	return &statement.UpdateStmt{
		Options:    updOptList,
		Tables:     neededTables,
		Join:       joinAst,
		JoinOn:     joinOnAst,
		Partitions: parList,
		UpdateCol:  updColumns,
		UpdateExpr: updExprList,
		Where:      GenerateExpr(neededColumns, 3),
		OrderBy:    orderByColumns,
		OrderOpt:   orderByOpt,
		Limit:      -1,
	}
}

func RandPickColumns(columns []*common.Column, randomSize int) []*common.Column {
	colNum := len(columns)
	elements := make([]*common.Column, len(columns))
	copy(elements, columns)
	rand.Shuffle(colNum, func(i, j int) {
		elements[i], elements[j] = elements[j], elements[i]
	})
	return elements[:randomSize]
}
