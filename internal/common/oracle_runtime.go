package common

import (
	"dbkit/internal/common/dbms"
	"dbkit/internal/common/oracle"
	"time"
)

type OracleRuntime interface {
	GetJobID() int
	GetOracleList() []oracle.Oracle
	GetTargetTypes() []dbms.DBMS
	GetLimit() float32
	GetModel() string
	GetComments() string
	GetStartTime() time.Time
	GetDeadline() time.Time
	GetEndTime() time.Time
	GetSqlCount() int
	GetTestRunCount() int
	GetReportCount() int
	GetDBList() []*Database
	IsAborted() bool
	IsFinished() bool

	SetEndTime(time.Time)
	IncrSqlCount(int)
	IncrTestRunCount(int)
	IncrReportCount(int)
}
