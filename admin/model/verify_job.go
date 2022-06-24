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
	DBName    string    `json:"DBName" db:"db_name"`
	Target    string    `json:"Target" db:"target"`
	Model     string    `json:"Model" db:"model"`
	Op        int       `json:"Op" db:"op"`
	State     int       `json:"State" db:"state"`
	Comments  string    `json:"comments" db:"comments"`
	CreatedAt time.Time `json:"CreatedAt" db:"created_at"`
	Deleted   int       `json:"Deleted" db:"deleted"`
}

type VerifyJobPageItem struct {
	Jid       int       `json:"Jid" db:"jid"`
	Target    string    `json:"Target" db:"target"`
	Model     string    `json:"Model" db:"model"`
	Op        int       `json:"Op" db:"op"`
	State     int       `json:"State" db:"state"`
	Rid       int       `json:"Rid" db:"rid"`
	Pass      int       `json:"Pass" db:"pass"`
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

func GetVerifyJobPage(offset, limit int) ([]VerifyJobPageItem, error) {
	var jobs []VerifyJobPageItem
	sql := fmt.Sprintf("select vj.jid,target,model,state, IFNULL(rid, 0) as rid, IFNULL(pass, 0) as pass,"+
		"vj.op, vj.comments, vj.created_at,vj.deleted from %s as vj left outer join %s as vr on vj.jid=vr.jid "+
		"WHERE vj.deleted = 0 ORDER BY vj.jid LIMIT %d,%d",
		tableNameVerifyJob, tableNameVerifyReport, offset, limit)
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

func DeleteVerifyJob(jid int) error {
	sql := fmt.Sprintf("UPDATE %s SET deleted = 1 WHERE jid = %d", tableNameVerifyJob, jid)
	_, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("删除VerifyJob失败：%s\n", err)
		log.Warnf(errMsg)
		return errors.New(errMsg)
	}
	return nil
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
