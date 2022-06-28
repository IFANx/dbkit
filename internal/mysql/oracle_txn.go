package mysql

import (
	"dbkit/internal"
	"dbkit/internal/common"
	"dbkit/internal/mysql/gen"

	log "github.com/sirupsen/logrus"
)

type MySQLTrocTester struct{}

func (tester *MySQLTrocTester) RunTask(ctx *internal.TaskContext) {
	table := &common.Table{
		DB:     ctx.DBList[0],
		Name:   "t",
		DBName: "test",
	}
	for {
		table.Build()

		for run := 0; run < 20; run++ {
			ctx.TestRunCount++
			predicate := gen.GenPredicate(table)
			log.Infof("生成新的谓词：%s", predicate)
			NoRECWithCtx(ctx, table, predicate)
			TLPWithCtx(ctx, table, predicate)
		}
	}
}
