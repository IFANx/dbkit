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

func AddVerifyJob(dsn, dbName, target, model, comments string, op int) (int, error) {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("INSERT INTO %s(dsn, db_name, target, "+
		"model, op, state, comments, created_at, deleted) "+
		"VALUES('%s', '%s', '%s', '%s', '%d', '%d', '%s', '%s', '%d')",
		tableNameVerifyJob, dsn, dbName, target, model, op, 1, comments, timeStr, 0)
	res, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("数据库写入VerifyJob记录失败：%s\n", err)
		log.Warnf(errMsg)
		return 0, errors.New(errMsg)
	}
	jid, err := res.LastInsertId()
	if err != nil {
		errMsg := fmt.Sprintf("数据库获取新插入记录jid失败：%s\n", err)
		log.Warnf(errMsg)
		return 0, errors.New(errMsg)
	}
	return int(jid), nil
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

func AlterVerifyJobDBName(jid int, dbName string) error {
	sql := fmt.Sprintf("UPDATE %s SET db_name = %s WHERE jid = %d", tableNameVerifyJob, dbName, jid)
	_, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("更新VerifyJob的数据库实例名称失败：%s\n", err)
		log.Warnf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func EndVerifyJob(jid int, success bool) error {
	state := 2
	if !success {
		state = -1
	}
	sql := fmt.Sprintf("UPDATE %s SET state = %d WHERE jid = %d", tableNameVerifyJob, state, jid)
	_, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("更新TestJob状态失败：%s\n", err)
		log.Warnf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}
