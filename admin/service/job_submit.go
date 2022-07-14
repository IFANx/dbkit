package service

import (
	"dbkit/internal"
	"dbkit/internal/common/dbms"
	"dbkit/internal/common/oracle"
	"dbkit/internal/model"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SubJob(ctx *gin.Context) {
	tp := ctx.DefaultPostForm("type", "")
	testType := ctx.DefaultPostForm("testType", "")
	oracleName := ctx.DefaultPostForm("oracle", "")
	target := getArrayFromPostFormMap(ctx.PostFormMap("target"))
	targetKind := ctx.DefaultPostForm("targetKind", "")
	deadline := ctx.DefaultPostForm("deadline", "false")
	time := ctx.DefaultPostForm("time", "-1")
	checkModel := ctx.DefaultPostForm("model", "")
	op := ctx.DefaultPostForm("op", "")
	desc := ctx.DefaultPostForm("desc", "")
	var taskSubmit *internal.TaskSubmit
	var err error
	if tp == "test" {
		taskSubmit, err = validateTestForm(oracleName, target, testType, deadline, time, desc)
	} else if tp == "diff" {
		taskSubmit, err = validateDiffForm(oracleName, target, targetKind, deadline, time, desc)
	} else if tp == "verify" {
		taskSubmit, err = validateVerifyForm(oracleName, target, checkModel, op, desc)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": "测试类型未知",
		})
	}
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	if jid, err := internal.BuildTaskFromSubmit(taskSubmit); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": jid,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
	}
}

func validateTestForm(oracleName string, target []string, testType, deadline, time, desc string) (*internal.TaskSubmit, error) {
	oracleObj := oracle.GetOracleFromStr(oracleName)
	if testType == "query" {
		if oracleObj != oracle.TLP && oracleObj != oracle.NoREC && oracleObj != oracle.DQE && oracleObj != oracle.NoREC2 && oracleObj != oracle.PQS {
			return nil, errors.New("不支持的oracle类型: " + oracleName)
		}
	} else if testType == "txn" {
		if oracleObj != oracle.Troc && oracleObj != oracle.TrocPlus {
			return nil, errors.New("不支持的oracle类型: " + oracleName)
		}
	} else {
		return nil, errors.New("不支持的测试类型: " + testType)
	}
	if len(target) < 2 {
		return nil, errors.New(fmt.Sprintf("测试对象参数非法: %v", target))
	}
	var (
		dbmsList []dbms.DBMS
		connList []*sqlx.DB
		dsnList  []string
	)
	targetDSN, err := model.GetDSNFromTypeAndVersion(target[0], target[1])
	if err != nil {
		return nil, errors.New("查询测试对象连接参数错误: " + err.Error())
	}
	conn, err := targetDSN.GetConn()
	if err != nil {
		return nil, errors.New("测试对象连接错误: " + err.Error())
	}
	dbmsList = append(dbmsList, dbms.GetDBMSFromStr(targetDSN.DBType))
	connList = append(connList, conn)
	dsnList = append(dsnList, targetDSN.GetDSN())
	var limit float32 = 0
	if deadline == "true" {
		timeLimit, err := strconv.ParseFloat(time, 32)
		if err != nil {
			return nil, errors.New("限时参数解析错误: " + time)
		}
		limit = float32(timeLimit)
	}
	return &internal.TaskSubmit{
		Type:        internal.TaskTypeTest,
		Oracle:      oracleObj,
		TargetTypes: dbmsList,
		ConnList:    connList,
		DSNList:     dsnList,
		Limit:       limit,
		Comments:    desc,
	}, nil
}

func validateDiffForm(oracleName string, target []string, targetKind, deadline, time, desc string) (*internal.TaskSubmit, error) {
	var oracleObj oracle.Oracle
	if oracleName == "query" {
		oracleObj = oracle.DIFF
	} else if oracleName == "txn" {
		oracleObj = oracle.DIFFTXN
	} else {
		return nil, errors.New("不支持的oracle类型: " + oracleName)
	}
	if len(target) < 4 {
		return nil, errors.New(fmt.Sprintf("测试对象参数非法: %v", target))
	}
	var (
		dbmsList []dbms.DBMS
		connList []*sqlx.DB
		dsnList  []string
	)
	for i := 0; i < len(target); i += 2 {
		targetDSN, err := model.GetDSNFromTypeAndVersion(target[i], target[i+1])
		if err != nil {
			return nil, errors.New("查询测试对象连接参数错误: " + err.Error())
		}
		conn, err := targetDSN.GetConn()
		if err != nil {
			return nil, errors.New("测试对象连接错误: " + err.Error())
		}
		dbmsList = append(dbmsList, dbms.GetDBMSFromStr(targetDSN.DBType))
		connList = append(connList, conn)
		dsnList = append(dsnList, targetDSN.GetDSN())
	}
	var limit float32 = 0
	if deadline == "true" {
		timeLimit, err := strconv.ParseFloat(time, 32)
		if err != nil {
			return nil, errors.New("限时参数解析错误: " + time)
		}
		limit = float32(timeLimit)
	}
	return &internal.TaskSubmit{
		Type:        internal.TaskTypeDiff,
		Oracle:      oracleObj,
		TargetTypes: dbmsList,
		ConnList:    connList,
		DSNList:     dsnList,
		Limit:       limit,
		Comments:    desc,
	}, nil
}

func validateVerifyForm(oracleName string, target []string, checkModel, op, desc string) (*internal.TaskSubmit, error) {
	var oracleObj oracle.Oracle
	if oracleName == "linear" {
		oracleObj = oracle.LINEAR
	} else {
		return nil, errors.New("不支持的oracle类型: " + oracleName)
	}
	if len(target) < 2 {
		return nil, errors.New(fmt.Sprintf("测试对象参数非法: %v", target))
	}
	var (
		dbmsList []dbms.DBMS
		connList []*sqlx.DB
		dsnList  []string
	)
	targetDSN, err := model.GetDSNFromTypeAndVersion(target[0], target[1])
	if err != nil {
		return nil, errors.New("查询测试对象连接参数错误: " + err.Error())
	}
	conn, err := targetDSN.GetConn()
	if err != nil {
		return nil, errors.New("测试对象连接错误: " + err.Error())
	}
	dbmsList = append(dbmsList, dbms.GetDBMSFromStr(targetDSN.DBType))
	connList = append(connList, conn)
	dsnList = append(dsnList, targetDSN.GetDSN())
	var limit float32 = 0
	timeLimit, err := strconv.ParseInt(op, 10, 32)
	if err != nil {
		return nil, errors.New("限时参数解析错误: " + op)
	}
	limit = float32(timeLimit)
	return &internal.TaskSubmit{
		Type:        internal.TaskTypeVerify,
		Oracle:      oracleObj,
		TargetTypes: dbmsList,
		ConnList:    connList,
		DSNList:     dsnList,
		Limit:       limit,
		Model:       checkModel,
		Comments:    desc,
	}, nil
}

func getArrayFromPostFormMap(formMap map[string]string) []string {
	res := make([]string, 0)
	for i := 0; i < len(formMap); i++ {
		val := formMap[strconv.Itoa(i)]
		res = append(res, val)
	}
	return res
}
