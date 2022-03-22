package main

import (
	"dbkit/internal/randomly"
	"dbkit/internal/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"math/rand"
	"os"
	"time"
)

var logFile *os.File

func init() {
	rand.Seed(time.Now().UnixNano())

	// 指定配置文件路径
	viper.SetConfigFile("./config/config.json")

	testID := time.Now().Format("20060102150405") + randomly.RandAlphabetStrLen(5)
	viper.Set("TestID", testID)

	logFileName := "./config/" + testID + ".log"
	if util.CheckFileIsExist(logFileName) { //如果文件存在
		log.Fatalf("Test ID is not unique, log file exits")
	}
	logFile, err := os.Create(logFileName)
	if err != nil {
		log.Fatalf("Faied to create log file")
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}
