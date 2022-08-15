package internal

import (
	"dbkit/config"
	"dbkit/internal/model"
	"errors"
	"fmt"
	_ "gitee.com/chunanyong/dm"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
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
	fmt.Println(dataSource.Username, dataSource.Password, dataSource.Host, dataSource.Port)
	//connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/dbkit?parseTime=true",
	//	dataSource.Username, dataSource.Password, dataSource.Host, dataSource.Port)
	connStr := fmt.Sprintf("dm://%s:%s@%s:%d/dbkit?parseTime=true&compatibleMode=mysql",
		dataSource.Username, dataSource.Password, dataSource.Host, dataSource.Port)
	state.DataSource, err = sqlx.Open("dm", connStr)
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

func (state *GlobalState) SubmitTask(task *TaskContext) {
	if task.Submit.Type == TaskTypeVerify {
		state.VerifyTasks[task.JobID] = task
	} else {
		state.TestTasks[task.JobID] = task
	}
	go task.Start()
}

func (state *GlobalState) AbortTask(taskType TaskType, jid int) error {
	var task *TaskContext
	if taskType == TaskTypeVerify {
		task = state.VerifyTasks[jid]
		err := model.AbortVerifyJob(jid)
		if err != nil {
			return errors.New("手动终止任务失败：" + err.Error())
		}
	} else {
		task = state.TestTasks[jid]
		err := model.AbortTestJob(jid)
		if err != nil {
			return errors.New("手动终止任务失败：" + err.Error())
		}
		_, err = model.AddStatistic(jid, task.SqlCount, task.TestRunCount, task.ReportCount, "手动终止")
		if err != nil {
			return errors.New("手动终止任务失败：" + err.Error())
		}
	}
	if task == nil {
		return errors.New("未找到该运行中的任务，请刷新页面")
	}
	task.Abort()
	return nil
}

func (state *GlobalState) AbortAllRunningTasks() {
	for _, task := range state.TestTasks {
		_ = state.AbortTask(TaskTypeTest, task.JobID)
	}
	for _, task := range state.VerifyTasks {
		_ = state.AbortTask(TaskTypeVerify, task.JobID)
	}
}
