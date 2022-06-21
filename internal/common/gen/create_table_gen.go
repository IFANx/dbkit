package gen

import (
	"dbkit/internal/common/statement"
	"dbkit/internal/mysql/gen"
	"dbkit/internal/randomly"
	"strconv"
)

func GenerateCreateTableStmt(tableName string) *statement.CreateTableStmt {

	colNum := randomly.RandIntGap(1, 3)
	columns := make([]statement.Column, 0)

	candidates := make([]int, 0)

	for i := statement.ColOptColumnFormat; i < statement.ColOptUnique; {
		candidates = append(candidates, i)
	}

	for i := 0; i < colNum; i++ {
		columns[i].Name = "c" + strconv.Itoa(i)
		columns[i].Type = randomly.RandIntGap(gen.TypeBigInt, gen.TypeYear)
		columns[i].Constraint = randomly.RandPickNotEmptyInt(candidates)
	}

	return &statement.CreateTableStmt{
		TableName: tableName,
		Columns:   columns,
	}
}
