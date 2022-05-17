package mysql

import (
	"dbkit/internal"
	"dbkit/internal/common"
	"dbkit/internal/mysql/gen"
	log "github.com/sirupsen/logrus"
)

type MySQLTrocTester struct {
	TestCtx *internal.TestContext
}

func NewMySQLTrocTester(testCtx *internal.TestContext) *MySQLTrocTester {
	return &MySQLTrocTester{TestCtx: testCtx}
}

func (tester *MySQLTrocTester) RunTest() {
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
			log.Infof("生成新的谓词：%s", predicate)
			NoRECWithCtx(ctx, table, predicate)
			TLPWithCtx(ctx, table, predicate)
		}
	}
}
