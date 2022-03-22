package txn

import (
	"dbkit/internal"
	log "github.com/sirupsen/logrus"
)

type IsolationLevel struct {
	Name  string
	Alias string
}

var (
	ReadUncommitted = IsolationLevel{Name: IsolationRUName, Alias: IsolationRUAlias}
	ReadCommitted   = IsolationLevel{Name: IsolationRCName, Alias: IsolationRCAlias}
	RepeatableRead  = IsolationLevel{Name: IsolationRRName, Alias: IsolationRRAlias}
	Serializable    = IsolationLevel{Name: IsolationSERName, Alias: IsolationSERAlias}
)

func GetIsolationFromAlias(alias string) IsolationLevel {
	switch alias {
	case IsolationRUAlias:
		return ReadUncommitted
	case IsolationRCAlias:
		return ReadCommitted
	case IsolationRRAlias:
		return RepeatableRead
	case IsolationSERAlias:
		return Serializable
	default:
		log.Infof("Unsupported isolation alias: %s", alias)
		panic("Unreachable")
	}
}

func GetAllowedIsolationOfDBMS(dbms internal.DBMS) []IsolationLevel {
	switch dbms {
	case internal.MYSQL:
		return []IsolationLevel{ReadUncommitted, ReadCommitted, RepeatableRead, Serializable}
	case internal.TIDB:
		return []IsolationLevel{ReadCommitted, RepeatableRead}
	case internal.COCKROACHDB:
		return []IsolationLevel{Serializable}
	case internal.ZNBASE:
		return []IsolationLevel{ReadCommitted, Serializable}
	case internal.SQLITE:
		return []IsolationLevel{Serializable}
	default:
		log.Infof("Unsupported isolation alias: %s", dbms)
		panic("Unreachable")
	}
}
