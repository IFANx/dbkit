package model

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

type TestJob struct {
	Jid       int       `json:"Jid" db:"jid"`
	DSN       string    `json:"DSN" db:"dsn"`
	DBName    string    `json:"DBName" db:"db_name"`
	Target    string    `json:"Target" db:"target"`
	Oracle    string    `json:"Oracle" db:"oracle"`
	State     int       `json:"State" db:"state"`
	TimeLimit float32   `json:"TimeLimit" db:"time_limit"`
	Comments  string    `json:"Comments" db:"comments"`
	CreatedAt time.Time `json:"CreatedAt" db:"created_at"`
	Deleted   int       `json:"Deleted" db:"deleted"`
}

var RunningTestJobs map[int]*TestJob

func GetTestJobCount() (int, error) {
	count := new(int)
	sql := fmt.Sprintf("SELECT count(*) FROM %s WHERE deleted = 0", tableNameTestJob)
	err := db.Get(count, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询所有TestJob数量出错: %s\n", err)
		log.Warnf(errMsg)
		return 0, errors.New(errMsg)
	}
	return *count, nil
}

func GetTestJobPage(offset, limit int) ([]TestJob, error) {
	var jobs []TestJob
	sql := fmt.Sprintf("SELECT * FROM %s WHERE deleted = 0 ORDER BY jid LIMIT %d,%d",
		tableNameTestJob, offset, limit)
	err := db.Select(&jobs, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询TestJob出错: %s\n", err)
		log.Warnf(errMsg)
		return nil, errors.New(errMsg)
	}
	return jobs, nil
}

func GetTestJobByJid(jid int) (*TestJob, error) {
	job := TestJob{}
	sql := fmt.Sprintf("SELECT * FROM %s WHERE jid = %d", tableNameTestJob, jid)
	err := db.Select(&job, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询TestJob出错: %s\n", err)
		log.Warnf(errMsg)
		return nil, errors.New(errMsg)
	}
	return &job, nil
}

func AbortTestJob(jid int) error {
	sql := fmt.Sprintf("UPDATE %s SET state = -1 WHERE jid = %d AND state = 1", tableNameTestJob, jid)
	_, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("更新TestJob状态失败：%s\n", err)
		log.Warnf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}
