package model

import (
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
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
	err := db.Get(&job, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询TestJob出错: %s\n", err)
		log.Warnf(errMsg)
		return nil, errors.New(errMsg)
	}
	return &job, nil
}

func AddTestJob(dsn, dbName, target, oracle, comments string, timeLimit float32) (int, error) {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("INSERT INTO %s(dsn, db_name, target, "+
		"oracle, state, time_limit, comments, created_at, deleted) "+
		"VALUES('%s', '%s', '%s', '%s', '%d', '%f', '%s', '%s', '%d')",
		tableNameTestJob, dsn, dbName, target, oracle, 1, timeLimit, comments, timeStr, 0)
	res, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("数据库写入TestJob记录失败: %s\n", err)
		log.Warnf(errMsg)
		return 0, errors.New(errMsg)
	}
	jid, err := res.LastInsertId()
	if err != nil {
		errMsg := fmt.Sprintf("数据库获取新插入记录jid失败: %s\n", err)
		log.Warnf(errMsg)
		return 0, errors.New(errMsg)
	}
	return int(jid), nil
}

func DeleteTestJob(jid int) error {
	sql := fmt.Sprintf("UPDATE %s SET deleted = 1 WHERE jid = %d", tableNameTestJob, jid)
	_, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("删除TestJob失败：%s\n", err)
		log.Warnf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func AlterTestJobDBName(jid int, dbName string) error {
	sql := fmt.Sprintf("UPDATE %s SET db_name = '%s' WHERE jid = %d", tableNameTestJob, dbName, jid)
	_, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("更新TestJob的数据库实例名称失败：%s\n", err)
		log.Warnf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func AbortTestJob(jid int) error {
	sql := fmt.Sprintf("UPDATE %s SET state = -1 WHERE jid = %d AND state = 1", tableNameTestJob, jid)
	_, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("更新TestJob状态失败:%s\n", err)
		log.Warnf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func EndTestJob(jid int, success bool) error {
	state := 2
	if !success {
		state = -1
	}
	sql := fmt.Sprintf("UPDATE %s SET state = %d WHERE jid = %d", tableNameTestJob, state, jid)
	_, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("更新TestJob状态失败: %s\n", err)
		log.Warnf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}
