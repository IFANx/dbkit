package gen

import "dbkit/internal/randomly"

type MySQLAggregate struct {
	name     string
	argCnt   int
	distinct bool
}

func GetRandomMySQLAggregate() MySQLAggregate {
	return aggregates[randomly.RandIntGap(0, len(aggregates)-1)]
}

var aggregates = []MySQLAggregate{
	{"AVG", 1, true}, {"BIT_AND", 1, false},
	{"BIT_OR", 1, false}, {"BIT_XOR", 1, false},
	{"COUNT", 1, true}, {"GROUP_CONCAT", 1, false},
	{"JSON_ARRAYAGG", 1, false}, {"JSON_OBJECTAGG", 2, false},
	{"MAX", 1, true}, {"MIN", 1, true},
	{"SUM", 1, true},
	{"STD", 1, false}, {"STDDEV", 1, false},
	{"STDDEV_POP", 1, false}, {"STDDEV_SAMP", 1, false},
	{"VAR_POP", 1, false}, {"VAR_SAMP", 1, false},
	{"VARIANCE", 1, false},
}
