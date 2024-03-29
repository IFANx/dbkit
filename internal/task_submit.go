package internal

import (
	"dbkit/internal/common"
	"dbkit/internal/common/dbms"
	"dbkit/internal/common/oracle"
	_ "dbkit/internal/dameng"
	mysql2 "dbkit/internal/dameng"
	"dbkit/internal/model"
	"dbkit/internal/mysql"
	"errors"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type TaskType int

const (
	TaskTypeTest = iota
	TaskTypeDiff
	TaskTypeVerify
)

type TaskSubmit struct {
	Type        TaskType
	Oracle      oracle.Oracle
	TargetTypes []dbms.DBMS
	ConnList    []*sqlx.DB
	DSNList     []string
	Limit       float32
	Model       string // TODO: Change to enumerate
	Comments    string
}

func BuildTaskFromSubmit(submit *TaskSubmit) (int, error) {
	// log.Printf("%+v\n", submit)
	var (
		jid    int
		runner TaskRunner
		err    error
	)
	runner, err = getTaskRunnerFromSubmit(submit)
	if err != nil {
		return 0, err
	}
	dsnStr := strings.Join(submit.DSNList, ",")
	targetTypeStrList := make([]string, 0)
	for _, tp := range submit.TargetTypes {
		targetTypeStrList = append(targetTypeStrList, tp.Alias)
	}
	targetTypeStr := strings.Join(targetTypeStrList, ",")
	if submit.Type == TaskTypeVerify {
		jid, err = model.AddVerifyJob(dsnStr, "", targetTypeStr, submit.Model, submit.Comments, int(submit.Limit))
		if err != nil {
			return 0, errors.New("创建VerifyJob失败: " + err.Error())
		}
	} else {
		jid, err = model.AddTestJob(dsnStr, "", targetTypeStr, submit.Oracle.Name, submit.Comments, submit.Limit)
		if err != nil {
			return 0, errors.New("创建TestJob失败: " + err.Error())
		}
	}
	task := &TaskContext{
		JobID:        jid,
		Submit:       submit,
		Runner:       runner,
		StartTime:    time.Time{},
		Deadline:     time.Time{},
		EndTime:      time.Time{},
		SqlCount:     0,
		TestRunCount: 0,
		ReportCount:  0,
		DBList:       make([]*common.Database, 0),
		Aborted:      0,
	}
	GetState().SubmitTask(task)
	return jid, nil
}

func getTaskRunnerFromSubmit(submit *TaskSubmit) (TaskRunner, error) {
	if submit.Oracle == oracle.NoREC2 {
		if submit.TargetTypes[0] == dbms.MYSQL {
			return &mysql.MySQLNoREC2{}, nil
		}
	}
	if submit.Oracle == oracle.Troc {
		if submit.TargetTypes[0] == dbms.MYSQL {
			return &mysql.MySQLTrocTester{}, nil
		}
	}
	if submit.Oracle == oracle.TrocPlus {
		if submit.TargetTypes[0] == dbms.MYSQL {
			return &mysql.MySQLTrocPlusTester{}, nil
		}
	}
	if submit.Oracle == oracle.TLP {
		if submit.TargetTypes[0] == dbms.MYSQL {
			return &mysql.MySQLTLPTester{}, nil
		}
		if submit.TargetTypes[0] == dbms.DAMENG {
			return &mysql2.DAMENGTLPTester{}, nil
		}
	}
	if submit.Oracle == oracle.NoREC {
		if submit.TargetTypes[0] == dbms.MYSQL {
			return &mysql.MySQLNoRECTester{}, nil
		}
	}
	if submit.Oracle == oracle.PQS {
		if submit.TargetTypes[0] == dbms.MYSQL {
			return &mysql.MySQLPQSTester{}, nil
		}
	}
	return nil, errors.New("该测试功能未实现")
}
