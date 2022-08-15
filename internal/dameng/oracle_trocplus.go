package dameng

import (
	"dbkit/internal/common"
	"dbkit/internal/common/gen"
	"dbkit/internal/common/stmt"
	"dbkit/internal/randomly"
	log "github.com/sirupsen/logrus"
)

type MySQLTrocPlusTester struct{}

func (tester *MySQLTrocPlusTester) RunTask(ctx common.OracleRuntime) {
	dbInstance := ctx.GetDBList()[0]
	table := &common.Table{
		DB:   dbInstance,
		Name: "t",
	}
	for {
		log.Info("生成table")
		table.Build()

		for run := 0; run < 5; run++ {
			ctx.IncrTestRunCount(1)
			log.Info("生成事务")
			RunTrocPlusWithCtx(ctx, table)
			if ctx.IsAborted() {
				break
			}
		}

		if ctx.IsAborted() {
			break
		}
	}
}

func RunTrocPlusWithCtx(ctx common.OracleRuntime, table *common.Table) {
	statements := []stmt.Statement{&stmt.BeginStmt{}}
	ctx.GetDBList()[0].ExecSQL(statements[0].String())
	tables := []*common.Table{table}
	queryTable := "SELECT * FROM " + table.Name
	n := randomly.RandIntGap(2, 5)
	for i := 0; i < n; i++ {
		x := randomly.RandIntGap(0, 3)
		switch x {
		case 0:
			s := gen.GenerateSelectStmt(tables)
			log.Info(s.String())
			res, err := ctx.GetDBList()[0].QuerySQL(s.String())
			if err != nil {
				break
			}
			log.Infof("Query result: %s", res)
			statements = append(statements, s)
		case 1:
			s := gen.GenerateUpdateStmt(tables, nil)
			log.Info(s.String())
			err := ctx.GetDBList()[0].ExecSQL(s.String())
			if err != nil {

			}
			statements = append(statements, s)
		case 2:
			s := gen.GenerateInsertStmt(table, nil)
			log.Info(s.String())
			err := ctx.GetDBList()[0].ExecSQL(s.String())
			if err != nil {

			}
			statements = append(statements, s)
		case 3:
			s := gen.GenerateDeleteStmt(tables, nil)
			log.Info(s.String())
			err := ctx.GetDBList()[0].ExecSQL(s.String())
			if err != nil {

			}
			statements = append(statements, s)
		}
	}
	res, err := ctx.GetDBList()[0].QuerySQL(queryTable)
	if err != nil {
		return
	}
	log.Infof("Query table in txn: %s", res)
	if true { // lack commit statement
		statements = append(statements, &stmt.RollbackStmt{})
		ctx.GetDBList()[0].ExecSQL(statements[len(statements)-1].String())
	}
	res, err = ctx.GetDBList()[0].QuerySQL(queryTable)
	if err != nil {
		return
	}
	log.Infof("Query table out txn: %s", res)
}
