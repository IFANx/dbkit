package mysql

import (
	"dbkit/internal"
	"dbkit/internal/common"
	"dbkit/internal/common/stmt"
	"dbkit/internal/mysql/gen"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type MySQLQueryTester struct {
	TestCtx *internal.TestContext
}

func NewMySQLQueryTester(testCtx *internal.TestContext) *MySQLQueryTester {
	return &MySQLQueryTester{TestCtx: testCtx}
}

func (tester *MySQLQueryTester) RunTest() {
	state := internal.GetState()
	ctx := tester.TestCtx
	table := &internal.Table{
		TestCtx:    tester.TestCtx,
		DBMS:       common.MYSQL,
		Name:       "t",
		DBName:     state.Config.MySQL.DBName,
		DBProvider: &MySQLProvider{},
	}
	for {
		table.Build()

		for run := 0; run < 20; run++ {
			ctx.CountTestRun()
			predicate := gen.GenPredicate(table)

			// NoREC
			norec := stmt.SelectStmt{
				TableName: table.Name,
				Targets:   []string{"count(1)"},
				Predicate: predicate,
				ForShare:  false,
				ForUpdate: false,
			}
			res, err := ctx.Query(norec)
			if err != nil {
				continue
			}
			count1Str := string(res[0][0].([]byte))
			count1, err := strconv.Atoi(count1Str)
			if err != nil {
				continue
			}
			norec.Predicate = "True"
			norec.Targets[0] = "((" + predicate + ") IS TRUE)"
			res, err = ctx.Query(norec)
			if err != nil {
				continue
			}
			count2 := 0
			for _, row := range res {
				tmpStr := string(row[0].([]byte))
				tmp, err := strconv.Atoi(tmpStr)
				if err != nil {
					continue
				}
				count2 += tmp
			}
			log.Infof("count1=%d, count2=%d", count1, count2)
			// TLP
		}
	}
}
