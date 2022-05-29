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
			"data": "hello, DBKit",
		})
	})

	r.GET("/TestJob/count", service.GetTestJobCount)
	r.GET("/TestJob/page", service.GetTestJobPage)
	r.GET("/TestJob/detail", service.GetTestJobDetail)
	r.POST("/TestJob/sub", service.SubTestJob)
	r.GET("/TestJob/delete", service.DeleteTestJob)
	r.GET("/TestJob/abort", service.AbortTestJob)

	r.GET("/TestReport/count", service.GetTestReportCount)
	r.GET("/TestReport/page", service.GetTestReportPage)
	r.GET("/TestReport/detail", service.GetTestReportDetail)
	r.GET("/TestReport/delete", service.DeleteTestReport)

	r.GET("/VerifyJob/count", service.GetVerifyJobCount)
	r.GET("/VerifyJob/page", service.GetVerifyJobPage)
	r.GET("/VerifyJob/detail", service.GetVerifyJobDetail)
	r.POST("/VerifyJob/sub", service.SubVerifyJob)
	r.GET("/VerifyJob/delete", service.DeleteVerifyJob)
	r.GET("/VerifyJob/abort", service.AbortVerifyJob)

	r.GET("/VerifyReport/count", service.GetVerifyReportCount)
	r.GET("/VerifyReport/page", service.GetVerifyReportPage)
	r.GET("/VerifyReport/detail", service.GetVerifyReportDetail)
	r.GET("/VerifyReport/delete", service.DeleteVerifyReport)

	r.GET("/TargetDSN/all", service.GetAllTargetDSN)
	r.GET("/TargetDSN/type", service.GetTargetDSNByType)
	r.POST("/TargetDSN/add", service.AddTargetDSN)
	r.GET("/TargetDSN/check", service.CheckTargetDSN)
	r.GET("/TargetDSN/delete", service.DeleteTargetDSN)
	r.GET("/TargetDSN/available", service.GetAvailableVersionByType)

	r.GET("/SysInfo/all", service.GetAllSysInfo)

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
