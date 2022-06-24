package oracle

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

type Oracle struct {
	Name         string
	Alias        string
	MultiTargets bool
}

func (oracle Oracle) HasMultiTargets() bool {
	return oracle.MultiTargets
}

func (oracle Oracle) String() string {
	return oracle.Name
}

var (
	TLP     = Oracle{TLPName, TLPAlias, false}
	NOREC   = Oracle{NoRecName, NoRecAlias, false}
	DQE     = Oracle{DQEName, DQEAlias, false}
	DIFF    = Oracle{DiffName, DiffAlias, true}
	Troc    = Oracle{TrocName, TrocAlias, false}
	DIFFTXN = Oracle{DiffTxnName, DiffTxnAlias, true}
	LINEAR  = Oracle{LinearName, LinearAlias, false}
)

var OracleMap = map[string]Oracle{
	TLPAlias:     TLP,
	NoRecAlias:   NOREC,
	DQEAlias:     DQE,
	DiffAlias:    DIFF,
	TrocAlias:    Troc,
	DiffTxnAlias: DIFFTXN,
	LinearAlias:  LINEAR,
}

func GetOracleFromStr(oracle string) Oracle {
	oracle = strings.ToLower(oracle)
	val, ok := OracleMap[oracle]
	if !ok {
		log.Errorf("Do not support oracle: %s", oracle)
	}
	return val
}
