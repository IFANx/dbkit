package gen

import "dbkit/internal/randomly"

type MySQLFunc struct {
	name   string
	argCnt int
}

func GetRandomMySQLFunc() MySQLFunc {
	n := len(functions)
	return functions[randomly.RandIntGap(0, n-1)]
}

var functions = []MySQLFunc{
	{"abs", 1}, {"acos", 1},
	{"ASCII", 1}, {"ASIN", 1},
	{"ATAN", 1}, {"ATAN2", 1},
	{"BIN", 1}, {"CEIL", 1},
	{"CEILING", 1}, {"CHAR", 1},
	{"COMPRESS", 1}, {"CONCAT", 1},
	{"COS", 1},
	{"COT", 1}, {"CRC32", 1},
	{"CURTIME", 0}, {"DATABASE", 0},
	{"DEGREES", 1}, {"EXP", 1},
	{"FLOOR", 1}, {"FROM_BASE64", 1},
	{"HEX", 1}, {"IF", 3},
	{"IFNULL", 2}, {"ISNULL", 1},
	{"LCASE", 1}, {"LEAST", 3},
	{"LEFT", 2}, {"LENGTH", 1},
	{"LN", 1}, {"LOG", 2},
	{"LOG", 2}, {"LOG10", 1},
	{"LOG2", 1}, {"LOWER", 1},
	{"LPAD", 3}, {"LTRIM", 1},
	{"MD5", 1}, {"MID", 3},
	{"MOD", 2}, {"NULLIF", 2},
	{"OCT", 1}, {"ORD", 1},
	{"PI", 0}, {"RAND", 0},
	{"REVERSE", 1}, {"RIGHT", 2},
	{"ROUND", 1}, {"RPAD", 3},
	{"RTRIM", 1}, {"SCHEMA", 0},
	{"SHA1", 1}, {"SIGN", 1},
	{"SIN", 1}, {"SPACE", 1},
	{"SQRT", 1}, {"STRCMP", 2},
	{"TAN", 1}, {"TO_BASE64", 1},
	{"UPPER", 1}, {"USER", 0},
	{"UUID", 0}, {"VERSION", 0},
}
