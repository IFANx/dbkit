package dameng

import (
	"dbkit/internal/common"
	"dbkit/internal/common/stmt"
	"dbkit/internal/mysql/gen"
	"dbkit/internal/randomly"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

type MySQLPQSTester struct {
}

func (tester *MySQLPQSTester) RunTask(ctx common.OracleRuntime) {
	table := &common.Table{
		DB:   ctx.GetDBList()[0],
		Name: "t",
	}
	pqsOrigin := stmt.SelectStmt{
		TableName: table.Name,
		Targets:   []string{"*"},
		Predicate: "True",
		ForShare:  false,
		ForUpdate: false,
	}

	for {
		table.Build()

		columns := table.ColumnNames
		n := len(columns)
		//查询出当前数据表中的所有数据
		rows, err := ctx.GetDBList()[0].Query(pqsOrigin)
		if err != nil {
			return
		}

		//创建一个map数组来存储所有查询出的数据
		columnStructs := make([]map[string]interface{}, 0)
		for i := 0; i < len(rows); i++ {
			// 创建一个map来存储一行数据
			columnStruct := make(map[string]interface{})
			for j := 0; j < n; j++ {
				col := rows[i]
				colType := table.Columns[columns[j]].Type.Name()
				fmt.Println(colType)
				columnStruct[columns[j]] = col[j]
			}
			columnStructs = append(columnStructs, columnStruct)
		}
		//随机选择一行数据作为pivotRow
		pivotRow := columnStructs[randomly.RandIntGap(0, n)]
		log.Infof("the chosen pivotrow is：%s", pivotRow)

		//为pivotRow构造20条谓词逻辑为true的where子句
		for run := 0; run < 20; run++ {
			ctx.IncrTestRunCount(1)
			time.Sleep(time.Second * 5)
			RectifiedPredicate, predicate := gen.GenPQS(table, pivotRow)
			//新建用于查询的语句，其值根据传入的predicate的改变
			//计算生成的predicate语句的取值
			expectedValue := getExpectedValue(predicate)
			if expectedValue == nil {
				RectifiedPredicate = "(" + RectifiedPredicate + ") IS NULL"
			}
			if expectedValue == false {
				RectifiedPredicate = "NOT (" + RectifiedPredicate + ")"
			}
			log.Infof("生成新的谓词：%s", RectifiedPredicate)
			log.Infof("生成新的谓词：%s", predicate)
			PQSWithCtx(ctx, table, RectifiedPredicate, pivotRow)
			if ctx.IsAborted() {
				break
			}
		}
		if ctx.IsAborted() {
			break
		}
	}
}

func getExpectedValue(predicate string) interface{} {
	expected := []string{"True", "False", "Null"}
	res := randomly.RandPickOneStr(expected)
	switch res {
	case "True":
		return true
	case "False":
		return false
	case "Null":
		return nil
	default:
		return false
	}
}

func PQSWithCtx(ctx common.OracleRuntime, table *common.Table, predicate string, pivotRow map[string]interface{}) {
	pqs := stmt.SelectStmt{
		TableName: table.Name,
		Targets:   []string{"*"},
		Predicate: predicate,
		ForShare:  false,
		ForUpdate: false,
	}
	columns := table.ColumnNames
	n := len(columns)
	rows, err := ctx.GetDBList()[0].Query(pqs)
	if err != nil {
		return
	}
	//如果当前查询到的结果集长度小于1，说明pivotRow未查询出来，若在构造谓词逻辑正确的情况下，则一定有错误发生
	if len(rows) < 1 {
		log.Infof("Potenial bugs found by " + predicate)
	}
	//	根据查询结果，查看选区的PivotRow是否在结果集中
	//创建一个map数组来存储所有查询出的数据
	columnStructs := make([]map[string]interface{}, 0)
	for i := 0; i < len(rows); i++ {
		// 创建一个map来存储一行数据
		columnStruct := make(map[string]interface{})
		for j := 0; j < n; j++ {
			col := rows[i]
			columnStruct[columns[j]] = col[j]
		}
		columnStructs = append(columnStructs, columnStruct)
	}
	if !IsPivotRowExist(table, pivotRow, columnStructs) {
		log.Infof("Potenial bugs found by " + predicate)
	}
	log.Infof("Next process")
}

func IsPivotRowExist(table *common.Table, pivotRow map[string]interface{}, columnStructs []map[string]interface{}) bool {
	for i := 0; i < len(columnStructs); i++ {
		if !mapEqual2(pivotRow, columnStructs[i]) {
			continue
		}
		return true
	}
	return false
}
func mapEqual2(map1, map2 map[string]interface{}) bool {
	var v1 string
	var v2 string
	for k, v := range map1 {
		m1, ok := v.([]byte)
		if ok {
			v1 = string(m1)
		}
		m2, ok := map2[k]
		if ok {
			m3, ok := m2.([]byte)
			if ok {
				v2 = string(m3)
			}
		}
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}
