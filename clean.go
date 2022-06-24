package main

import (
	"dbkit/internal"
)

func clean() {
	logFile.Close()

	state := internal.GetState()
	state.DataSource.Close()
}
