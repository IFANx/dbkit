package main

import (
	"dbkit/admin"
	"dbkit/internal"
	"dbkit/internal/mysql"

	log "github.com/sirupsen/logrus"
)

func main() {
	defer func() {
		log.Info("End of run, clean up resources")
		clean()
	}()
	admin.StartServer(8080)
}

func RunMySQLQueryTest() {
	testCtx := internal.NewTestContext()
	tester := mysql.NewMySQLQueryTester(testCtx)
	testCtx.SetTester(tester)
	state.Tests = append(state.Tests)
	testCtx.Start()
}
