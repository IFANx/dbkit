package model

import (
	"dbkit/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var db *sqlx.DB

const (
	tableNameTestJob       = "test_job"
	tableNameTestReport    = "test_report"
	tableNameTestStatistic = "test_statistic"
	tableNameVerifyJob     = "verify_job"
	tableNameVerifyReport  = "verify_report"
)

func init() {
	// 指定配置文件路径
	viper.SetConfigFile("./config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取全局配置文件出错")
	}
	dbKitConfig := config.DBKitConfig{}
	err = viper.Unmarshal(&dbKitConfig)
	if err != nil {
		log.Fatalf("解析配置文件错误")
	}
	dataSource := dbKitConfig.DataSource
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True",
		dataSource.Username, dataSource.Password, dataSource.Host, dataSource.Port, "dbkit")
	log.Infof("数据库连接参数：%s", connStr)
	db, err = sqlx.Open("mysql", connStr)
	if err != nil {
		log.Panic(err.Error())
		return
	}
}

func CloseDB() {
	db.Close()
}

func CleanUpAbortedJobs() {
	for jid := range RunningTestJobs {
		_ = AbortTestJob(jid)
	}
	for jid := range RunningVerifyJobs {
		_ = AbortVerifyJob(jid)
	}
}
