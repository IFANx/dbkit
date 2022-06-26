package model

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type TargetDSN struct {
	Tid     int    `json:"Tid" db:"tid"`
	DBType  string `json:"DBType" db:"db_type"`
	DBHost  string `json:"DBHost" db:"db_host"`
	DBPort  int    `json:"DBPort" db:"db_port"`
	DBUser  string `json:"DBUser" db:"db_user"`
	DBPwd   string `json:"DBPwd" db:"db_pwd"`
	DBName  string `json:"DBName" db:"db_name"`
	Params  string `json:"Params" db:"params"`
	State   int    `json:"State" db:"state"`
	Version string `json:"Version" db:"version"`
	Deleted string `json:"Deleted" db:"deleted"`
}

func (targetDSN *TargetDSN) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", targetDSN.DBUser, targetDSN.DBPwd,
		targetDSN.DBHost, targetDSN.DBPort, targetDSN.DBName, targetDSN.Params)
}

func (targetDSN *TargetDSN) GetConn() (*sqlx.DB, error) {
	dbType := targetDSN.DBType
	if dbType == "tidb" || dbType == "mariadb" {
		dbType = "mysql"
	}
	conn, err := sqlx.Open(dbType, targetDSN.GetDSN())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func GetAllTargetDSN() ([]TargetDSN, error) {
	var dsnList []TargetDSN
	sql := fmt.Sprintf("SELECT * FROM %s WHERE deleted = 0", tableNameTargetDSN)
	err := db.Select(&dsnList, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询连接参数出错: %s\n", err)
		log.Warnf(errMsg)
		return nil, errors.New(errMsg)
	}
	return dsnList, nil
}

func GetTargetDSNByType(tp string) ([]TargetDSN, error) {
	var dsnList []TargetDSN
	sql := fmt.Sprintf("SELECT * FROM %s WHERE db_type = '%s' AND deleted = 0", tableNameTargetDSN, tp)
	err := db.Select(&dsnList, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询连接参数出错: %s\n", err)
		log.Warnf(errMsg)
		return nil, errors.New(errMsg)
	}
	return dsnList, nil
}

func GetAllAvailableDSN() ([]TargetDSN, error) {
	var dsnList []TargetDSN
	sql := fmt.Sprintf("SELECT * FROM %s WHERE state = 1 AND deleted = 0", tableNameTargetDSN)
	err := db.Select(&dsnList, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询可用连接参数出错: %s\n", err)
		log.Warnf(errMsg)
		return nil, errors.New(errMsg)
	}
	return dsnList, nil
}

func GetAvailableDSNByType(tp string) ([]TargetDSN, error) {
	var dsnList []TargetDSN
	sql := fmt.Sprintf("SELECT * FROM %s WHERE db_type = '%s' AND state = 1 AND deleted = 0",
		tableNameTargetDSN, tp)
	err := db.Select(&dsnList, sql)
	if err != nil {
		errMsg := fmt.Sprintf("查询可用连接参数出错: %s\n", err)
		log.Warnf(errMsg)
		return nil, errors.New(errMsg)
	}
	return dsnList, nil
}

func GetAllAvailableDSNVersion() (map[string]map[string]string, error) {
	typeList := []string{"mysql", "tidb", "mariadb", "sqlite", "postgresql", "cockroachdb"}
	res := make(map[string]map[string]string)
	for _, tp := range typeList {
		verDsnMap, err := GetAvailableDSNVersionByType(tp)
		if err != nil {
			res[tp] = make(map[string]string)
		}
		res[tp] = verDsnMap
	}
	return res, nil
}

func GetAvailableTypeVersionMap() (map[string][]string, error) {
	dsnList, err := GetAllAvailableDSN()
	if err != nil {
		return nil, err
	}
	typeList := []string{"mysql", "tidb", "mariadb", "sqlite", "postgresql", "cockroachdb"}
	verMap := make(map[string][]string)
	for _, tp := range typeList {
		verMap[tp] = make([]string, 0)
	}
	for _, dsn := range dsnList {
		version := dsn.Version
		verMap[dsn.DBType] = append(verMap[dsn.DBType], version)
	}
	return verMap, nil
}

func GetAvailableDSNVersionByType(tp string) (map[string]string, error) {
	dsnList, err := GetAvailableDSNByType(tp)
	if err != nil {
		return nil, err
	}
	verDsnMap := make(map[string]string)
	for _, dsn := range dsnList {
		version := dsn.Version
		dsnStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
			dsn.DBUser, dsn.DBPwd, dsn.DBHost, dsn.DBPort, dsn.DBName, dsn.Params)
		verDsnMap[version] = dsnStr
	}
	return verDsnMap, nil
}

func AddTargetDSN(tp, host, user, pwd, dbName, params string, port int) (int, error) {
	sql := fmt.Sprintf("INSERT INTO %s(db_type, db_host, db_port, db_user, db_pwd, db_name, params) "+
		"VALUES ('%s', '%s', '%d', '%s', '%s', '%s', '%s')", tableNameTargetDSN, tp, host, port, user, pwd, dbName, params)
	execRes, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("新增连接参数出错: %s\n", err)
		log.Warnf(errMsg)
		return 0, errors.New(errMsg)
	}
	tid, err := execRes.LastInsertId()
	return int(tid), err
}

func DeleteTargetDSN(tid int) error {
	sql := fmt.Sprintf("UPDATE %s SET deleted = 1 WHERE tid = %d", tableNameTargetDSN, tid)
	_, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("删除连接参数出错: %s\n", err)
		log.Warnf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func GetTargetDSNVersion(tp, host, user, pwd, dbName, params string, port int) (string, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", user, pwd, host, port, dbName, params)
	if tp == "tidb" || tp == "mariadb" {
		tp = "mysql"
	}
	tmpDB, err := sqlx.Open(tp, dsn)
	if err != nil {
		return "", err
	}
	var version string
	err = tmpDB.Get(&version, "SELECT VERSION()")
	if err != nil {
		return "", err
	}
	return version, nil
}

func GetTargetDSNByTid(tid int) (*TargetDSN, error) {
	var dsn = TargetDSN{}
	sql := fmt.Sprintf("SELECT * FROM %s WHERE tid = %d", tableNameTargetDSN, tid)
	err := db.Get(&dsn, sql)
	if err != nil {
		return nil, err
	}
	return &dsn, nil
}

func UpdateStateAndVersionByTid(tid, state int, version string) error {
	sql := fmt.Sprintf("UPDATE %s SET state = %d, version = '%s' WHERE tid = %d",
		tableNameTargetDSN, state, version, tid)
	_, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("更新连接状态出错: %s\n", err)
		log.Warnf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func ClearAllDSNStateAndVersion() error {
	sql := fmt.Sprintf("UPDATE %s SET state = 0, version = '-'", tableNameTargetDSN)
	_, err := db.Exec(sql)
	if err != nil {
		errMsg := fmt.Sprintf("更新连接状态出错: %s\n", err)
		log.Warnf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func GetDSNFromTypeAndVersion(dbType, version string) (*TargetDSN, error) {
	var dsn = TargetDSN{}
	sql := fmt.Sprintf("SELECT * FROM %s WHERE db_type = '%s' AND version = '%s'",
		tableNameTargetDSN, dbType, version)
	err := db.Get(&dsn, sql)
	if err != nil {
		return nil, err
	}
	return &dsn, nil
}
