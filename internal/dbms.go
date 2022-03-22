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
	MARIADB     = DBMS{DBMARIADBName, DBMARIADBAlias}
	TIDB        = DBMS{DBTiDBName, DBTiDBAlias}
	COCKROACHDB = DBMS{DBCockroachName, DBCockroachDBAlias}
	ZNBASE      = DBMS{DBZNBaseName, DBZNBaseAlias}
	SQLITE      = DBMS{DBSQLiteName, DBSQLiteAlias}
)
