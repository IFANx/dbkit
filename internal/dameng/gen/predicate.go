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
func GenPQS(table *common.Table, pivotRow map[string]interface{}) (string, string) {
	exprGen := NewExprGen(table, 3)
	exprWithName, exprWithValue := exprGen.GenPQSExpr(pivotRow, 0, "")
	if exprWithName == "" || exprWithValue == "" {
		return "True", "True"
	} else {
		return exprWithName, exprWithValue
	}
}

func (exprGen *ExprGen) GenPQSExpr(pivotRow map[string]interface{}, depth int, str string) (string, string) {
	if randomly.ProbFiveOne() || depth > exprGen.depthLimit {
		return exprGen.GenPQSLeaf(pivotRow)
	}
	opNames := []string{"GenPQSColumn", "GenPQSConstant",
		"GenPQSUaryPrefixOp", "GenPQSUaryPostfixOp",
		"GenPQSBinaryLogicalOp", "GenPQSBinaryBitOp", "GenPQSBinaryMathOp", "GenPQSBinaryCompOp",
		"GenPQSInOp", "GenPQSBetweenOp", "GenPQSCastOp", "GenPQSFunction"}

	opName := randomly.RandPickOneStr(opNames)
	switch opName {
	case "GenPQSColumn":
		return exprGen.GenPQSColumn(pivotRow, depth+1, str)
	case "GenPQSConstant":
		return exprGen.GenPQSConstant(pivotRow, depth+1, str)
	case "GenPQSUaryPrefixOp":
		return exprGen.GenPQSUaryPrefixOp(pivotRow, depth+1, str)
	case "GenPQSUaryPostfixOp":
		return exprGen.GenPQSUaryPostfixOp(pivotRow, depth+1, str)
	case "GenPQSBinaryLogicalOp":
		return exprGen.GenPQSBinaryLogicalOp(pivotRow, depth+1, str)
	case "GenPQSBinaryBitOp":
		return exprGen.GenPQSBinaryBitOp(pivotRow, depth+1, str)
	case "GenPQSBinaryMathOp":
		return exprGen.GenPQSBinaryMathOp(pivotRow, depth+1, str)
	case "GenPQSBinaryCompOp":
		return exprGen.GenPQSBinaryCompOp(pivotRow, depth+1, str)
	case "GenPQSInOp":
		return exprGen.GenPQSInOp(pivotRow, depth+1, str)
	case "GenPQSBetweenOp":
		return exprGen.GenPQSBetweenOp(pivotRow, depth+1, str)
	case "GenPQSCastOp":
		return exprGen.GenPQSCastOp(pivotRow, depth+1, str)
	case "GenPQSFunction":
		return exprGen.GenPQSFunction(pivotRow, depth+1, str)
	default:
		return "", ""

	}
	//return reflect.ValueOf(exprGen).MethodByName(opName).Call(paramList)[0].String()

}
func (exprGen *ExprGen) GenPQSLeaf(pivotRow map[string]interface{}) (string, string) {
	if randomly.RandBool() {
		return exprGen.GenPQSColumn(pivotRow, 0, "")
	} else {
		return exprGen.GenPQSConstant(pivotRow, 0, "")
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

func (exprGen *ExprGen) GenPQSColumn(pivotRow map[string]interface{}, depth int, str string) (string, string) {
	colName := randomly.RandPickOneStr(exprGen.table.ColumnNames)
	colValue, ok := pivotRow[colName].([]byte)
	if ok {
		//根据列类型生成相应的sql，比如varchar需要带“”，int不用
		return colName, string(colValue)
	}
	return exprGen.GenPQSConstant(pivotRow, 0, "")
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
func (exprGen *ExprGen) GenPQSConstant(pivotRow map[string]interface{}, depth int, str string) (res1, res string) {
	if randomly.ProbTwentyOne() {
		res = "NULL"
		res1 = "NULL"
		return
	}
	constType := randomly.RandPickOneStr([]string{"INT", "STRING", "DOUBLE"})
	switch constType {
	case "INT":
		res = fmt.Sprintf("%d", randomly.RandInt32())
		res1 = res
	case "STRING":
		res = "'" + randomly.RandNormStrLen(randomly.RandIntGap(5, 10)) + "'"
		res1 = res
	case "DOUBLE":
		if randomly.RandBool() {
			res = strconv.FormatFloat(float64(randomly.RandFloat()), 'f',
				randomly.RandIntGap(0, 5), 32)
			res1 = res
		} else {
			res = strconv.FormatFloat(randomly.RandDouble(), 'f',
				randomly.RandIntGap(0, 10), 64)
			res1 = res
		}
	default:
		res = "0"
		res1 = res
	}
	// log.Infof("Generate %s constant：%s", constType, res)
	return
}

func (exprGen *ExprGen) GenUaryPrefixOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"NOT", "!", "+", "-"})
	return op + "(" + exprGen.GenExpr(depth+1) + ")"
}
func (exprGen *ExprGen) GenPQSUaryPrefixOp(pivotRow map[string]interface{}, depth int, str string) (string, string) {
	op := randomly.RandPickOneStr([]string{"NOT", "!", "+", "-"})
	str1, str2 := exprGen.GenPQSExpr(pivotRow, depth+1, str)
	return op + "(" + str1 + ")", op + "(" + str2 + ")"
}

func (exprGen *ExprGen) GenUaryPostfixOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"IS NULL", "IS FALSE", "IS TRUE"})
	return "(" + exprGen.GenExpr(depth+1) + ")" + op
}
func (exprGen *ExprGen) GenPQSUaryPostfixOp(pivotRow map[string]interface{}, depth int, str string) (string, string) {
	op := randomly.RandPickOneStr([]string{"IS NULL", "IS FALSE", "IS TRUE"})
	str1, str2 := exprGen.GenPQSExpr(pivotRow, depth+1, str)
	return "(" + str1 + ")" + op, "(" + str2 + ")" + op
}

func (exprGen *ExprGen) GenBinaryLogicalOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"AND", "&&", "OR", "||", "XOR"})
	return "(" + exprGen.GenExpr(depth+1) + ")" + op +
		"(" + exprGen.GenExpr(depth+1) + ")"
}
func (exprGen *ExprGen) GenPQSBinaryLogicalOp(pivotRow map[string]interface{}, depth int, str string) (string, string) {
	op := randomly.RandPickOneStr([]string{"AND", "&&", "OR", "||", "XOR"})
	str1, str2 := exprGen.GenPQSExpr(pivotRow, depth+1, str)
	str3, str4 := exprGen.GenPQSExpr(pivotRow, depth+1, str)
	return "(" + str1 + ")" + op + "(" + str3 + ")", "(" + str2 + ")" + op + "(" + str4 + ")"
}

func (exprGen *ExprGen) GenBinaryBitOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"&", "|", "^", ">>", "<<"})
	return "(" + exprGen.GenExpr(depth+1) + ")" + op +
		"(" + exprGen.GenExpr(depth+1) + ")"
}
func (exprGen *ExprGen) GenPQSBinaryBitOp(pivotRow map[string]interface{}, depth int, str string) (string, string) {
	op := randomly.RandPickOneStr([]string{"&", "|", "^", ">>", "<<"})
	str1, str2 := exprGen.GenPQSExpr(pivotRow, depth+1, str)
	str3, str4 := exprGen.GenPQSExpr(pivotRow, depth+1, str)
	return "(" + str1 + ")" + op + "(" + str3 + ")", "(" + str2 + ")" + op + "(" + str4 + ")"
}

func (exprGen *ExprGen) GenBinaryMathOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"+", "-", "*", "/", "%"})
	return "(" + exprGen.GenExpr(depth+1) + ")" + op +
		"(" + exprGen.GenExpr(depth+1) + ")"
}
func (exprGen *ExprGen) GenPQSBinaryMathOp(pivotRow map[string]interface{}, depth int, str string) (string, string) {
	op := randomly.RandPickOneStr([]string{"+", "-", "*", "/", "%"})
	str1, str2 := exprGen.GenPQSExpr(pivotRow, depth+1, str)
	str3, str4 := exprGen.GenPQSExpr(pivotRow, depth+1, str)
	return "(" + str1 + ")" + op + "(" + str3 + ")", "(" + str2 + ")" + op + "(" + str4 + ")"
}

func (exprGen *ExprGen) GenBinaryCompOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"=", "!=", "<", "<=", ">", ">=", "LIKE"})
	return "(" + exprGen.GenExpr(depth+1) + ")" + op +
		"(" + exprGen.GenExpr(depth+1) + ")"
}
func (exprGen *ExprGen) GenPQSBinaryCompOp(pivotRow map[string]interface{}, depth int, str string) (string, string) {
	op := randomly.RandPickOneStr([]string{"=", "!=", "<", "<=", ">", ">=", "LIKE"})
	str1, str2 := exprGen.GenPQSExpr(pivotRow, depth+1, str)
	str3, str4 := exprGen.GenPQSExpr(pivotRow, depth+1, str)
	return "(" + str1 + ")" + op + "(" + str3 + ")", "(" + str2 + ")" + op + "(" + str4 + ")"
}

func (exprGen *ExprGen) GenInOp(depth int) string {
	exprList := []string{"0"}
	for i := 0; i < randomly.RandIntGap(1, 3); i++ {
		exprList = append(exprList, exprGen.GenExpr(depth+1))
	}
	return "(" + exprGen.GenExpr(depth+1) + ") IN ((" +
		strings.Join(exprList, "), (") + "))"
}
func (exprGen *ExprGen) GenPQSInOp(pivotRow map[string]interface{}, depth int, str string) (string, string) {
	exprList1 := []string{"0"}
	exprList2 := []string{"0"}
	for i := 0; i < randomly.RandIntGap(1, 3); i++ {
		str1, str2 := exprGen.GenPQSExpr(pivotRow, depth+1, str)
		exprList1 = append(exprList1, str1)
		exprList2 = append(exprList2, str2)
	}
	str3, str4 := exprGen.GenPQSExpr(pivotRow, depth+1, str)
	return "(" + str3 + ") IN ((" + strings.Join(exprList1, "), (") + "))", "(" + str4 + ") IN ((" + strings.Join(exprList2, "), (") + "))"
}

func (exprGen *ExprGen) GenBetweenOp(depth int) string {
	fromExpr := exprGen.GenExpr(depth + 1)
	toExpr := exprGen.GenExpr(depth + 1)
	return "(" + exprGen.GenExpr(depth+1) + ") BETWEEN (" +
		fromExpr + ") AND (" + toExpr + ")"
}
func (exprGen *ExprGen) GenPQSBetweenOp(pivotRow map[string]interface{}, depth int, str string) (string, string) {
	str1, str2 := exprGen.GenPQSExpr(pivotRow, depth+1, str)
	str3, str4 := exprGen.GenPQSExpr(pivotRow, depth+1, str)
	str5, str6 := exprGen.GenPQSExpr(pivotRow, depth+1, str)
	//fromExpr := exprGen.GenPQSExpr(pivotRow, depth+1, str)
	//toExpr := exprGen.GenPQSExpr(pivotRow, depth+1, str)
	return "(" + str5 + ") BETWEEN (" + str1 + ") AND (" + str3 + ")", "(" + str6 + ") BETWEEN (" + str2 + ") AND (" + str4 + ")"
}

func (exprGen *ExprGen) GenCastOp(depth int) string {
	castedExpr := exprGen.GenExpr(depth + 1)
	castType := randomly.RandPickOneStr([]string{"DATE", "DECIMAL", "SIGNED", "UNSIGNED", "CHAR", "BINARY"})
	return "CAST((" + castedExpr + ") AS " + castType + ")"
}
func (exprGen *ExprGen) GenPQSCastOp(pivotRow map[string]interface{}, depth int, str string) (string, string) {
	str1, str2 := exprGen.GenPQSExpr(pivotRow, depth+1, str)
	//castedExpr := exprGen.GenPQSExpr(pivotRow, depth+1, str)
	castType := randomly.RandPickOneStr([]string{"DATE", "DECIMAL", "SIGNED", "UNSIGNED", "CHAR", "BINARY"})
	return "CAST((" + str1 + ") AS " + castType + ")", "CAST((" + str2 + ") AS " + castType + ")"
}

func (exprGen *ExprGen) GenFunction(depth int) string {
	function := RandMySQLFunc()
	var argList []string
	for i := 0; i < function.argCnt; i++ {
		argList = append(argList, exprGen.GenExpr(depth+1))
	}
	return function.name + "((" + strings.Join(argList, "), (") + "))"
}
func (exprGen *ExprGen) GenPQSFunction(pivotRow map[string]interface{}, depth int, str string) (string, string) {
	function := RandMySQLFunc()
	var argList1 []string
	var argList2 []string
	for i := 0; i < function.argCnt; i++ {
		str1, str2 := exprGen.GenPQSExpr(pivotRow, depth+1, str)
		argList1 = append(argList1, str1)
		argList2 = append(argList2, str2)
	}
	return function.name + "((" + strings.Join(argList1, "), (") + "))", function.name + "((" + strings.Join(argList2, "), (") + "))"
}
