package model

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

type TestReport struct {
	Rid        int       `json:"Rid" db:"rid"`
	Jid        int       `json:"Jid" db:"jid"`
	InputStmt  string    `json:"InputStmt" db:"input_stmt"`
	InputRes   string    `json:"InputRes" db:"input_res"`
	OracleStmt string    `json:"OracleStmt" db:"oracle_stmt"`
	OracleRes  string    `json:"OracleRes" db:"oracle_res"`
	Category   string    `json:"Category" db:"category"`
	ReportTime time.Time `json:"ReportTime" db:"report_time"`
	State      string    `json:"State" db:"state"`
	URL        string    `json:"URL" db:"url"`
	Comments   string    `json:"Comments" db:"comments"`
	Deleted    int       `json:"Deleted" db:"deleted"`
}

type TestReportPageItem struct {
	Rid        int       `json:"Rid" db:"rid"`
	Jid        int       `json:"Jid" db:"jid"`
	Target     string    `json:"Target" db:"target"`
	Oracle     string    `json:"Oracle" db:"oracle"`
	Category   string    `json:"Category" db:"category"`
	ReportTime time.Time `json:"ReportTime" db:"report_time"`
	State      string    `json:"State" db:"state"`
	URL        string    `json:"URL" db:"url"`
	Comments   string    `json:"Comments" db:"comments"`
	Deleted    int       `json:"Deleted" db:"deleted"`
}

func GetTestReportByJid(jid int) ([]TestReport, error) {
	var reports []TestReport
	sql := fmt.Sprintf("SELECT * FROM %s WHERE jid = %d", tableNameTestReport, jid)
	err := db.Select(&reports, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询TestReports出错: %s\n", err)
		log.Warnf(errMsg)
		return nil, errors.New(errMsg)
	}
	return reports, nil
}

func GetTestReportByRid(rid int) (*TestReport, error) {
	report := TestReport{}
	sql := fmt.Sprintf("SELECT * FROM %s WHERE rid = %d", tableNameTestReport, rid)
	err := db.Get(&report, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询TestReports出错: %s\n", err)
		log.Warnf(errMsg)
		return nil, errors.New(errMsg)
	}
	return &report, nil
}

func GetTestReportCount() (int, error) {
	count := new(int)
	sql := fmt.Sprintf("SELECT count(*) FROM %s WHERE deleted = 0", tableNameTestReport)
	err := db.Get(count, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询所有TestReport数量出错: %s\n", err)
		log.Warnf(errMsg)
		return 0, errors.New(errMsg)
	}
	return *count, nil
}

func GetTestReportPage(offset, limit int) ([]TestReportPageItem, error) {
	var reports []TestReportPageItem
	sql := fmt.Sprintf("select rid,tr.jid,target,oracle,category,"+
		"report_time,tr.state,url,tr.comments,tr.deleted from %s as tr"+
		" join %s as tj on tr.jid=tj.jid WHERE tr.deleted = 0 ORDER BY tr.jid LIMIT %d,%d",
		tableNameTestReport, tableNameTestJob, offset, limit)
	err := db.Select(&reports, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询TstJob出错: %s\n", err)
		log.Warnf(errMsg)
		return nil, errors.New(errMsg)
	}
	return reports, nil
}

func DeleteTestReport(rid int) error {
	sql := fmt.Sprintf("UPDATE %s SET deleted = 1 WHERE rid = %d", tableNameTestReport, rid)
	_, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("删除TestReport失败：%s\n", err)
		log.Warnf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}
