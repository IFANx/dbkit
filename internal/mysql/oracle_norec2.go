package mysql

import (
	"dbkit/internal/common"
	"dbkit/internal/common/stmt"
	"dbkit/internal/mysql/gen"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type MySQLNoREC2 struct{}

func (tester *MySQLNoREC2) RunTask(ctx common.OracleRuntime) {
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
			NoREC2WithCtx(ctx, table, predicate)
			if ctx.IsAborted() {
				break
			}
		}

		if ctx.IsAborted() {
			break
		}
	}
}

func NoREC2WithCtx(ctx common.OracleRuntime, table *common.Table, predicate string) {
	query1 := stmt.SelectStmt{
		TableName: table.Name,
		Targets:   []string{"count(1)"},
		Predicate: predicate,
		ForShare:  false,
		ForUpdate: false,
	}
	res, err := ctx.GetDBList()[0].Query(query1)
	if err != nil || len(res) < 1 {
		return
	}
	count1Str := string(res[0][0].([]byte))
	count1, err := strconv.Atoi(count1Str)
	if err != nil {
		return
	}

	alter1 := "ALTER TABLE " + table.Name + " ADD COLUMN cg DECIMAL(10,6) AS (" + predicate + ")" // TODO: change into statement generation
	ctx.GetDBList()[0].ExecSQL(alter1)

	query1.Predicate = "cg"
	res, err = ctx.GetDBList()[0].Query(query1)
	if err != nil {
		return
	}

	dropCg := "ALTER TABLE " + table.Name + " DROP COLUMN cg" //not good
	ctx.GetDBList()[0].ExecSQL(dropCg)

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
