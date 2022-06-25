package model

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

type TestStatistic struct {
	Sid         int       `json:"Sid" db:"sid"`
	Jid         int       `json:"Jid" db:"jid"`
	SQLCount    int       `json:"SQLCount" db:"sql_count"`
	CaseCount   int       `json:"CaseCount" db:"case_count"`
	ReportCount int       `json:"ReportCount" db:"report_count"`
	FailCause   string    `json:"FailCause" db:"fail_cause"`
	EndTime     time.Time `json:"EndTime" db:"end_time"`
}

func GetStatisticByJid(jid int) (*TestStatistic, error) {
	var stat = TestStatistic{}
	sql := fmt.Sprintf("SELECT * FROM %s WHERE jid = %d", tableNameTestStatistic, jid)
	err := db.Get(&stat, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询TestStatistic出错: %s\n", err)
		log.Warnf(errMsg)
		return nil, errors.New(errMsg)
	}
	return &stat, nil
}
