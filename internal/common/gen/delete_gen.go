package gen

import (
	"dbkit/internal/common"
	"dbkit/internal/common/ast"
	"dbkit/internal/common/statement"
	"dbkit/internal/randomly"
	"math/rand"
	"time"
)

func GenerateDeleteStmt(tables []*common.Table, partitions []string) *statement.DeleteStmt {
	rand.Seed(time.Now().UnixNano())

	delOptList := make([]statement.DeleteOption, 0)
	// 需要添加控制选项的开关
	if true { // 可以生成参数
		// 需要引擎参数
		if true { // innodb引擎
			if randomly.RandBool() {
				delOptList = RandPickOptions(delOptList, statement.DelOptIgnore)
			} else {
				delOptList = nil
			}
		} else {
			if randomly.RandBool() {
				delOptList = RandPickOptions(delOptList, statement.DelOptQuick)
			} else {
				delOptList = nil
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

	orderByColumns := make([]*common.Column, 0)
	if len(tables) > 0 {
		// 需要添加控制选项的开关
		if true { // 可以生成ORDER BY
			if randomly.RandBool() {
				orderByColumns = RandPickColumns(neededColumns)
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

	return &statement.DeleteStmt{
		Options:    delOptList,
		Tables:     neededTables,
		Join:       joinAst,
		JoinOn:     joinOnAst,
		Partitions: parList,
		Where:      GenerateExpr(neededColumns, 5),
		OrderBy:    orderByColumns,
		OrderOpt:   orderByOpt,
		Limit:      -1,
	}
}

func RandPickOptions(optList []int, optSize int) []int {
	optList = rand.Perm(optSize + 1) // len(optList) = optSize + 1
	delOption := randomly.RandIntGap(1, optSize+1)
	rand.Shuffle(optSize+1, func(i, j int) {
		optList[i], optList[j] = optList[j], optList[i]
	})
	optList = optList[:delOption]
	return optList
}

func RandPickStrings(stringList []string) []string {
	listSize := len(stringList)
	randomSize := randomly.RandIntGap(1, listSize)
	elements := make([]string, len(stringList))
	copy(elements, stringList)
	rand.Shuffle(listSize, func(i, j int) {
		elements[i], elements[j] = elements[j], elements[i]
	})
	return elements[:randomSize]
}

func RandPickColumns(columns []*common.Column) []*common.Column {
	colNum := len(columns)
	randomSize := randomly.RandIntGap(1, colNum)
	elements := make([]*common.Column, len(columns))
	copy(elements, columns)
	rand.Shuffle(colNum, func(i, j int) {
		elements[i], elements[j] = elements[j], elements[i]
	})
	return elements[:randomSize]
}
