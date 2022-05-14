package model

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

type VerifyJob struct {
	Jid       int       `json:"Jid" db:"jid"`
	DSN       string    `json:"DSN" db:"dsn"`
	DBName    string    `json:"DBName" db:"dbname"`
	Target    string    `json:"Target" db:"target"`
	Op        int       `json:"Op" db:"op"`
	State     int       `json:"State" db:"state"`
	Comments  string    `json:"comments" db:"comments"`
	CreatedAt time.Time `json:"CreatedAt" db:"created_at"`
	Deleted   int       `json:"Deleted" db:"deleted"`
}

var RunningVerifyJobs map[int]*VerifyJob

func GetVerifyJobCount() (int, error) {
	count := new(int)
	sql := fmt.Sprintf("SELECT count(*) FROM %s WHERE deleted = 0", tableNameVerifyJob)
	err := db.Get(count, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询所有VerifyJob数量出错: %s\n", err)
		log.Warnf(errMsg)
		return 0, errors.New(errMsg)
	}
	return *count, nil
}

func GetVerifyJobPage(offset, limit int) ([]VerifyJob, error) {
	var jobs []VerifyJob
	sql := fmt.Sprintf("SELECT * FROM %s WHERE deleted = 0 ORDER BY jid LIMIT %d,%d",
		tableNameVerifyJob, offset, limit)
	err := db.Select(&jobs, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询VerifyJob出错: %s\n", err)
		log.Warnf(errMsg)
		return nil, errors.New(errMsg)
	}
	return jobs, nil
}

func GetVerifyJobByJid(jid int) (*VerifyJob, error) {
	job := VerifyJob{}
	sql := fmt.Sprintf("SELECT * FROM %s WHERE jid = %d", tableNameVerifyJob, jid)
	err := db.Get(&job, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询VerifyJob出错: %s\n", err)
		log.Warnf(errMsg)
		return nil, errors.New(errMsg)
	}
	return &job, nil
}

func AbortVerifyJob(jid int) error {
	sql := fmt.Sprintf("UPDATE %s SET state = -1 WHERE jid = %d AND state = 1", tableNameVerifyJob, jid)
	_, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("更新VerifyJob状态失败：%s\n", err)
		log.Warnf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}
