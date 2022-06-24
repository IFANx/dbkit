package statement

type CreateIndexStmt struct { //lack index_option
	OptionCreate    CreateOption
	IndexName       string
	TypeIndex       IndexType
	TableName       string
	ColumnName      string //incomplete
	OptionAlgorithm AlgorithmOption
	OptionLock      LockOption
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
	sql += "(" + stmt.ColumnName + ")"
	if stmt.OptionAlgorithm != -1 {
		algOptDic := []string{" COPY ", " DEFAULT ", " INPLACE "}
		sql += " ALGORITHM =" + algOptDic[stmt.OptionAlgorithm]
	}
	if stmt.OptionLock != -1 {
		lockOptDic := []string{" DEFAULT ", " EXCLUSIVE ", " NONE ", " SHARED "}
		sql += " LOCK =" + lockOptDic[stmt.OptionLock]
	}

	return sql
}
