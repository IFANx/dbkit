package mysql

import (
	"dbkit/internal/common"
	"dbkit/internal/common/stmt"
	"dbkit/internal/mysql/gen"
	"dbkit/internal/randomly"
	log "github.com/sirupsen/logrus"
	"time"
)

type MySQLTLPTester struct {
}

func (tester *MySQLTLPTester) RunTask(ctx common.OracleRuntime) {
	table := &common.Table{
		DB:   ctx.GetDBList()[0],
		Name: "t",
	}
	for {
		table.Build()

		for run := 0; run < 20; run++ {
			ctx.IncrTestRunCount(1)
			time.Sleep(time.Second * 5)
			predicate := gen.GenPredicate(table)
			log.Infof("生成新的谓词,s%", predicate)
			TLPWithCtx1(ctx, table, predicate)
			if ctx.IsAborted() {
				break
			}
		}
		if ctx.IsAborted() {
			break
		}
	}
}

func TLPWithCtx1(ctx common.OracleRuntime, table *common.Table, predicate string) {
	target := randomly.RandPickOneStr(table.ColumnNames)
	targets := []string{target}
	tlpOrigin := stmt.SelectStmt{
		TableName: table.Name,
		Targets:   targets,
		Predicate: "True",
		ForShare:  false,
		ForUpdate: false,
	}
	firstQuery := stmt.SelectStmt{
		TableName: table.Name,
		Targets:   targets,
		Predicate: "(" + predicate + ") IS NULL",
		ForShare:  false,
		ForUpdate: false,
	}
	secondQuery := stmt.SelectStmt{
		TableName: table.Name,
		Targets:   targets,
		Predicate: predicate,
		ForShare:  false,
		ForUpdate: false,
	}
	thirdQuery := stmt.SelectStmt{
		TableName: table.Name,
		Targets:   targets,
		Predicate: "NOT(" + predicate + ")",
		ForShare:  false,
		ForUpdate: false,
	}

	tlpQuery := "(" + firstQuery.String() + ") UNION ALL (" + secondQuery.String() +
		") UNION ALL (" + thirdQuery.String() + ")"

	res, err := ctx.GetDBList()[0].Query(tlpOrigin)
	if err != nil || len(res) < 1 {
		return
	}

	resMap1 := make(map[string]int)
	var key string
	for _, rec := range res {
		if rec[0] == nil {
			key = "nil"
		} else {
			key = string(rec[0].([]byte))
		}
		resMap1[key]++
	}
	log.Infof("res: %+v", resMap1)

	res, err = ctx.GetDBList()[0].QuerySQL(tlpQuery)
	if err != nil {
		return
	}
	resMap2 := make(map[string]int)
	for _, rec := range res {
		if rec[0] == nil {
			key = "nil"
		} else {
			key = string(rec[0].([]byte))
		}
		resMap2[key]++
	}
	log.Infof("res: %+v", resMap2)
	if !mapEqual1(resMap1, resMap2) {
		log.Infof("TLP结果不一致，predicate:" + predicate)
	}
}

func mapEqual1(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}
	for k, v := range map1 {
		v2, ok := map2[k]
		if !ok || v != v2 {
			return false
		}
	}
	return true
}
