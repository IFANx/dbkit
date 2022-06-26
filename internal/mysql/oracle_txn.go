package mysql

import (
	"dbkit/internal"
	"dbkit/internal/common"
	"dbkit/internal/common/dbms"
	"dbkit/internal/mysql/gen"

	log "github.com/sirupsen/logrus"
)

type MySQLTrocTester struct {
	TestCtx *internal.TaskContext
}

func NewMySQLTrocTester(testCtx *internal.TaskContext) *MySQLTrocTester {
	return &MySQLTrocTester{TestCtx: testCtx}
}

func (tester *MySQLTrocTester) RunTest() {
	ctx := tester.TestCtx
	table := &common.Table{
		TestCtx:    tester.TestCtx,
		DBMS:       dbms.MYSQL,
		Name:       "t",
		DBName:     "test",
		DBProvider: &MySQLProvider{},
	}
	for {
		table.Build()

		for run := 0; run < 20; run++ {
			ctx.CountTestRun()
			predicate := gen.GenPredicate(table)
			log.Infof("生成新的谓词：%s", predicate)
			NoRECWithCtx(ctx, table, predicate)
			TLPWithCtx(ctx, table, predicate)
		}
	}
}
