package statement

import (
	"dbkit/internal/common"
	"dbkit/internal/common/ast"
)

type CreateIndexStmt struct { //lack index_option
	OptionCreate    CreateOption
	IndexName       string
	TypeIndex       IndexType
	TableName       string
	KeyPart         ast.AstNode
	Columns         []*common.Column
	OptionAlgorithm AlgorithmOption
	OptionLock      LockOption
	Where           ast.AstNode
}

type CreateOption = int

const (
	CreOptFullText = iota
	CreOptSpatial
	CreOptUnique
)

type IndexType = int

const (
	IndexBtree = iota
	IndexHash
)

type AlgorithmOption = int

const (
	AlgorCopy = iota
	AlgorDefault
	AlgorInplace
)

type LockOption = int

const (
	LockDefault = iota
	LockExclusive
	LockNone
	LockShared
)

func (stmt *CreateIndexStmt) String() string {
	sql := "CREATE"
	if stmt.OptionCreate != -1 {
		creOptDic := []string{" FULLTEXT ", " SPATIAL ", " UNIQUE "}
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
	sql += "("

	colCount := len(stmt.Columns) - 1
	for idx, val := range stmt.Columns {
		sql += val.Name
		if val.Type.IsString() {
			if stmt.KeyPart == nil {
				sql += "(10)" //TODO support random length
			} else {
				sql += stmt.KeyPart.String()
			}
		} else {
			if stmt.KeyPart != nil {
				sql += stmt.KeyPart.String()
			}
		}
		if idx != colCount {
			sql += ", "
		}
	}

	sql += ")"
	if stmt.OptionAlgorithm != -1 {
		algOptDic := []string{" COPY ", " DEFAULT ", " INPLACE "}
		sql += " ALGORITHM =" + algOptDic[stmt.OptionAlgorithm]
	}
	if stmt.OptionLock != -1 {
		lockOptDic := []string{" DEFAULT ", " EXCLUSIVE ", " NONE ", " SHARED "}
		sql += " LOCK =" + lockOptDic[stmt.OptionLock]
	}

	if stmt.Where != nil {
		sql += "WHERE " + stmt.Where.String() + " "
	}

	return sql
}
