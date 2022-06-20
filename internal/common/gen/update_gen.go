package gen

import (
	"dbkit/internal"
	"dbkit/internal/ast"
	"dbkit/internal/randomly"
	"math/rand"
	"time"
)

func GenerateUpdateStmt(table *internal.Table, partitions []string) *ast.UpdateStmt {
	// 需要添加控制选项的开关
	updOption := []int{0, 1, 2}
	rand.Seed(time.Now().UnixNano())
	optNum := rand.Intn(len(updOption) + 1) // 从0到updOption的长度中随机选一个数
	switch optNum {
	case 0:
		updOption = nil
	case 1:
		one := rand.Intn(len(updOption))
		updOption = append(updOption[one : one+1])
	case 2:

	}
	// 需要添加控制选项的开关
	if randomly.RandBool() {
		partitions = nil
	}

	neededColumns := make([]*internal.Column, 0)
	for _, col := range table.Columns {
		neededColumns = append(neededColumns, col)
	}

	var updColumns []*internal.Column
	updColNum := rand.Intn(len(neededColumns)) + 1
	updColumns = GenerateRandColumns(neededColumns, updColNum)

	updExprList := make([]ast.AstNode, updColNum)
	for i := 0; i < updColNum; i++ {
		if randomly.RandBool() {
			updExprList = append(updExprList) // 待修改
		} else {
			updExprList = append(updExprList, GenerateExpr(updColumns, 3))
		}
	}

	// 需要添加控制选项的开关
	var orderByColumns []*internal.Column
	orderNum := rand.Intn(len(neededColumns) + 1)
	if orderNum == 0 {
		orderByColumns = nil
	} else {
		orderByColumns = GenerateRandColumns(neededColumns, orderNum)
	}

	var orderByOpt ast.OrderOption
	if randomly.RandBool() {
		orderByOpt = rand.Intn(2)
	} else {
		orderByOpt = -1
	}

	// 待修改
	// 需要添加控制选项的开关
	if randomly.RandBool() && orderByOpt == -1 {

	}

	return &ast.UpdateStmt{
		Options:    updOption,
		Table:      *table,
		Partitions: partitions,
		UpdateCol:  updColumns,
		Where:      GenerateExpr(neededColumns, 5),
		OrderBy:    orderByColumns,
		OrderOpt:   orderByOpt,
		Limit:      -1,
	}
}

func GenerateRandColumns(neededColumns []*internal.Column, colNum int) []*internal.Column {
	orderByColumns := make([]*internal.Column, len(neededColumns))
	copy(orderByColumns, neededColumns)
	for i := len(neededColumns); i > colNum; i-- {
		r := rand.Intn(len(orderByColumns))
		orderByColumns = append(orderByColumns[:r], orderByColumns[r+1:]...)
	}
	return orderByColumns
}
