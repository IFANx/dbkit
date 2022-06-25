package service

import (
	model2 "dbkit/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTestJobCount(ctx *gin.Context) {
	count, err := model2.GetTestJobCount()
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
	jobs, err := model2.GetTestJobPage(pageSize*(pageNum-1), pageSize)
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
	job, err := model2.GetTestJobByJid(jid)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	reports, _ := model2.GetTestReportByJid(jid)
	statistic, _ := model2.GetStatisticByJid(jid)
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
	jidStr := ctx.Query("jid")
	jid, err := strconv.Atoi(jidStr)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	err = model2.DeleteTestJob(jid)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": jid,
	})
}

func AbortTestJob(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ok":  false,
		"err": "暂不允许手动终止任务",
	})
}
