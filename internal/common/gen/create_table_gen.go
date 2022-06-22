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

	for i := statement.ColOptColumnFormat - 1; i < statement.ColOptUnique; i++ {
		candidates = append(candidates, i)
	}

	hasPrimaryKey := false

	for i := 0; i < colNum; i++ {
		columns[i].Name = "c" + strconv.Itoa(i)
		columns[i].Type = gen.MySQLDataType(randomly.RandIntGap(gen.TypeBigInt, gen.TypeYear))
		constraint := randomly.RandPickNotEmptyInt(candidates)
		isNull := false
		for idx, val := range constraint {
			if val == statement.ColOptNull {
				isNull = true
			} else if val == statement.ColOptPrimaryKey {
				if !hasPrimaryKey && !isNull && columns[i].Type.CanBePrimary() {
					hasPrimaryKey = true
				} else {
					constraint[idx] = -1
				}
			}
		}
		columns[i].Constraint = constraint
	}

	candidates = make([]int, 0)

	for i := statement.TabOptAutoIncrement - 1; i < statement.TabOptStatsSamplePages; i++ {
		candidates = append(candidates, i)
	}
	candidates = randomly.RandPickNotEmptyInt(candidates)

	var tableEngine statement.TableEngines = statement.TabEngInnoDB
	for _, val := range candidates {
		if val == statement.TabOptEngine {
			tableEngine = statement.TableEngines(randomly.RandIntGap(statement.TabEngArchive, statement.TabEngMyISAM))
		}
	}

	partOpt := statement.PartitionOptions(randomly.RandIntGap(statement.PartOptHASH-1, statement.PartOptKEY))

	return &statement.CreateTableStmt{
		TableName:       tableName,
		Columns:         columns,
		TableOptions:    candidates,
		TableEngine:     tableEngine,
		PartitionOption: partOpt,
	}
}
