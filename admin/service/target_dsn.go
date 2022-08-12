package service

import (
	"dbkit/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllTargetDSN(ctx *gin.Context) {
	dsn, err := model.GetAllTargetDSN()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": dsn,
		})
	}
}

func GetTargetDSNByType(ctx *gin.Context) {
	tp := ctx.DefaultQuery("type", "mysql")
	dsn, err := model.GetTargetDSNByType(tp)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": dsn,
		})
	}
}

func GetAllAvailableTypeVersionMap(ctx *gin.Context) {
	verMap, err := model.GetAvailableTypeVersionMap()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": verMap,
		})
	}
}

func GetAvailableVersionByType(ctx *gin.Context) {
	tp := ctx.DefaultQuery("type", "mysql")
	dsnMap, err := model.GetAvailableDSNVersionByType(tp)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": dsnMap,
		})
	}
}

func AddTargetDSN(ctx *gin.Context) {
	tp := ctx.DefaultPostForm("type", "")
	host := ctx.DefaultPostForm("host", "127.0.0.1")
	user := ctx.DefaultPostForm("user", "root")
	pwd := ctx.DefaultPostForm("pwd", "")
	dbName := ctx.DefaultPostForm("dbName", "test")
	params := ctx.DefaultPostForm("params", "")
	portStr := ctx.DefaultPostForm("port", "")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	if tp != "mysql" && tp != "tidb" && tp != "mariadb" && tp != "postgresql" && tp != "cockroachdb" {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": "Database type is not supported:" + tp,
		})
		return
	}
	tid, err := model.AddTargetDSN(tp, host, user, pwd, dbName, params, port)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": tid,
	})
}

func DeleteTargetDSN(ctx *gin.Context) {
	tidStr := ctx.Query("tid")
	tid, err := strconv.Atoi(tidStr)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	err = model.DeleteTargetDSN(tid)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": tid,
	})
}

func CheckTargetDSN(ctx *gin.Context) {
	tidStr := ctx.Query("tid")
	tid, err := strconv.Atoi(tidStr)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	dsn, err := model.GetTargetDSNByTid(tid)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	version, err := model.GetTargetDSNVersion(dsn.DBType, dsn.DBHost, dsn.DBUser,
		dsn.DBPwd, dsn.DBName, dsn.Params, dsn.DBPort)
	//version = "8.0.1"
	if err != nil {
		_ = model.UpdateStateAndVersionByTid(tid, -1, "-")
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	err = model.UpdateStateAndVersionByTid(tid, 1, version)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": version,
	})
}
