package service

import (
	"dbkit/admin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTestJobCount(ctx *gin.Context) {
	count, err := model.GetTestJobCount()
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

func GetTestJobPage(ctx *gin.Context) {
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
	jobs, err := model.GetTestJobPage(pageSize*(pageNum-1), pageSize)
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

func GetTestJobDetail(ctx *gin.Context) {
	jidStr := ctx.DefaultQuery("jid", "0")
	jid, err := strconv.Atoi(jidStr)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	job, err := model.GetTestJobByJid(jid)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	reports, _ := model.GetTestReportByJid(jid)
	statistic, _ := model.GetStatisticByJid(jid)
	ctx.JSON(http.StatusOK, gin.H{
		"ok": true,
		"data": map[string]interface{}{
			"job":       job,
			"reports":   reports,
			"statistic": statistic,
		},
	})
}

func SubTestJob(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ok":  false,
		"err": "功能待实现",
	})
}

func DeleteTestJob(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ok":  false,
		"err": "暂不允许删除操作",
	})
}

func AbortTestJob(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ok":  false,
		"err": "暂不允许手动终止任务",
	})
}
