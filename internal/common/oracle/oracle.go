package oracle

import (
	"strings"

	log "github.com/sirupsen/logrus"
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
	TLP      = Oracle{TLPName, TLPAlias, false}
	PQS      = Oracle{PQSName, PQSAlias, false}
	NoREC    = Oracle{NoRECName, NoRECAlias, false}
	NoREC2   = Oracle{NoREC2Name, NoREC2Alias, false}
	DQE      = Oracle{DQEName, DQEAlias, false}
	DIFF     = Oracle{DiffName, DiffAlias, true}
	Troc     = Oracle{TrocName, TrocAlias, false}
	TrocPlus = Oracle{TrocPlusName, TrocPlusAlias, false}
	DIFFTXN  = Oracle{DiffTxnName, DiffTxnAlias, true}
	LINEAR   = Oracle{LinearName, LinearAlias, false}
)

var OracleMap = map[string]Oracle{
	TLPAlias:      TLP,
	NoRECAlias:    NoREC,
	NoREC2Alias:   NoREC2,
	DQEAlias:      DQE,
	DiffAlias:     DIFF,
	TrocAlias:     Troc,
	TrocPlusAlias: TrocPlus,
	DiffTxnAlias:  DIFFTXN,
	LinearAlias:   LINEAR,
	PQSAlias:      PQS,
}

func GetOracleFromStr(oracle string) Oracle {
	oracle = strings.ToLower(oracle)
	val, ok := OracleMap[oracle]
	if !ok {
		log.Errorf("Do not support oracle: %s", oracle)
	}
	return val
}
