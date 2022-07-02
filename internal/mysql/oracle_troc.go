package mysql

import (
	"dbkit/internal/common"
	"time"
)

type MySQLTrocTester struct{}

func (tester *MySQLTrocTester) RunTask(ctx common.OracleRuntime) {
	dbInstance := ctx.GetDBList()[0]
	table := &common.Table{
		DB:   dbInstance,
		Name: "t",
	}
	for {
		table.Build()
		time.Sleep(time.Minute)
	}
}
