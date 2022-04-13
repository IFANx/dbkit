package main

import (
	"dbkit/internal"
	"dbkit/internal/mysql"
	log "github.com/sirupsen/logrus"
)

func main() {
	defer func() {
		log.Info("运行结束，清理资源")
		clean()
	}()
	testCtx := internal.NewTestContext()
	tester := mysql.NewMySQLQueryTester(testCtx)
	testCtx.SetTester(tester)
	state.Tests = append(state.Tests)
	testCtx.Start()
}
