package admin

import (
	"dbkit/admin/model"
	"dbkit/admin/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

func StartServer(port int) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Static("/uploads", "./uploads")
	r.Static("/results", "./results")

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": "hello, dbkit",
		})
	})
	r.GET("/TestJob/count", service.GetTestJobCount)
	r.GET("/TestJob/page", service.GetTestJobPage)

	r.GET("/TestReport/count", service.GetTestReportCount)
	r.GET("/TestReport/page", service.GetTestReportPage)

	r.GET("/VerifyJob/count", service.GetVerifyJobCount)
	r.GET("/VerifyJob/page", service.GetVerifyJobPage)

	r.GET("/VerifyReport/count", service.GetVerifyReportCount)
	r.GET("/VerifyReport/page", service.GetVerifyReportPage)

	log.Infof("服务器启动，在[%d]端口监听请求...", port)
	err := r.Run("0.0.0.0:" + strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Infof("程序退出，将未执行完毕Job改为执行错误...")
	model.CleanUpAbortedJobs()

	log.Info("关闭数据库链接...")
	model.CloseDB()
}
