package internal

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
)

type GlobalState struct {
	TestID string
	Config *DBKitConfig
}

var globalState *GlobalState

var once sync.Once

func GetState() *GlobalState {
	once.Do(func() {
		globalState = makeGlobalState()
	})
	return globalState
}

func makeGlobalState() *GlobalState {
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件错误")
	}
	config := DBKitConfig{}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("解析配置文件错误")
	}
	return &GlobalState{
		TestID: viper.GetString("TestID"),
		Config: &config,
	}
}
