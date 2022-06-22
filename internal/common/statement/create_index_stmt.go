package statement

import (
	"strings"
)

type CreateIndexStmt struct { //lack index_option
	OptionCreate    CreateOption
	IndexName       string
	TypeIndex       IndexType
	TableName       string
	Columns         []string //incomplete
	OptionAlgorithm AlgorithmOption
	OptionLock      LockOption
}

type CreateOption = int

const (
	CreOptUnique = iota
	CreOptFullText
	CreOptSpatial
)

type IndexType = int

const (
	IndexBtree = iota
	IndexHash
)

type AlgorithmOption = int

const (
	AlgorDefault = iota
	AlgorInplace
	AlgorCopy
)

type LockOption = int

const (
	LockDefault = iota
	LockNone
	LockShared
	LockExclusive
)

func (stmt *CreateIndexStmt) String() string {
	sql := "CREATE"
	if stmt.OptionCreate != -1 {
		creOptDic := []string{" UNIQUE ", " FULLTEXT ", " SPATIAL "}
		sql += creOptDic[stmt.OptionCreate]
	}
	sql += "INDEX "
	sql += stmt.IndexName
	if stmt.TypeIndex != -1 {
		indexTypeDic := []string{" BTREE ", " HASH "}
		sql += " USING" + indexTypeDic[stmt.TypeIndex]
	}
	sql += "ON "
	sql += stmt.TableName
	sql += "(" + strings.Join(stmt.Columns, ",") + ")"
	if stmt.OptionAlgorithm != -1 {
		algOptDic := []string{" DEFAULT ", " INPLACE ", " COPY "}
		sql += " ALGORITHM =" + algOptDic[stmt.OptionAlgorithm]
	}
	if stmt.OptionLock != -1 {
		lockOptDic := []string{" DEFAULT ", " NONE ", " SHARED ", " EXCLUSIVE "}
		sql += " LOCK =" + lockOptDic[stmt.OptionLock]
	}

	return sql
}
