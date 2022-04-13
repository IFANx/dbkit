package main

import (
	"dbkit/internal"
	"dbkit/internal/common"
)

func clean() {
	logFile.Close()

	state := internal.GetState()
	state.DataSource.Close()
	for _, dbms := range common.DBMSSet {
		if state.ConnStates[dbms] == 1 {
			state.Connections[dbms].Close()
		}
	}
}
