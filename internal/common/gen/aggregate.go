package gen

import "dbkit/internal/randomly"

type MySQLAggregate struct {
	name   string
	argCnt int
}

func GetRandomMySQLAggregate() MySQLAggregate {
	n := len(aggregates)
	return aggregates[randomly.RandIntGap(0, n)]
}

var aggregates = []MySQLAggregate{
	{"AVG", 1}, {"BIT_AND", 1},
	{"BIT_OR", 1}, {"BIT_XOR", 1},
	{"COUNT", 1}, {"GROUP_CONCAT", 1},
	{"MAX", 1}, {"MIN", 1},
	{"SUM", 1},
	{"STD", 1}, {"STDDEV", 1},
	{"STDDEV_POP", 1}, {"STDDEV_SAMP", 1},
	{"VAR_POP", 1}, {"VAR_SAMP", 1},
	{"VARIANCE", 1},
}
