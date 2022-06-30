package internal

import (
	"dbkit/internal/common"
	"dbkit/internal/common/dbms"
	"dbkit/internal/common/oracle"
	"fmt"
	"sync/atomic"
	"time"
)

type TaskContext struct {
	JobID        int
	Submit       *TaskSubmit
	Runner       TaskRunner
	StartTime    time.Time
	Deadline     time.Time
	EndTime      time.Time
	SqlCount     int
	TestRunCount int
	ReportCount  int
	DBList       []*common.Database
	Aborted      int32
	Finished     int32
}

func (ctx *TaskContext) Start() {
	defer func() {
		if err := recover(); err != nil {
			ctx.EndTime = time.Now()
			if !ctx.IsAborted() {

			}
		} else {
			ctx.EndTime = time.Now()
			if !ctx.IsAborted() {

			}
		}
		ctx.Clean()
	}()
	ctx.StartTime = time.Now()
	ctx.initDBList()
	ctx.Runner.RunTask(ctx)
}

func (ctx *TaskContext) initDBList() {
	var dbName string
	switch ctx.Submit.Type {
	case TaskTypeTest, TaskTypeDiff:
		dbName = fmt.Sprintf("test%d", ctx.JobID)
	case TaskTypeVerify:
		dbName = fmt.Sprintf("verify%d", ctx.JobID)
	}
	n := len(ctx.Submit.ConnList)
	dbList := make([]*common.Database, n)
	for i := 0; i < n; i++ {
		db := &common.Database{
			DBMS:   ctx.Submit.TargetTypes[i],
			DBName: dbName,
			Conn:   ctx.Submit.ConnList[i],
			Tables: make([]*common.Table, 0),
		}
		err := db.Refresh()
		if err != nil {
			panic("初始化测试数据库实例失败：" + err.Error())
		}
		dbList[i] = db
	}
}

func (ctx *TaskContext) Abort() {
	atomic.StoreInt32(&ctx.Aborted, 1)
}

func (ctx *TaskContext) IsAborted() bool {
	return atomic.LoadInt32(&ctx.Aborted) == 1
}

func (ctx *TaskContext) Finish() {
	atomic.StoreInt32(&ctx.Finished, 1)
}

func (ctx *TaskContext) IsFinished() bool {
	return atomic.LoadInt32(&ctx.Finished) == 1
}

func (ctx *TaskContext) Clean() {

}

func (ctx *TaskContext) GetJobID() int {
	return ctx.JobID
}

func (ctx *TaskContext) GetOracleList() []oracle.Oracle {
	return ctx.Submit.OracleList
}

func (ctx *TaskContext) GetTargetTypes() []dbms.DBMS {
	return ctx.GetTargetTypes()
}

func (ctx *TaskContext) GetLimit() float32 {
	return ctx.Submit.Limit
}

func (ctx *TaskContext) GetModel() string {
	return ctx.Submit.Model
}

func (ctx *TaskContext) GetComments() string {
	return ctx.Submit.Comments
}

func (ctx *TaskContext) GetStartTime() time.Time {
	return ctx.StartTime
}

func (ctx *TaskContext) GetDeadline() time.Time {
	return ctx.Deadline
}

func (ctx *TaskContext) GetEndTime() time.Time {
	return ctx.EndTime
}

func (ctx *TaskContext) GetSqlCount() int {
	return ctx.SqlCount
}

func (ctx *TaskContext) GetTestRunCount() int {
	return ctx.TestRunCount
}

func (ctx *TaskContext) GetReportCount() int {
	return ctx.ReportCount
}

func (ctx *TaskContext) GetDBList() []*common.Database {
	return ctx.DBList
}

func (ctx *TaskContext) SetEndTime(t time.Time) {
	ctx.EndTime = t
}

func (ctx *TaskContext) IncrSqlCount(x int) {
	ctx.SqlCount += x
}

func (ctx *TaskContext) IncrTestRunCount(x int) {
	ctx.TestRunCount += x
}

func (ctx *TaskContext) IncrReportCount(x int) {
	ctx.ReportCount += x
}
