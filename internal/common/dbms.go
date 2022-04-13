package common

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

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

var DBMSSet = []DBMS{MYSQL, MARIADB, TIDB, SQLITE}

var DBMSMap = map[string]DBMS{
	DBMySQLAlias:       MYSQL,
	DBMARIADBAlias:     MARIADB,
	DBTiDBAlias:        TIDB,
	DBCockroachDBAlias: COCKROACHDB,
	DBZNBaseAlias:      ZNBASE,
	DBSQLiteAlias:      SQLITE,
}

func GetDBMSFromStr(dbms string) DBMS {
	dbms = strings.ToLower(dbms)
	val, ok := DBMSMap[dbms]
	if !ok {
		log.Errorf("不支持数据库名称：%s", dbms)
	}
	return val
}
