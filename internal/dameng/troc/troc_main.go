package troc

import (
	"dbkit/internal/common"
	"dbkit/internal/common/isolation"
	"dbkit/internal/common/stmt"
	"dbkit/internal/randomly"
	"time"
)

func RunTrocWithCtx(ctx common.OracleRuntime) {
	dbInstance := ctx.GetDBList()[0]
	table := &common.Table{
		DB:   dbInstance,
		Name: "t",
	}
	for {
		table.Build()
		time.Sleep(time.Second * 5)
		if ctx.IsAborted() {
			break
		}
	}
}

func genTransaction(table *common.Table) *common.Transaction {
	level := isolation.GetRandomIsolationLevel()
	statements := []stmt.Statement{&stmt.BeginStmt{}}
	n := randomly.RandIntGap(2, 5)
	for i := 0; i < n; i++ {
		x := randomly.RandIntGap(0, 3)
		switch x {
		case 0:
			statements = append(statements, table.DB.DBProvider.GenSelectStmt(table))
		case 1:
			statements = append(statements, table.DB.DBProvider.GenUpdateStmt(table))
		case 2:
			statements = append(statements, table.DB.DBProvider.GenInsertStmt(table))
		case 3:
			statements = append(statements, table.DB.DBProvider.GenDeleteStmt(table))
		}
	}
	return &common.Transaction{
		Isolation: level,
		Stmts:     statements,
	}
}
