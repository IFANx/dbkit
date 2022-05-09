package gen

import (
	"dbkit/internal"
	"dbkit/internal/randomly"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type ExprGen struct {
	table      *internal.Table
	depthLimit int
}

func NewExprGen(table *internal.Table, depthLimit int) *ExprGen {
	return &ExprGen{table, depthLimit}
}

func GenPredicate(table *internal.Table) string {
	exprGen := NewExprGen(table, 6)
	expr := exprGen.GenExpr(0)
	if expr == "" {
		return "True"
	} else {
		return expr
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

func (exprGen *ExprGen) GenUaryPrefixOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"NOT", "!", "+", "-"})
	return op + "(" + exprGen.GenExpr(depth+1) + ")"
}

func (exprGen *ExprGen) GenUaryPostfixOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"IS NULL", "IS FALSE", "IS TRUE"})
	return "(" + exprGen.GenExpr(depth+1) + ")" + op
}

func (exprGen *ExprGen) GenBinaryLogicalOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"AND", "&&", "OR", "||", "XOR"})
	return "(" + exprGen.GenExpr(depth+1) + ")" + op +
		"(" + exprGen.GenExpr(depth+1) + ")"
}

func (exprGen *ExprGen) GenBinaryBitOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"&", "|", "^", ">>", "<<"})
	return "(" + exprGen.GenExpr(depth+1) + ")" + op +
		"(" + exprGen.GenExpr(depth+1) + ")"
}

func (exprGen *ExprGen) GenBinaryMathOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"+", "-", "*", "/", "%"})
	return "(" + exprGen.GenExpr(depth+1) + ")" + op +
		"(" + exprGen.GenExpr(depth+1) + ")"
}

func (exprGen *ExprGen) GenBinaryCompOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"=", "!=", "<", "<=", ">", ">=", "LIKE"})
	return "(" + exprGen.GenExpr(depth+1) + ")" + op +
		"(" + exprGen.GenExpr(depth+1) + ")"
}

func (exprGen *ExprGen) GenInOp(depth int) string {
	exprList := []string{"0"}
	for i := 0; i < randomly.RandIntGap(1, 3); i++ {
		exprList = append(exprList, exprGen.GenExpr(depth+1))
	}
	return "(" + exprGen.GenExpr(depth+1) + ") IN ((" +
		strings.Join(exprList, "), (") + "))"
}

func (exprGen *ExprGen) GenBetweenOp(depth int) string {
	fromExpr := exprGen.GenExpr(depth + 1)
	toExpr := exprGen.GenExpr(depth + 1)
	return "(" + exprGen.GenExpr(depth+1) + ") BETWEEN (" +
		fromExpr + ") AND (" + toExpr + ")"
}

func (exprGen *ExprGen) GenCastOp(depth int) string {
	castedExpr := exprGen.GenExpr(depth + 1)
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
