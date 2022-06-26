package gen

import (
	"dbkit/internal/common"
	"dbkit/internal/common/statement"
	"dbkit/internal/randomly"
)

func GenerateInsertStmt(table *common.Table, partitions []string) *statement.InsertStmt {
	insOptList := make([]statement.InsertOption, 0)
	// 需要添加控制选项的开关
	if true { // 可以生成参数
		// 需要引擎参数
		if true { // innodb引擎
			if randomly.RandBool() {
				insOptList = RandPickOptions(insOptList, statement.InsOptIgnore)
			} else {
				insOptList = nil
			}
		} else {
			if randomly.RandBool() {
				insOptSize := randomly.RandIntGap(1, statement.InsOptHighPriority+1)
				if insOptSize == 1 {
					randOpt := randomly.RandIntGap(statement.InsOptIgnore, statement.InsOptLowPriority)
					insOptList = append(insOptList, randOpt)
				} else {
					insOptList = append(insOptList, statement.InsOptIgnore)
					randOpt := randomly.RandIntGap(statement.InsOptHighPriority, statement.InsOptLowPriority)
					insOptList = append(insOptList, randOpt)
				}
			} else {
				insOptList = nil
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

	neededColumns := make([]*common.Column, 0)
	for _, col := range table.Columns {
		neededColumns = append(neededColumns, col)
	}
	insValueList := make([]string, 0)
	for i := 0; i < len(neededColumns); i++ {
		if true { // 待修改
			insValueList = append(insValueList, neededColumns[i].Type.GenRandomVal())
		} else {
			// 待修改
			//insValueList = append(insValueList, GenerateExpr(neededColumns, 3))
		}
	}

	var dupColumns []*common.Column
	var dupExprList []string
	// 需要添加控制选项的开关
	if true { // 可以生成Duplicate
		dupColNum := randomly.RandIntGap(1, len(neededColumns))
		dupColumns = RandPickColumns(neededColumns)
		dupExprList = make([]string, 0)
		for i := 0; i < dupColNum; i++ {
			if true { // 待修改
				dupExprList = append(dupExprList, dupColumns[i].Type.GenRandomVal())
			} else {
				// 待修改
				//dupExprList = append(dupExprList, GenerateExpr(dupColumns, 3))
			}
		}
	}

	return &statement.InsertStmt{
		Options:     insOptList,
		Table:       table,
		Partitions:  parList,
		InsertCol:   neededColumns,
		InsertValue: insValueList,
		DupCol:      dupColumns,
		DupValue:    dupExprList,
	}
}
