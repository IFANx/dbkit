package internal

import (
	"dbkit/config"
	"dbkit/internal/model"
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type GlobalState struct {
	Config      *config.DBKitConfig
	DataSource  *sqlx.DB
	TestTasks   map[int]*TaskContext
	VerifyTasks map[int]*TaskContext
}

var globalState *GlobalState

var once sync.Once

func GetState() *GlobalState {
	once.Do(func() {
		globalState = makeGlobalState()
		log.Info("Initialize global state")
	})
	return globalState
}

func makeGlobalState() *GlobalState {
	// 从配置文件读取配置
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fail to read config")
	}
	dbKitConfig := config.DBKitConfig{}
	err = viper.Unmarshal(&dbKitConfig)
	if err != nil {
		log.Fatalf("Fail to parse config")
	}

	state := GlobalState{
		Config:      &dbKitConfig,
		TestTasks:   make(map[int]*TaskContext),
		VerifyTasks: make(map[int]*TaskContext),
		DataSource:  nil,
	}

	// 根据配置文件建立连接
	dataSource := state.Config.DataSource
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/dbkit?parseTime=true",
		dataSource.Username, dataSource.Password, dataSource.Host, dataSource.Port)
	state.DataSource, err = sqlx.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("Fail to connect database")
	} else {
		log.Println("连接数据库成功")
	}
	model.Setup(state.DataSource)

	return &state
}

func (state *GlobalState) GetDataSourceConn() *sqlx.DB {
	return state.DataSource
}

func (state *GlobalState) submitTask(task *TaskContext) {
	if task.Submit.Type == TaskTypeVerify {
		state.TestTasks[task.JobID] = task
	} else {
		state.VerifyTasks[task.JobID] = task
	}
	go task.Start()
}
