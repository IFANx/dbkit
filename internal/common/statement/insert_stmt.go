package statement

import (
	"dbkit/internal/common"
	"dbkit/internal/common/ast"
	"strings"
)

type InsertStmt struct {
	Options     []InsertOption
	Table       *common.Table
	Partitions  []string
	InsertCol   []ast.ColRefNode
	InsertValue []string // 结构待调整
	DupCol      []ast.ColRefNode
	DupValue    []string
}

func (stmt *InsertStmt) String() string {
	res := "INSERT "
	if stmt.Options != nil && len(stmt.Options) > 0 {
		delOptDict := []string{"IGNORE", "HIGH_PRIORITY", "LOW_PRIORITY", "DELAYED"}
		optionStrList := make([]string, 0)
		for _, opt := range stmt.Options {
			optionStrList = append(optionStrList, delOptDict[opt])
		}
		res += strings.Join(optionStrList, " ")
		res += " "
	}
	res += "INTO " + stmt.Table.Name + " "
	if stmt.Partitions != nil && len(stmt.Partitions) > 0 {
		res += "PARTITION(" + strings.Join(stmt.Partitions, ",") + ") "
	}
	res += "("
	for i, col := range stmt.InsertCol {
		if i != 0 {
			res += ", "
		}
		res += col.String()
	}
	res += ") VALUES "
	c := 0
	for j, val := range stmt.InsertValue {
		c++
		if c == 1 {
			res += "("
		}
		if j != 0 {
			res += ", "
		}
		res += val
		if c == len(stmt.InsertCol) {
			res += ")"
		}
	}
	res += " "
	if stmt.DupCol != nil && len(stmt.DupCol) > 0 {
		res += "ON DUPLICATE KEY UPDATE "
		for k, dup := range stmt.DupValue {
			if k != 0 {
				res += ", "
			}
			res += stmt.DupCol[k].String() + " = " + dup
		}
		res += " "
	}
	return res
}

type InsertOption = int

// Maybe should change to struct
const (
	InsOptIgnore = iota
	InsOptHighPriority
	InsOptLowPriority
	//InsOptDelayed
)
