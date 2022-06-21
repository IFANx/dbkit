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
	UNIQUE = iota
	FULLTEXT
	SPATIAL
)

type IndexType = int

const (
	BTREE = iota
	HASH
)

type AlgorithmOption = int

const (
	DEFAULT = iota
	INPLACE
	COPY
)

type LockOption = int

const (
	// DEFAULT = iota
	NONE = iota + 1
	SHARED
	EXCLUSIVE
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
