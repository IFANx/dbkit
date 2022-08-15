package dameng

import (
	"dbkit/internal/common"
	"dbkit/internal/common/stmt"
	"dbkit/internal/mysql/gen"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type MySQLNoRECTester struct {
}

func (tester *MySQLNoRECTester) RunTask(ctx common.OracleRuntime) {
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
			log.Infof("生成新的谓词：%s", predicate)
			NoRECWithCtx1(ctx, table, predicate)
			if ctx.IsAborted() {
				break
			}
		}
		if ctx.IsAborted() {
			break
		}
	}
}

func NoRECWithCtx1(ctx common.OracleRuntime, table *common.Table, predicate string) {
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
	norec.Targets[0] = "((" + predicate + ") is TRUE)"
	res, err = ctx.GetDBList()[0].Query(norec)
	if err != nil {
		return
	}
	count2 := 0
	for _, row := range res {
		tmpStr := string(row[0].([]byte))
		tmp, err := strconv.Atoi(tmpStr)
		if err != nil {
			return
		}
		count2 += tmp
	}
	log.Infof("count1=%d,count2=%d", count1, count2)
	if count1 != count2 {
		log.Errorf("Potential bugs found by:" + predicate)
	}
}
