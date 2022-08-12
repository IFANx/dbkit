package mysql

import (
	"dbkit/internal/common"
	"dbkit/internal/mysql/troc"
)

type MySQLTrocTester struct{}

func (tester *MySQLTrocTester) RunTask(ctx common.OracleRuntime) {
	troc.RunTrocWithCtx(ctx)
}
