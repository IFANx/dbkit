package internal

import (
	"dbkit/internal/common"
	"dbkit/internal/common/dbms"
	"dbkit/internal/common/oracle"
	"dbkit/internal/model"
	"errors"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type TaskType int

const (
	TaskTypeTest = iota
	TaskTypeDiff
	TaskTypeVerify
)

type TaskSubmit struct {
	Type        TaskType
	OracleList  []oracle.Oracle
	TargetTypes []dbms.DBMS
	ConnList    []*sqlx.DB
	DSNList     []string
	Limit       float32
	Model       string // TODO: Change to enumerate
	Comments    string
}

func BuildTaskFromSubmit(submit *TaskSubmit) (int, error) {
	// log.Printf("%+v\n", submit)
	dsnStr := strings.Join(submit.DSNList, ",")
	targetTypeStrList := make([]string, len(submit.TargetTypes))
	for _, tp := range submit.TargetTypes {
		targetTypeStrList = append(targetTypeStrList, tp.Alias)
	}
	targetTypeStr := strings.Join(targetTypeStrList, ",")
	oracleStrList := make([]string, len(submit.OracleList))
	for _, oc := range submit.OracleList {
		oracleStrList = append(oracleStrList, oc.Alias)
	}
	oracleStr := strings.Join(oracleStrList, ",")
	var (
		jid int
		err error
	)
	if submit.Type == TaskTypeVerify {
		jid, err = model.AddVerifyJob(dsnStr, "", targetTypeStr, submit.Model, submit.Comments, int(submit.Limit))
		if err != nil {
			return 0, errors.New("创建VerifyJob失败：" + err.Error())
		}
	} else {
		jid, err = model.AddTestJob(dsnStr, "", targetTypeStr, oracleStr, submit.Comments, submit.Limit)
		if err != nil {
			return 0, errors.New("创建TestJob失败：" + err.Error())
		}
	}
	task := &TaskContext{
		JobID:        jid,
		Submit:       submit,
		Conn:         submit.ConnList[0],
		StartTime:    time.Time{},
		Deadline:     time.Time{},
		EndTime:      time.Time{},
		SqlCount:     0,
		TestRunCount: 0,
		ReportCount:  0,
		DBList:       make([]*common.Database, 0),
	}
	GetState().submitTask(task)
	return jid, nil
}
