package internal

type DBMS struct {
	Name  string
	Alias string
}

func (dbms *DBMS) String() string {
	return dbms.Name
}

var (
	MYSQL       = DBMS{DBMySQLName, DBMySQLAlias}
	TIDB        = DBMS{DBTiDBName, DBTiDBAlias}
	COCKROACHDB = DBMS{DBCockroachName, DBCockroachDBAlias}
	ZNBASE      = DBMS{DBZNBaseName, DBZNBaseAlias}
)
