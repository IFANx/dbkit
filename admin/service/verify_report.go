package service

import (
	"dbkit/admin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetVerifyReportCount(ctx *gin.Context) {
	count, err := model.GetVerifyReportCount()
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

func GetVerifyReportPage(ctx *gin.Context) {
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
	reports, err := model.GetVerifyReportPage(pageSize*(pageNum-1), pageSize)
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

func GetVerifyReportDetail(ctx *gin.Context) {
	ridStr := ctx.DefaultQuery("rid", "0")
	rid, err := strconv.Atoi(ridStr)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	report, err := model.GetVerifyReportByRid(rid)
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

func DeleteVerifyReport(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ok":  false,
		"err": "暂不允许删除操作",
	})
}
