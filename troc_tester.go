package main

import (
	"dbkit/internal"
	"dbkit/internal/common"
	"dbkit/internal/common/dbms"
	"dbkit/internal/common/oracle"
	"dbkit/internal/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"time"
)

func TestMySQLTroc() {
	state := internal.GetState()
	log.Println(state.DataSource.DriverName())
	dsn := "root:tobeno.1@tcp(127.0.0.1:3306)/test?"
	conn, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Println("创建数据库连接出错：", err)
	}
	submit := internal.TaskSubmit{
		Type:        internal.TaskTypeTest,
		Oracle:      oracle.Troc,
		TargetTypes: []dbms.DBMS{dbms.MYSQL},
		ConnList:    []*sqlx.DB{conn},
		DSNList:     []string{dsn},
		Limit:       1.5,
		Model:       "",
		Comments:    "",
	}
	ctx := internal.TaskContext{
		JobID:        0,
		Submit:       &submit,
		Runner:       &mysql.MySQLTrocTester{},
		StartTime:    time.Now(),
		Deadline:     time.Now().Add(time.Hour),
		EndTime:      time.Time{},
		SqlCount:     0,
		TestRunCount: 0,
		ReportCount:  0,
		DBList:       make([]*common.Database, 1),
		Aborted:      0,
		Finished:     0,
	}
	ctx.StartTime = time.Now()
	db := &common.Database{
		DBMS:       ctx.Submit.TargetTypes[0],
		DBProvider: &mysql.MySQLProvider{},
		DBName:     "test0",
		Conn:       ctx.Submit.ConnList[0],
		Tables:     make([]*common.Table, 0),
	}
	err = db.Refresh()
	if err != nil {
		panic("初始化测试数据库实例失败：" + err.Error())
	}
	ctx.DBList[0] = db
	ctx.Runner.RunTask(&ctx)
}
