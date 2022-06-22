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
			} else if val == statement.ColOptColumnFormat {
				columns[i].ColForMat = statement.ColFormats(randomly.RandIntGap(statement.ColForMatDEFAULT, statement.ColForMatFixed))
			} else if val == statement.ColOptStorage {
				columns[i].Storage = statement.StorageOption(randomly.RandIntGap(statement.StorageDisk, statement.StorageMemory))
			} else if val == statement.ColOptComment {
				columns[i].Comment = randomly.RandAlphabetStrLen(4)
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
	partCol := columns[randomly.RandIntGap(0, colNum-1)]
	partCols := make([]statement.Column, 0)
	for _, val := range columns {
		if randomly.RandBool() {
			partCols = append(partCols, val)
		}
	}

	return &statement.CreateTableStmt{
		TableName:          tableName,
		Columns:            columns,
		TableOptions:       candidates,
		AutoIncrement:      randomly.RandIntGap(1, 10),
		AvgRowLength:       randomly.RandIntGap(1, 10),
		Compression:        statement.Compressions(randomly.RandIntGap(statement.CompreLZ4, statement.CompreZLIB)),
		DelayKeyWrite:      randomly.RandBool(),
		InsertMethod:       statement.InsertMethods(randomly.RandIntGap(statement.InsertFirst, statement.InsertNo)),
		KeyBlockSize:       randomly.RandIntGap(1, 10),
		MaxRows:            randomly.RandIntGap(5, 10),
		MinRows:            randomly.RandIntGap(1, 5),
		PackKey:            randomly.RandIntGap(-1, 1),
		StatsAutoRecalc:    randomly.RandIntGap(-1, 1),
		StatsPersistent:    randomly.RandIntGap(-1, 1),
		StatsSamplePages:   randomly.RandIntGap(1, 10),
		TableEngine:        tableEngine,
		PartitionOption:    partOpt,
		PartitionColumn:    partCol,
		PartitionAlgorithm: randomly.RandIntGap(-1, 1),
		PartitionColumns:   partCols,
	}
}
