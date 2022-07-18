package gen

import (
	"dbkit/internal/common"
	"dbkit/internal/randomly"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type ExprGen struct {
	table      *common.Table
	depthLimit int
}

func NewExprGen(table *common.Table, depthLimit int) *ExprGen {
	return &ExprGen{table, depthLimit}
}

func GenPredicate(table *common.Table) string {
	exprGen := NewExprGen(table, 3)
	expr := exprGen.GenExpr(0)
	if expr == "" {
		return "True"
	} else {
		return expr
	}
}
func GenPQS(table *common.Table, pivotrow map[string]interface{}) string {
	exprGen := NewExprGen(table, 3)
	expr := exprGen.GenPQSExpr(pivotrow, 0, "")
	if expr == "" {
		return "True"
	} else {
		return expr
	}
}

func (exprGen *ExprGen) GenPQSExpr(pivotrow map[string]interface{}, depth int, str string) string {
	if randomly.ProbFiveOne() || depth > exprGen.depthLimit {
		return exprGen.GenPQSLeaf(pivotrow)
	}
	opNames := []string{"GenPQSColumn", "GenPQSConstant",
		"GenPQSUaryPrefixOp", "GenPQSUaryPostfixOp",
		"GenPQSBinaryLogicalOp", "GenPQSBinaryBitOp", "GenPQSBinaryMathOp", "GenPQSBinaryCompOp",
		"GenPQSInOp", "GenPQSBetweenOp", "GenPQSCastOp", "GenPQSFunction"}

	opName := randomly.RandPickOneStr(opNames)
	switch opName {
	case "GenPQSColumn":
		return exprGen.GenPQSColumn(pivotrow, depth+1, str)
	case "GenPQSConstant":
		return exprGen.GenPQSConstant(pivotrow, depth+1, str)
	case "GenPQSUaryPrefixOp":
		return exprGen.GenPQSUaryPrefixOp(pivotrow, depth+1, str)
	case "GenPQSUaryPostfixOp":
		return exprGen.GenPQSUaryPostfixOp(pivotrow, depth+1, str)
	case "GenPQSBinaryLogicalOp":
		return exprGen.GenPQSBinaryLogicalOp(pivotrow, depth+1, str)
	case "GenPQSBinaryBitOp":
		return exprGen.GenPQSBinaryBitOp(pivotrow, depth+1, str)
	case "GenPQSBinaryMathOp":
		return exprGen.GenPQSBinaryMathOp(pivotrow, depth+1, str)
	case "GenPQSBinaryCompOp":
		return exprGen.GenPQSBinaryCompOp(pivotrow, depth+1, str)
	case "GenPQSInOp":
		return exprGen.GenPQSInOp(pivotrow, depth+1, str)
	case "GenPQSBetweenOp":
		return exprGen.GenPQSBetweenOp(pivotrow, depth+1, str)
	case "GenPQSCastOp":
		return exprGen.GenPQSCastOp(pivotrow, depth+1, str)
	case "GenPQSFunction":
		return exprGen.GenPQSFunction(pivotrow, depth+1, str)
	default:
		return ""

	}
	//return reflect.ValueOf(exprGen).MethodByName(opName).Call(paramList)[0].String()

}
func (exprGen *ExprGen) GenPQSLeaf(pivotrow map[string]interface{}) string {
	if randomly.RandBool() {
		return exprGen.GenPQSColumn(pivotrow, 0, "")
	} else {
		return exprGen.GenConstant(0)
	}
}

func (exprGen *ExprGen) GenExpr(depth int) string {
	if randomly.ProbFiveOne() || depth > exprGen.depthLimit {
		return exprGen.GenLeaf()
	}
	opNames := []string{"GenColumn", "GenConstant",
		"GenUaryPrefixOp", "GenUaryPostfixOp",
		"GenBinaryLogicalOp", "GenBinaryBitOp", "GenBinaryMathOp", "GenBinaryCompOp",
		"GenInOp", "GenBetweenOp", "GenCastOp", "GenFunction"}
	opName := randomly.RandPickOneStr(opNames)
	paramList := []reflect.Value{
		reflect.ValueOf(depth),
	}
	// log.Infof("opName: %s", opName)
	return reflect.ValueOf(exprGen).MethodByName(opName).Call(paramList)[0].String()
}

func (exprGen *ExprGen) GenLeaf() string {
	if randomly.RandBool() {
		return exprGen.GenColumn(0)
	} else {
		return exprGen.GenConstant(0)
	}
}

func (exprGen *ExprGen) GenPQSColumn(pivotrow map[string]interface{}, depth int, str string) string {
	colName := randomly.RandPickOneStr(exprGen.table.ColumnNames)
	colValue, ok := pivotrow[colName].([]byte)
	//log.Infof("colvalue:%s", colValue)
	if ok {
		return string(colValue)
	}
	return exprGen.GenConstant(0)
}

func (exprGen *ExprGen) GenColumn(depth int) string {
	colName := randomly.RandPickOneStr(exprGen.table.ColumnNames)
	// log.Infof("Pick column name：" + colName)
	return colName
}

func (exprGen *ExprGen) GenConstant(depth int) (res string) {
	if randomly.ProbTwentyOne() {
		res = "NULL"
		return
	}
	constType := randomly.RandPickOneStr([]string{"INT", "STRING", "DOUBLE"})
	switch constType {
	case "INT":
		res = fmt.Sprintf("%d", randomly.RandInt32())
	case "STRING":
		res = "'" + randomly.RandNormStrLen(randomly.RandIntGap(5, 10)) + "'"
	case "DOUBLE":
		if randomly.RandBool() {
			res = strconv.FormatFloat(float64(randomly.RandFloat()), 'f',
				randomly.RandIntGap(0, 5), 32)
		} else {
			res = strconv.FormatFloat(randomly.RandDouble(), 'f',
				randomly.RandIntGap(0, 10), 64)
		}
	default:
		res = "0"
	}
	// log.Infof("Generate %s constant：%s", constType, res)
	return
}

func (exprGen *ExprGen) GenPQSConstant(pivotrow map[string]interface{}, depth int, str string) (res string) {
	if randomly.ProbTwentyOne() {
		res = "NULL"
		return
	}
	constType := randomly.RandPickOneStr([]string{"INT", "STRING", "DOUBLE"})
	switch constType {
	case "INT":
		res = fmt.Sprintf("%d", randomly.RandInt32())
	case "STRING":
		res = "'" + randomly.RandNormStrLen(randomly.RandIntGap(5, 10)) + "'"
	case "DOUBLE":
		if randomly.RandBool() {
			res = strconv.FormatFloat(float64(randomly.RandFloat()), 'f',
				randomly.RandIntGap(0, 5), 32)
		} else {
			res = strconv.FormatFloat(randomly.RandDouble(), 'f',
				randomly.RandIntGap(0, 10), 64)
		}
	default:
		res = "0"
	}
	// log.Infof("Generate %s constant：%s", constType, res)
	return
}

func (exprGen *ExprGen) GenUaryPrefixOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"NOT", "!", "+", "-"})
	return op + "(" + exprGen.GenExpr(depth+1) + ")"
}
func (exprGen *ExprGen) GenPQSUaryPrefixOp(pivotrow map[string]interface{}, depth int, str string) string {
	op := randomly.RandPickOneStr([]string{"NOT", "!", "+", "-"})
	return op + "(" + exprGen.GenPQSExpr(pivotrow, depth+1, str) + ")"
}

func (exprGen *ExprGen) GenUaryPostfixOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"IS NULL", "IS FALSE", "IS TRUE"})
	return "(" + exprGen.GenExpr(depth+1) + ")" + op
}
func (exprGen *ExprGen) GenPQSUaryPostfixOp(pivotrow map[string]interface{}, depth int, str string) string {
	op := randomly.RandPickOneStr([]string{"IS NULL", "IS FALSE", "IS TRUE"})
	return "(" + exprGen.GenPQSExpr(pivotrow, depth+1, str) + ")" + op
}

func (exprGen *ExprGen) GenBinaryLogicalOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"AND", "&&", "OR", "||", "XOR"})
	return "(" + exprGen.GenExpr(depth+1) + ")" + op +
		"(" + exprGen.GenExpr(depth+1) + ")"
}
func (exprGen *ExprGen) GenPQSBinaryLogicalOp(pivotrow map[string]interface{}, depth int, str string) string {
	op := randomly.RandPickOneStr([]string{"AND", "&&", "OR", "||", "XOR"})
	return "(" + exprGen.GenPQSExpr(pivotrow, depth+1, str) + ")" + op +
		"(" + exprGen.GenPQSExpr(pivotrow, depth+1, str) + ")"
}

func (exprGen *ExprGen) GenBinaryBitOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"&", "|", "^", ">>", "<<"})
	return "(" + exprGen.GenExpr(depth+1) + ")" + op +
		"(" + exprGen.GenExpr(depth+1) + ")"
}
func (exprGen *ExprGen) GenPQSBinaryBitOp(pivotrow map[string]interface{}, depth int, str string) string {
	op := randomly.RandPickOneStr([]string{"&", "|", "^", ">>", "<<"})
	return "(" + exprGen.GenPQSExpr(pivotrow, depth+1, str) + ")" + op +
		"(" + exprGen.GenPQSExpr(pivotrow, depth+1, str) + ")"
}

func (exprGen *ExprGen) GenBinaryMathOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"+", "-", "*", "/", "%"})
	return "(" + exprGen.GenExpr(depth+1) + ")" + op +
		"(" + exprGen.GenExpr(depth+1) + ")"
}
func (exprGen *ExprGen) GenPQSBinaryMathOp(pivotrow map[string]interface{}, depth int, str string) string {
	op := randomly.RandPickOneStr([]string{"+", "-", "*", "/", "%"})
	return "(" + exprGen.GenPQSExpr(pivotrow, depth+1, str) + ")" + op +
		"(" + exprGen.GenPQSExpr(pivotrow, depth+1, str) + ")"
}

func (exprGen *ExprGen) GenBinaryCompOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"=", "!=", "<", "<=", ">", ">=", "LIKE"})
	return "(" + exprGen.GenExpr(depth+1) + ")" + op +
		"(" + exprGen.GenExpr(depth+1) + ")"
}
func (exprGen *ExprGen) GenPQSBinaryCompOp(pivotrow map[string]interface{}, depth int, str string) string {
	op := randomly.RandPickOneStr([]string{"=", "!=", "<", "<=", ">", ">=", "LIKE"})
	return "(" + exprGen.GenPQSExpr(pivotrow, depth+1, str) + ")" + op +
		"(" + exprGen.GenPQSExpr(pivotrow, depth+1, str) + ")"
}

func (exprGen *ExprGen) GenInOp(depth int) string {
	exprList := []string{"0"}
	for i := 0; i < randomly.RandIntGap(1, 3); i++ {
		exprList = append(exprList, exprGen.GenExpr(depth+1))
	}
	return "(" + exprGen.GenExpr(depth+1) + ") IN ((" +
		strings.Join(exprList, "), (") + "))"
}
func (exprGen *ExprGen) GenPQSInOp(pivotrow map[string]interface{}, depth int, str string) string {
	exprList := []string{"0"}
	for i := 0; i < randomly.RandIntGap(1, 3); i++ {
		exprList = append(exprList, exprGen.GenPQSExpr(pivotrow, depth+1, str))
	}
	return "(" + exprGen.GenPQSExpr(pivotrow, depth+1, str) + ") IN ((" +
		strings.Join(exprList, "), (") + "))"
}

func (exprGen *ExprGen) GenBetweenOp(depth int) string {
	fromExpr := exprGen.GenExpr(depth + 1)
	toExpr := exprGen.GenExpr(depth + 1)
	return "(" + exprGen.GenExpr(depth+1) + ") BETWEEN (" +
		fromExpr + ") AND (" + toExpr + ")"
}
func (exprGen *ExprGen) GenPQSBetweenOp(pivotrow map[string]interface{}, depth int, str string) string {
	fromExpr := exprGen.GenPQSExpr(pivotrow, depth+1, str)
	toExpr := exprGen.GenPQSExpr(pivotrow, depth+1, str)
	return "(" + exprGen.GenPQSExpr(pivotrow, depth+1, str) + ") BETWEEN (" +
		fromExpr + ") AND (" + toExpr + ")"
}

func (exprGen *ExprGen) GenCastOp(depth int) string {
	castedExpr := exprGen.GenExpr(depth + 1)
	castType := randomly.RandPickOneStr([]string{"DATE", "DECIMAL", "SIGNED", "UNSIGNED", "CHAR", "BINARY"})
	return "CAST((" + castedExpr + ") AS " + castType + ")"
}
func (exprGen *ExprGen) GenPQSCastOp(pivotrow map[string]interface{}, depth int, str string) string {
	castedExpr := exprGen.GenPQSExpr(pivotrow, depth+1, str)
	castType := randomly.RandPickOneStr([]string{"DATE", "DECIMAL", "SIGNED", "UNSIGNED", "CHAR", "BINARY"})
	return "CAST((" + castedExpr + ") AS " + castType + ")"
}

func (exprGen *ExprGen) GenFunction(depth int) string {
	function := RandMySQLFunc()
	var argList []string
	for i := 0; i < function.argCnt; i++ {
		argList = append(argList, exprGen.GenExpr(depth+1))
	}
	return function.name + "((" + strings.Join(argList, "), (") + "))"
}
func (exprGen *ExprGen) GenPQSFunction(pivotrow map[string]interface{}, depth int, str string) string {
	function := RandMySQLFunc()
	var argList []string
	for i := 0; i < function.argCnt; i++ {
		argList = append(argList, exprGen.GenPQSExpr(pivotrow, depth+1, str))
	}
	return function.name + "((" + strings.Join(argList, "), (") + "))"
}
