package gen

import (
	"dbkit/internal/common"
	"dbkit/internal/common/ast"
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
	// 生成一行数据
	for i := 0; i < len(neededColumns); i++ {
		if true { // 待修改
			insValueList = append(insValueList, neededColumns[i].Type.GenRandomVal())
		} else {
			// 待修改
			//insValueList = append(insValueList, GenerateExpr(neededColumns, 3))
		}
	}

	insertColumns := make([]ast.ColRefNode, 0)
	for _, col := range neededColumns {
		insertColumns = append(insertColumns, ast.ColRefNode{Column: col})
	}

	dupColumns := make([]ast.ColRefNode, 0)
	var dupExprList []string
	// 需要添加控制选项的开关
	if true { // 可以生成Duplicate
		randDup := randomly.RandIntGap(0, 3)
		if randDup == 3 {
			dupColNum := randomly.RandIntGap(1, len(neededColumns))
			randColumns := RandPickColumns(neededColumns, dupColNum)
			for _, col := range randColumns {
				dupColumns = append(dupColumns, ast.ColRefNode{Column: col})
			}

			dupExprList = make([]string, 0)
			for i := 0; i < dupColNum; i++ {
				if true { // 待修改
					dupExprList = append(dupExprList, randColumns[i].Type.GenRandomVal())
				} else {
					// 待修改
					//dupExprList = append(dupExprList, GenerateExpr(dupColumns, 3))
				}
			}
		}
	}

	return &statement.InsertStmt{
		Options:     insOptList,
		Table:       table,
		Partitions:  parList,
		InsertCol:   insertColumns,
		InsertValue: insValueList,
		DupCol:      dupColumns,
		DupValue:    dupExprList,
	}
}
