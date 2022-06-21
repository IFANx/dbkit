package gen

import (
	"dbkit/internal/common"
	"dbkit/internal/common/ast"
	"dbkit/internal/common/statement"
	"math/rand"
	"time"
)

func GenerateUpdateStmt(table *common.Table, partitions []string) *statement.UpdateStmt {
	rand.Seed(time.Now().UnixNano())
	// 需要添加控制选项的开关
	updOption := 0
	if rand.Intn(2) == 1 {
		updOption = 1
	}

	// 需要添加控制选项的开关
	if rand.Intn(2) == 0 {
		partitions = nil
	}

	neededColumns := make([]*common.Column, 0)
	for _, col := range table.Columns {
		neededColumns = append(neededColumns, col)
	}

	var updColumns []*common.Column
	updColNum := rand.Intn(len(neededColumns)) + 1
	updColumns = GenerateRandColumns(neededColumns, updColNum)

	updExprList := make([]ast.AstNode, updColNum)
	for i := 0; i < updColNum; i++ {
		if rand.Intn(2) == 1 {
			updExprList = append(updExprList) // 待修改
		} else {
			updExprList = append(updExprList, GenerateExpr(updColumns, 3))
		}
	}

	// 需要添加控制选项的开关
	var orderByColumns []*common.Column
	orderNum := rand.Intn(len(neededColumns) + 1)
	if orderNum == 0 {
		orderByColumns = nil
	} else {
		orderByColumns = GenerateRandColumns(neededColumns, orderNum)
	}

	var orderByOpt statement.OrderOption
	if rand.Intn(2) == 1 && orderNum != 0 {
		orderByOpt = rand.Intn(2)
	} else {
		orderByOpt = -1
	}

	// 待修改
	// 需要添加控制选项的开关
	if rand.Intn(2) == 1 && orderByOpt == -1 {

	}

	return &statement.UpdateStmt{
		Option:     updOption,
		Table:      *table,
		Partitions: partitions,
		UpdateCol:  updColumns,
		Where:      GenerateExpr(neededColumns, 5),
		OrderBy:    orderByColumns,
		OrderOpt:   orderByOpt,
		Limit:      -1,
	}
}

func GenerateRandColumns(neededColumns []*common.Column, colNum int) []*common.Column {
	var orderByColumns []*common.Column
	if colNum == len(neededColumns) {
		orderByColumns = make([]*common.Column, 0)
		for _, v := range rand.Perm(len(neededColumns)) {
			orderByColumns = append(orderByColumns, neededColumns[v])
		}
	} else {
		orderByColumns = make([]*common.Column, len(neededColumns))
		copy(orderByColumns, neededColumns)
		for i := len(neededColumns); i > colNum; i-- {
			r := rand.Intn(len(orderByColumns))
			orderByColumns = append(orderByColumns[:r], orderByColumns[r+1:]...)
		}
	}
	return orderByColumns
}
