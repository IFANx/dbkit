package gen

import (
	"dbkit/internal"
	"dbkit/internal/randomly"
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
	expr := exprGen.genExpr(0)
	if expr == "" {
		return "True"
	} else {
		return expr
	}
}

func (exprGen *ExprGen) genExpr(depth int) string {
	if randomly.RandBool() || depth > exprGen.depthLimit {
		return exprGen.genLeaf()
	}
	opNames := []string{"genColumn", "genConstant",
		"genUaryPrefixOp", "genUaryPostfixOp",
		"genBinaryLogicalOp", "genBinaryBitOp", "genBinaryMathOp", "genBinaryCompOp",
		"genInOp", "genBetweenOp", "genCastOp", "genFunction"}
	opName := randomly.RandPickOneStr(opNames)
	paramList := []reflect.Value{
		reflect.ValueOf(depth),
	}
	return reflect.ValueOf(exprGen).MethodByName(opName).Call(paramList)[0].String()
}

func (exprGen *ExprGen) genLeaf() string {
	if randomly.RandBool() {
		return exprGen.genColumn()
	} else {
		return exprGen.genConstant()
	}
}

func (exprGen *ExprGen) genColumn() string {
	return randomly.RandPickOneStr(exprGen.table.ColumnNames)
}

func (exprGen *ExprGen) genConstant() string {
	if randomly.ProbTwentyOne() {
		return "NULL"
	}
	constType := randomly.RandPickOneStr([]string{"INT", "STRING", "DOUBLE"})
	switch constType {
	case "INT":
		return string(randomly.RandInt32())
	case "STRING":
		return "\"" + randomly.RandPrintStrLen(randomly.RandIntGap(5, 10)) + "\""
	case "DOUBLE":
		if randomly.RandBool() {
			return strconv.FormatFloat(float64(randomly.RandFloat()), 'f',
				randomly.RandIntGap(0, 5), 32)
		} else {
			return strconv.FormatFloat(randomly.RandDouble(), 'f',
				randomly.RandIntGap(0, 10), 64)
		}
	}
	return "0"
}

func (exprGen *ExprGen) genUaryPrefixOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"NOT", "!", "+", "-"})
	return op + "(" + exprGen.genExpr(depth+1) + ")"
}

func (exprGen *ExprGen) genUaryPostfixOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"IS NULL", "IS FALSE", "IS TRUE"})
	return "(" + exprGen.genExpr(depth+1) + ")" + op
}

func (exprGen *ExprGen) genBinaryLogicalOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"AND", "&&", "OR", "||", "XOR"})
	return "(" + exprGen.genExpr(depth+1) + ")" + op +
		"(" + exprGen.genExpr(depth+1) + ")"
}

func (exprGen *ExprGen) genBinaryBitOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"&", "|", "^", ">>", "<<"})
	return "(" + exprGen.genExpr(depth+1) + ")" + op +
		"(" + exprGen.genExpr(depth+1) + ")"
}

func (exprGen *ExprGen) genBinaryMathOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"+", "-", "*", "/", "%"})
	return "(" + exprGen.genExpr(depth+1) + ")" + op +
		"(" + exprGen.genExpr(depth+1) + ")"
}

func (exprGen *ExprGen) genBinaryCompOp(depth int) string {
	op := randomly.RandPickOneStr([]string{"=", "!=", "<", "<=", ">", ">=", "LIKE"})
	return "(" + exprGen.genExpr(depth+1) + ")" + op +
		"(" + exprGen.genExpr(depth+1) + ")"
}

func (exprGen *ExprGen) genInOp(depth int) string {
	exprList := []string{"0"}
	for i := 0; i < randomly.RandIntGap(1, 3); i++ {
		exprList = append(exprList, exprGen.genExpr(depth+1))
	}
	return "(" + exprGen.genExpr(depth+1) + ") IN ((" +
		strings.Join(exprList, "), (") + "))"
}

func (exprGen *ExprGen) genBetweenOp(depth int) string {
	fromExpr := exprGen.genExpr(depth + 1)
	toExpr := exprGen.genExpr(depth + 1)
	return "(" + exprGen.genExpr(depth+1) + ") BETWEEN (" +
		fromExpr + ") AND (" + toExpr + ")"
}

func (exprGen *ExprGen) genCastOp(depth int) string {
	castedExpr := exprGen.genExpr(depth + 1)
	castType := randomly.RandPickOneStr([]string{"INT", "FLOAT", "DOUBLE", "CHAR"})
	return "CAST((" + castedExpr + ") AS " + castType + ")"
}

func (exprGen *ExprGen) genFunction(depth int) string {
	function := RandMySQLFunc()
	var argList []string
	for i := 0; i < function.argCnt; i++ {
		argList = append(argList, exprGen.genExpr(depth+1))
	}
	return function.name + "((" + strings.Join(argList, "), (") + "))"
}
