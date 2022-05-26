package ast

import (
	"dbkit/internal"
	"strconv"
	"strings"
)

type SelectStmt struct {
	Options    []SelectOption
	SelectExpr []AstNode
	Tables     []*internal.Table
	Join       AstNode
	JoinOn     AstNode
	Partitions []string
	Where      AstNode
	GroupBy    AstNode
	Having     AstNode
	OrderBy    AstNode
	OrderOpt   OrderOption
	Limit      int
	Offset     int
	ForOpt     ForOption
}

func (stmt *SelectStmt) String() string {
	res := "SELECT "
	selOptDict := []string{"ALL", "DISTINCT", "DISTINCTROW", "HIGH_PRIORITY", "STRAIGHT_JOIN",
		"SQL_SMALL_RESULT", "SQL_BIG_RESULT", "SQL_BUFFER_RESULT", "SQL_NO_CACHE", "SQL_CALC_FOUND_ROWS"}
	optionStrList := make([]string, 0)
	for opt := range stmt.Options {
		optionStrList = append(optionStrList, selOptDict[opt])
	}
	res += strings.Join(optionStrList, " ")
	res += " FROM "
	if stmt.Join == nil {
		res += stmt.Tables[0].Name
	} else {
		res += stmt.Join.String()
		res += " ON " + stmt.JoinOn.String()
	}
	res += " "
	if stmt.Partitions != nil && len(stmt.Partitions) > 0 {
		res += "PARTITION(" + strings.Join(stmt.Partitions, ",") + ") "
	}
	if stmt.Where != nil {
		res += "WHERE " + stmt.Where.String() + " "
	}
	if stmt.GroupBy != nil {
		res += "GROUP BY " + stmt.GroupBy.String() + " "
	}
	if stmt.Having != nil {
		res += "HAVING " + stmt.Having.String() + " "
	}
	if stmt.OrderBy != nil {
		res += "ORDER BY " + stmt.OrderBy.String() + " "
	}
	if stmt.OrderOpt > -1 {
		orderOptDict := []string{"ASC", "DESC"}
		res += orderOptDict[stmt.OrderOpt] + " "
	}
	if stmt.Limit > -1 {
		res += "LIMIT " + strconv.Itoa(stmt.Limit) + " "
	}
	if stmt.Offset > -1 {
		res += "OFFSET " + strconv.Itoa(stmt.Offset) + " "
	}
	if stmt.ForOpt > -1 {
		forOptDict := []string{"SHARE", "UPDATE", "KEY SHARE", "NO KEY UPDATE"}
		res += "FOR " + forOptDict[stmt.ForOpt]
	}
	return res
}

type SelectOption = int

// Maybe should change to struct
const (
	SelOptAll = iota
	SelOptDistinct
	SelOptDistinctrow
	SelOptHighPriority
	SelOptSqlSmallResult
	SelOptSqlBigResult
	SelOptSqlBufferResult
	SelOptSqlNoCache
	SelOptSqlCalcFoundRows
)

type OrderOption = int

const (
	OrderOptASC = iota
	OrderOptDESC
)

type ForOption = int

const (
	ForOptShare = iota
	ForOptUpdate
	ForOptKeyShare
	ForOptNoKeyUpdate
)

type JoinType = int

const (
	JoinTypeInner = iota
	JoinTypeCross
	JoinTypeStraight
	JoinTypeLeft
	JoinTypeLeftOuter
	JoinTypeRight
	JoinTypeRightOuter
	JoinTypeNatural
)

type JoinNode struct {
	JoinType JoinType
	Left     AstNode
	Right    AstNode
}

func (node *JoinNode) Name() string {
	return "Constant"
}

func (node *JoinNode) Type() NodeType {
	return NodeTypeJoin
}

func (node *JoinNode) String() string {
	joinTypeDict := []string{"INNER JOIN", "CROSS JOIN", "STRAIGHT_JOIN", "LEFT JOIN",
		"LEFT OUTER JOIN", "RIGHT JOIN", "RIGHT OUTER JOIN", "NATURAL JOIN"}
	return node.Left.String() + " " + joinTypeDict[node.JoinType] + " " + node.Right.String()
}
