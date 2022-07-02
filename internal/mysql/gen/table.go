package gen

import (
	"dbkit/internal/common"
	"dbkit/internal/common/stmt"
	"dbkit/internal/randomly"
	"strconv"
)

func GenCreateTableStmt(table *common.Table) *stmt.CreateTableStmt {
	var (
		colNames   = make([]string, 0)
		colTypes   = make(map[string]string)
		colOptions = make(map[string]string)
		tabOptions = make(map[string]string)
		partition  = ""
	)

	columnCnt := randomly.RandIntGap(2, 5)
	for i := 0; i < columnCnt; i++ {
		colName := "c" + strconv.Itoa(i)
		colNames = append(colNames, colName)
		colType := RandMySQLType()
		colTypes[colName] = getTypeString(colType)
		colOptions[colName] = getColOptions(colType)
	}

	return &stmt.CreateTableStmt{
		TableName:       table.Name,
		Columns:         colNames,
		ColumnTypes:     colTypes,
		ColumnOptions:   colOptions,
		TableOptions:    tabOptions,
		PartitionOption: partition,
	}
}

func getColOptions(dataType MySQLDataType) string {
	return ""
}

func getTypeString(dataType MySQLDataType) string {
	var res string
	switch dataType {
	case TypeDecimal:
		res = "DECIMAL"
		if randomly.RandBool() {
			res += getPrecisionAndScale()
		}
	case TypeInt:
		res = randomly.RandPickOneStr([]string{"TINYINT", "SMALLINT", "MEDIUMINT", "INT", "BIGINT"})
	case TypeVarchar:
		res = randomly.RandPickOneStr([]string{"VARCHAR(50)", "TINYTEXT", "TEXT", "MEDIUMTEXT", "LONGTEXT"})
	case TypeFloat:
		res = randomly.RandPickOneStr([]string{"DOUBLE", "FLOAT"})
		if randomly.RandBool() {
			res += getPrecisionAndScale()
		}
	default:
		panic("Unreachable")
	}
	return res
}

func getPrecisionAndScale() string {
	// For float(M,D), double(M,D) or decimal(M,D), M must be >= D
	// The maximum number of digits (M) for DECIMAL is 65
	// The maximum number of supported decimals (D) is 30
	m := randomly.RandIntGap(1, 65)
	d := randomly.RandIntGap(1, 30)
	if d > m {
		d = m
	}
	return "(" + strconv.Itoa(m) + "," + strconv.Itoa(d) + ")"
}
