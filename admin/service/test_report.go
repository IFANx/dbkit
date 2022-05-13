package service

import (
	"dbkit/admin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTestReportCount(ctx *gin.Context) {
	count, err := model.GetTestReportCount()
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

func GetTestReportPage(ctx *gin.Context) {
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
	reports, err := model.GetTestReportPage(pageSize*(pageNum-1), pageSize)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": reports,
		})
	}
}

func GetTestReportDetail(ctx *gin.Context) {
	ridStr := ctx.DefaultQuery("rid", "0")
	rid, err := strconv.Atoi(ridStr)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	report, err := model.GetTestReportByRid(rid)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": report,
		})
	}
}
