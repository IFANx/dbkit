package main

import (
	"dbkit/internal"
	"dbkit/internal/util"
	"io"
	"math/rand"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var logFile *os.File
var state internal.GlobalState

func init() {
	rand.Seed(time.Now().UnixNano())

	// 指定配置文件路径
	viper.SetConfigFile("config.json")

	var (
		logFile *os.File
		err     error
	)

	logFileName := "./log/" + time.Now().Format("2006010215") + ".log"
	if util.CheckFileIsExist(logFileName) { //如果文件存在
		logFile, err = os.OpenFile(logFileName, os.O_APPEND, 0666) //打开文件
		if err != nil {
			log.Fatalf("Failed to open log file.")
		}
	} else {
		logFile, err = os.Create(logFileName)
		if err != nil {
			log.Fatalf("Failed to create log file.")
		}
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	state := internal.GetState()
	conn := state.GetDataSourceConn()
	err = conn.Ping()
	if err != nil {
		log.Fatalf("Failed to build connection.")
	}
}
