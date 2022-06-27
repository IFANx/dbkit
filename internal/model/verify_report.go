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

func AddVerifyReport(jid, pass int, filePath, comments string) (int, error) {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("INSERT INTO %s(jid, pass, file_path, comments, created_at, deleted) "+
		"VALUES('%d', '%d', '%s', '%s', '%s', '%d')",
		tableNameVerifyReport, jid, pass, filePath, comments, timeStr, 0)
	res, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("数据库写入VerifyReport记录失败：%s\n", err)
		log.Warnf(errMsg)
		return 0, errors.New(errMsg)
	}
	rid, err := res.LastInsertId()
	if err != nil {
		errMsg := fmt.Sprintf("数据库获取新插入记录rid失败：%s\n", err)
		log.Warnf(errMsg)
		return 0, errors.New(errMsg)
	}
	return int(rid), nil
}

func DeleteVerifyReport(rid int) error {
	sql := fmt.Sprintf("UPDATE %s SET deleted = 1 WHERE rid = %d", tableNameVerifyReport, rid)
	_, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("删除VerifyReport失败：%s\n", err)
		log.Warnf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}
