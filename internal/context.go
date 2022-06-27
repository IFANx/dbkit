package internal

import (
	"dbkit/internal/common"
	"time"
)

type TaskContext struct {
	JobID        int
	Submit       *TaskSubmit
	Runner       TaskRunner
	StartTime    time.Time
	Deadline     time.Time
	EndTime      time.Time
	SqlCount     int64
	TestRunCount int64
	ReportCount  int64
	DBList       []*common.Database
	Aborted      int32
}

func (ctx *TaskContext) Start() {
	defer func() {
		if err := recover(); err != nil {
			ctx.EndTime = time.Now()
		} else {
			ctx.EndTime = time.Now()
		}
		ctx.Clean()
	}()
	ctx.StartTime = time.Now()
	ctx.Runner.RunTask(ctx)
}

func (ctx *TaskContext) Clean() {

}
