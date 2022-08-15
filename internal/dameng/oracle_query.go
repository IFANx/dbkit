package dameng

import (
	"dbkit/internal/common"
	"dbkit/internal/common/stmt"
	"dbkit/internal/mysql/gen"
	"dbkit/internal/randomly"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type MySQLQueryTester struct{}

func (tester *MySQLQueryTester) RunTask(ctx common.OracleRuntime) {
	table := &common.Table{
		DB:   ctx.GetDBList()[0],
		Name: "t",
	}
	for {
		table.Build()

		for run := 0; run < 20; run++ {
			ctx.IncrTestRunCount(1)
			predicate := gen.GenPredicate(table)
			log.Infof("生成新的谓词：%s", predicate)
			NoRECWithCtx(ctx, table, predicate)
			TLPWithCtx(ctx, table, predicate)
		}
	}
}

func NoRECWithCtx(ctx common.OracleRuntime, table *common.Table, predicate string) {
	norec := stmt.SelectStmt{
		TableName: table.Name,
		Targets:   []string{"count(1)"},
		Predicate: predicate,
		ForShare:  false,
		ForUpdate: false,
	}
	res, err := ctx.GetDBList()[0].Query(norec)
	if err != nil || len(res) < 1 {
		return
	}
	count1Str := string(res[0][0].([]byte))
	count1, err := strconv.Atoi(count1Str)
	if err != nil {
		return
	}
	norec.Predicate = "True"
	norec.Targets[0] = "((" + predicate + ") IS TRUE)"
	res, err = ctx.GetDBList()[0].Query(norec)
	if err != nil {
		return
	}
	count2 := 0
	for _, row := range res {
		tmpStr := string(row[0].([]byte))
		tmp, err := strconv.Atoi(tmpStr)
		if err != nil {
			continue
		}
		count2 += tmp
	}
	log.Infof("count1=%d, count2=%d", count1, count2)
	if count2 != count1 {
		log.Errorf("Potential bugs found by:" + predicate)
	}
}

func TLPWithCtx(ctx common.OracleRuntime, table *common.Table, predicate string) {
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

	if !mapEqual(resMap1, resMap2) {
		log.Infof("TLP结果不一致, predicate:" + predicate)
	}
}

func mapEqual(map1, map2 map[string]int) bool {
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
