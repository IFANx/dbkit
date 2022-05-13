package service

import (
	"dbkit/admin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetVerifyJobCount(ctx *gin.Context) {
	count, err := model.GetVerifyJobCount()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": count,
		})
	}
}

func GetVerifyJobPage(ctx *gin.Context) {
	pageSize := 10
	page := ctx.DefaultQuery("page", "1")
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	jobs, err := model.GetVerifyJobPage(pageSize*(pageNum-1), pageSize)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": jobs,
		})
	}
}

func GetVerifyJobDetail(ctx *gin.Context) {
	jidStr := ctx.DefaultQuery("jid", "0")
	jid, err := strconv.Atoi(jidStr)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	job, err := model.GetVerifyJobByJid(jid)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	report, _ := model.GetVerifyReportByJid(jid)
	ctx.JSON(http.StatusOK, gin.H{
		"ok": true,
		"data": map[string]interface{}{
			"job":    job,
			"report": report,
		},
	})
}

func SubVerifyJob(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ok":  false,
		"err": "功能待实现",
	})
}
