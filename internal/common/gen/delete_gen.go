package gen

import (
	"dbkit/internal"
	"dbkit/internal/ast"
	"math/rand"
	"time"
)

func GenerateDeleteStmt(table *internal.Table, partitions []string) *ast.DeleteStmt {
	rand.Seed(time.Now().UnixNano())
	// 需要添加控制选项的开关
	delOption := 0
	if rand.Intn(2) == 1 {
		delOption = 1
	}
	// 需要添加控制选项的开关
	if rand.Intn(2) == 0 {
		partitions = nil
	}

	neededColumns := make([]*internal.Column, 0)
	for _, col := range table.Columns {
		neededColumns = append(neededColumns, col)
	}

	// 需要添加控制选项的开关
	colNum := rand.Intn(len(neededColumns) + 1)
	var orderByColumns []*internal.Column
	if colNum == 0 {
		orderByColumns = nil
	} else {
		if colNum == len(neededColumns) {
			orderByColumns = make([]*internal.Column, 0)
			for _, v := range rand.Perm(len(neededColumns)) {
				orderByColumns = append(orderByColumns, neededColumns[v])
			}
		} else {
			orderByColumns = make([]*internal.Column, len(neededColumns))
			copy(orderByColumns, neededColumns)
			for i := len(neededColumns); i > colNum; i-- {
				r := rand.Intn(len(orderByColumns))
				orderByColumns = append(orderByColumns[:r], orderByColumns[r+1:]...)
			}
		}
	}

	var orderByOpt ast.OrderOption
	if rand.Intn(2) == 1 {
		orderByOpt = rand.Intn(2)
	} else {
		orderByOpt = -1
	}

	// 待修改
	// 需要添加控制选项的开关
	if rand.Intn(2) == 1 && orderByOpt == -1 {

	}

	return &ast.DeleteStmt{
		Option:     delOption,
		Table:      *table,
		Partitions: partitions,
		Where:      GenerateExpr(neededColumns, 5),
		OrderBy:    orderByColumns,
		OrderOpt:   orderByOpt,
		Limit:      -1,
	}
}
