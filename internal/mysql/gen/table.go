package gen

import (
	"dbkit/internal/common/stmt"
	"dbkit/internal/mysql"
	"dbkit/internal/randomly"
	"strconv"
)

func GenCreateTableStmt(tableName string) stmt.CreateTableStmt {
	var (
		colNames   []string
		colTypes   map[string]string
		colOptions map[string]string
		tabOptions map[string]string
		partition  = ""
	)

	columnCnt := randomly.RandIntGap(2, 6)
	for i := 0; i < columnCnt; i++ {
		colName := "c" + strconv.Itoa(i)
		colType := mysql.RandMySQLType()
		colTypes[colName] = getTypeString(colType)
		colOptions[colName] = getColOptions(colType)
	}

	return stmt.CreateTableStmt{
		TableName:       tableName,
		Columns:         colNames,
		ColumnTypes:     colTypes,
		ColumnOptions:   colOptions,
		TableOptions:    tabOptions,
		PartitionOption: partition,
	}
}

func getColOptions(dataType mysql.MySQLDataType) string {
	return ""
}

func getTypeString(dataType mysql.MySQLDataType) string {
	var res string
	switch dataType {
	case mysql.TypeDecimal:
		res = "DECIMAL"
		if randomly.RandBool() {
			res += getPrecisionAndScale()
		}
	case mysql.TypeInt:
		res = randomly.RandPickOneStr([]string{"TINYINT", "SMALLINT", "MEDIUMINT", "INT", "BIGINT"})
	case mysql.TypeVarchar:
		res = randomly.RandPickOneStr([]string{"VARCHAR(50)", "TINYTEXT", "TEXT", "MEDIUMTEXT", "LONGTEXT"})
	case mysql.TypeFloat:
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
