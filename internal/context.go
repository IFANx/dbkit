package internal

import (
	"dbkit/internal/common"
	"fmt"
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

func (ctx *TaskContext) Clean() {

}
