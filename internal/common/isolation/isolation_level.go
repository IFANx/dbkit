package isolation

import (
	"dbkit/internal/randomly"
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

func GetRandomIsolationLevel() IsolationLevel {
	levels := []IsolationLevel{ReadUncommitted, ReadCommitted, RepeatableRead, Serializable}
	idx := randomly.RandIntGap(0, 3)
	return levels[idx]
}
