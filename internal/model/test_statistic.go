package model

import (
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
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

func AddStatistic(jid, sqlCount, caseCount, reportCount int, failCause string) (int, error) {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("INSERT INTO %s(jid, sql_count, "+
		"case_count, report_count, fail_cause, end_time) "+
		"VALUES('%d', '%d', '%d', '%d', '%s', '%s')",
		tableNameTestStatistic, jid, sqlCount, caseCount, reportCount, failCause, timeStr)
	res, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("数据库写入TestStatistic记录失败：%s\n", err)
		log.Warnf(errMsg)
		return 0, errors.New(errMsg)
	}
	sid, err := res.LastInsertId()
	if err != nil {
		errMsg := fmt.Sprintf("数据库获取新插入记录sid失败：%s\n", err)
		log.Warnf(errMsg)
		return 0, errors.New(errMsg)
	}
	return int(sid), nil
}
