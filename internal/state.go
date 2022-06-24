package internal

import (
	"dbkit/config"
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type GlobalState struct {
	Config     *config.DBKitConfig
	Tasks      []*TaskContext
	DataSource *sqlx.DB
	TableCount int
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
		Config:     &dbKitConfig,
		Tasks:      make([]*TaskContext, 0),
		DataSource: nil,
	}

	// 根据配置文件建立连接
	dataSource := state.Config.DataSource
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/dbkit",
		dataSource.Username, dataSource.Password, dataSource.Host, dataSource.Port)
	state.DataSource, err = sqlx.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("Fail to connect database")
	}

	return &state
}

func (state *GlobalState) GetDataSourceConn() *sqlx.DB {
	return state.DataSource
}

func (state *GlobalState) buildTestContext(submit *TaskSubmit) *TaskContext {
	ctx := NewTestContext(submit)
	state.Tasks = append(state.Tasks, ctx)
	return ctx
}
