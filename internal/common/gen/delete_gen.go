package gen

import (
	"dbkit/internal"
	"dbkit/internal/ast"
	"dbkit/internal/randomly"
	"math/rand"
	"time"
)

func GenerateDeleteStmt(table *internal.Table, partitions []string) *ast.DeleteStmt {
	// 需要添加控制选项的开关
	delOption := []int{0, 1, 2}
	rand.Seed(time.Now().UnixNano())
	optNum := rand.Intn(len(delOption) + 1) // 从0到delOption的长度中随机选一个数
	switch optNum {
	case 0:
		delOption = nil
	case 1:
		one := rand.Intn(len(delOption))
		delOption = append(delOption[one : one+1])
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

	// 需要添加控制选项的开关
	colNum := rand.Intn(len(neededColumns) + 1)
	orderByColumns := make([]*internal.Column, len(neededColumns))
	if colNum == 0 {
		orderByColumns = nil
	} else {
		copy(orderByColumns, neededColumns)
		for i := len(neededColumns); i > colNum; i-- {
			r := rand.Intn(len(orderByColumns))
			orderByColumns = append(orderByColumns[:r], orderByColumns[r+1:]...)
		}
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

	return &ast.DeleteStmt{
		Options:    delOption,
		Table:      *table,
		Partitions: partitions,
		Where:      GenerateExpr(neededColumns, 5),
		OrderBy:    orderByColumns,
		OrderOpt:   orderByOpt,
		Limit:      -1,
	}
}
