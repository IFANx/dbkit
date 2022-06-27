package internal

import (
	"dbkit/internal/common"
	"dbkit/internal/common/dbms"
	"dbkit/internal/common/stmt"
	"time"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type TaskContext struct {
	JobID        int
	Submit       *TaskSubmit
	Conn         *sqlx.DB
	StartTime    time.Time
	Deadline     time.Time
	EndTime      time.Time
	SqlCount     int64
	TestRunCount int64
	ReportCount  int64
	DBList       []*common.Database
}

func (ctx *TaskContext) Start() {
	ctx.StartTime = time.Now()

}

func (ctx *TaskContext) End() {
	ctx.EndTime = time.Now()
}

func (ctx *TaskContext) CountSQL() {
	ctx.SqlCount++
}

func (ctx *TaskContext) CountTestRun() {
	ctx.TestRunCount++
}

func (ctx *TaskContext) CountReport() {
	ctx.ReportCount++
}

func (ctx *TaskContext) Queryx(query string) (*sqlx.Rows, error) {
	return ctx.Conn.Queryx(query)
}

func (ctx *TaskContext) QuerySQL(query string) ([][]interface{}, error) {
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

func (ctx *TaskContext) Query(stmt stmt.SelectStmt) ([][]interface{}, error) {
	var query string
	if ctx.Submit.TargetTypes[0] == dbms.TIDB {
		query = stmt.StringInMode()
	} else {
		query = stmt.String()
	}
	return ctx.QuerySQL(query)
}

func (ctx *TaskContext) ExecSQLIgnoreRes(sql string) {
	_, err := ctx.Conn.Exec(sql)
	if err != nil {
		log.Warnf("Fail to execute: %s, cause: %s", sql, err)
	}
}

func (ctx *TaskContext) ExecSQL(sql string) error {
	_, err := ctx.Conn.Exec(sql)
	if err != nil {
		log.Warnf("Fail to execute: %s, cause: %s", sql, err)
	}
	return err
}

func (ctx *TaskContext) ExecSQLAffectedRow(sql string) (int, error) {
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

func (ctx *TaskContext) ExecUpdate(stmt stmt.UpdateStmt) (int, error) {
	return ctx.ExecSQLAffectedRow(stmt.String())
}

func (ctx *TaskContext) ExecDelete(stmt stmt.DeleteStmt) (int, error) {
	return ctx.ExecSQLAffectedRow(stmt.String())
}

func (ctx *TaskContext) ExecInsert(stmt stmt.InsertStmt) error {
	_, err := ctx.Conn.Exec(stmt.String())
	if err != nil {
		log.Warnf("Fail to execute: %s, cause: %s", stmt.String(), err)
	}
	return err
}
