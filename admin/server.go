package admin

import (
	"context"
	"dbkit/admin/service"
	"dbkit/internal/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func StartServer(port int) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Static("/uploads", "./uploads")
	router.Static("/results", "./results")

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": "hello, DBKit",
		})
	})

	router.POST("/Job/sub", service.SubJob)

	router.GET("/TestJob/count", service.GetTestJobCount)
	router.GET("/TestJob/page", service.GetTestJobPage)
	router.GET("/TestJob/detail", service.GetTestJobDetail)
	router.POST("/TestJob/sub", service.SubTestJob)
	router.GET("/TestJob/delete", service.DeleteTestJob)
	router.GET("/TestJob/abort", service.AbortTestJob)

	router.GET("/TestReport/count", service.GetTestReportCount)
	router.GET("/TestReport/page", service.GetTestReportPage)
	router.GET("/TestReport/detail", service.GetTestReportDetail)
	router.GET("/TestReport/delete", service.DeleteTestReport)

	router.GET("/VerifyJob/count", service.GetVerifyJobCount)
	router.GET("/VerifyJob/page", service.GetVerifyJobPage)
	router.GET("/VerifyJob/detail", service.GetVerifyJobDetail)
	router.POST("/VerifyJob/sub", service.SubVerifyJob)
	router.GET("/VerifyJob/delete", service.DeleteVerifyJob)
	router.GET("/VerifyJob/abort", service.AbortVerifyJob)

	router.GET("/VerifyReport/count", service.GetVerifyReportCount)
	router.GET("/VerifyReport/page", service.GetVerifyReportPage)
	router.GET("/VerifyReport/detail", service.GetVerifyReportDetail)
	router.GET("/VerifyReport/delete", service.DeleteVerifyReport)

	router.GET("/TargetDSN/all", service.GetAllTargetDSN)
	router.GET("/TargetDSN/type", service.GetTargetDSNByType)
	router.POST("/TargetDSN/add", service.AddTargetDSN)
	router.GET("/TargetDSN/check", service.CheckTargetDSN)
	router.GET("/TargetDSN/delete", service.DeleteTargetDSN)
	router.GET("/TargetDSN/typeAvailable", service.GetAvailableVersionByType)
	router.GET("/TargetDSN/allAvailable", service.GetAllAvailableTypeVersionMap)

	router.GET("/SysInfo/all", service.GetAllSysInfo)

	srv := &http.Server{
		Addr:    "0.0.0.0:" + strconv.Itoa(port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("服务器出错: %s\n", err)
			os.Exit(1)
		}
	}()

	log.Infof("服务器启动，在[%d]端口监听请求...\n", port)
	log.Infof("输入 Ctrl + C 关闭服务器\n")

	// 手动终止服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("等待服务器关闭...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("服务器关闭:", err)
	}

	log.Infof("服务器退出，将未执行完毕Job改为执行错误...")
	model.CleanUpAbortedJobs()

	log.Info("关闭数据库链接...")
	model.CloseDB()

	log.Println("程序退出")
}
