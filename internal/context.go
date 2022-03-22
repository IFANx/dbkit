package internal

import "time"

type TestContext struct {
	TestID       int64
	Targets      []DBMS
	StartTime    time.Time
	EndTime      time.Time
	SqlCount     int64
	TestRunCount int64
	ReportCount  int64
}
