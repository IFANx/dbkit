package internal

import (
	"dbkit/internal/common"
	"dbkit/internal/common/stmt"
	"dbkit/internal/randomly"
	"time"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type TestContext struct {
	TestID       string
	Oracle       string
	Target       common.DBMS
	DBTester     Tester
	Conn         *sqlx.DB
	StartTime    time.Time
	EndTime      time.Time
	SqlCount     int64
	TestRunCount int64
	ReportCount  int64
	Tables       []*Table
}

func NewTestContext() *TestContext {
	state := GetState()
	config := state.Config
	testID := config.Oracle + time.Now().Format("060102150405") + randomly.RandAlphabetStrLen(5)
	target := common.GetDBMSFromStr(config.Target)
	return &TestContext{
		TestID:       testID,
		Oracle:       config.Oracle,
		Target:       target,
		DBTester:     nil,
		Conn:         state.Connections[target],
		StartTime:    time.Time{},
		EndTime:      time.Time{},
		SqlCount:     0,
		TestRunCount: 0,
		ReportCount:  0,
		Tables:       make([]*Table, 0),
	}
}

func (ctx *TestContext) SetTester(tester Tester) {
	ctx.DBTester = tester
}

func (ctx *TestContext) Start() {
	ctx.StartTime = time.Now()
	ctx.DBTester.RunTest()
}

func (ctx *TestContext) End() {
	ctx.EndTime = time.Now()
}

func (ctx *TestContext) CountSQL() {
	ctx.SqlCount++
}

func (ctx *TestContext) CountTestRun() {
	ctx.TestRunCount++
}

func (ctx *TestContext) CountReport() {
	ctx.ReportCount++
}

func (ctx *TestContext) Queryx(query string) (*sqlx.Rows, error) {
	return ctx.Conn.Queryx(query)
}

func (ctx *TestContext) QuerySQL(query string) ([][]interface{}, error) {
	rows, err := ctx.Conn.Queryx(query)
	if err != nil {
		log.Warnf("Fail to query: %s, cause: %s", query, err)
		return nil, err
	}
	defer rows.Close()
	res := make([][]interface{}, 0)
	for rows.Next() {
		cols, err := rows.SliceScan()
		if err != nil {
			log.Warnf("Fail to query: %s, cause: %s", query, err)
			return nil, err
		}
		res = append(res, cols)
	}
	return res, nil
}

func (ctx *TestContext) Query(stmt stmt.SelectStmt) ([][]interface{}, error) {
	var query string
	if ctx.Target == common.TIDB {
		query = stmt.StringInMode()
	} else {
		query = stmt.String()
	}
	return ctx.QuerySQL(query)
}

func (ctx *TestContext) ExecSQLIgnoreRes(sql string) {
	_, err := ctx.Conn.Exec(sql)
	if err != nil {
		log.Warnf("Fail to execute: %s, cause: %s", sql, err)
	}
}

func (ctx *TestContext) ExecSQL(sql string) error {
	_, err := ctx.Conn.Exec(sql)
	if err != nil {
		log.Warnf("Fail to execute: %s, cause: %s", sql, err)
	}
	return err
}

func (ctx *TestContext) ExecSQLAffectedRow(sql string) (int, error) {
	res, err := ctx.Conn.Exec(sql)
	if err != nil {
		log.Warnf("Fail to execute: %s, cause: %s", sql, err)
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Warnf("Fail to execute: %s, cause: %s", sql, err)
		return 0, err
	}
	return int(count), err
}

func (ctx *TestContext) ExecUpdate(stmt stmt.UpdateStmt) (int, error) {
	return ctx.ExecSQLAffectedRow(stmt.String())
}

func (ctx *TestContext) ExecDelete(stmt stmt.DeleteStmt) (int, error) {
	return ctx.ExecSQLAffectedRow(stmt.String())
}

func (ctx *TestContext) ExecInsert(stmt stmt.InsertStmt) error {
	_, err := ctx.Conn.Exec(stmt.String())
	if err != nil {
		log.Warnf("Fail to execute: %s, cause: %s", stmt.String(), err)
	}
	return err
}
