package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var db *sqlx.DB

const (
	tableNameTestJob       = "test_job"
	tableNameTestReport    = "test_report"
	tableNameTestStatistic = "test_statistic"
	tableNameVerifyJob     = "verify_job"
	tableNameVerifyReport  = "verify_report"
	tableNameTargetDSN     = "target_dsn"
)

func Setup(conn *sqlx.DB) {
	db = conn
	err := ClearAllDSNStateAndVersion()
	if err != nil {
		log.Info("初始化DSN连接状态和版本出错")
	} else {
		log.Info("初始化DSN连接状态和版本成功")
	}
}

func CloseDB() {
	_ = db.Close()
}
