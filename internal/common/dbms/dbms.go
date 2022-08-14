package dbms

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

type DBMS struct {
	Name  string
	Alias string
}

func (dbms DBMS) String() string {
	return dbms.Name
}

var (
	MYSQL       = DBMS{MySQLName, MySQLAlias}
	MARIADB     = DBMS{MARIADBName, DBMARIADBAlias}
	TIDB        = DBMS{TiDBName, DBTiDBAlias}
	COCKROACHDB = DBMS{CockroachName, DBCockroachDBAlias}
	ZNBASE      = DBMS{ZNBaseName, DBZNBaseAlias}
	SQLITE      = DBMS{SQLiteName, DBSQLiteAlias}
	DAMENG      = DBMS{DAMENGName, DBDAMENGAlias}
)

var DBMSMap = map[string]DBMS{
	MySQLAlias:         MYSQL,
	DBMARIADBAlias:     MARIADB,
	DBTiDBAlias:        TIDB,
	DBCockroachDBAlias: COCKROACHDB,
	DBZNBaseAlias:      ZNBASE,
	DBSQLiteAlias:      SQLITE,
	DBDAMENGAlias:      DAMENG,
}

func GetDBMSFromStr(dbms string) DBMS {
	dbms = strings.ToLower(dbms)
	val, ok := DBMSMap[dbms]
	if !ok {
		log.Errorf("Do not support %s", dbms)
	}
	return val
}
