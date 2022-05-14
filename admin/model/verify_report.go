package model

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

type VerifyReport struct {
	Rid       int       `json:"Rid" db:"rid"`
	Jid       int       `json:"Jid" db:"jid"`
	Pass      int       `json:"Pass" db:"pass"`
	FilePath  string    `json:"FilePath" db:"file_path"`
	Comments  string    `json:"comments" db:"comments"`
	CreatedAt time.Time `json:"CreatedAt" db:"created_at"`
	Deleted   int       `json:"Deleted" db:"deleted"`
}

func GetVerifyReportCount() (int, error) {
	count := new(int)
	sql := fmt.Sprintf("SELECT count(*) FROM %s WHERE deleted = 0", tableNameVerifyReport)
	err := db.Get(count, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询所有VerifyReport数量出错: %s\n", err)
		log.Warnf(errMsg)
		return 0, errors.New(errMsg)
	}
	return *count, nil
}

func GetVerifyReportPage(offset, limit int) ([]TestReport, error) {
	var reports []TestReport
	sql := fmt.Sprintf("SELECT * FROM %s WHERE deleted = 0 ORDER BY jid LIMIT %d,%d",
		tableNameVerifyReport, offset, limit)
	err := db.Select(&reports, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询VerifyReport出错: %s\n", err)
		log.Warnf(errMsg)
		return nil, errors.New(errMsg)
	}
	return reports, nil
}

func GetVerifyReportByRid(rid int) (*VerifyReport, error) {
	report := VerifyReport{}
	sql := fmt.Sprintf("SELECT * FROM %s WHERE rid = %d", tableNameVerifyReport, rid)
	err := db.Get(&report, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询VerifyReports出错: %s\n", err)
		log.Warnf(errMsg)
		return nil, errors.New(errMsg)
	}
	return &report, nil
}

func GetVerifyReportByJid(jid int) (*VerifyReport, error) {
	report := VerifyReport{}
	sql := fmt.Sprintf("SELECT * FROM %s WHERE jid = %d", tableNameVerifyReport, jid)
	err := db.Get(&report, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询VerifyReports出错: %s\n", err)
		log.Warnf(errMsg)
		return nil, errors.New(errMsg)
	}
	return &report, nil
}
