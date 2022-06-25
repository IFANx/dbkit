package main

import (
	"dbkit/admin"
	log "github.com/sirupsen/logrus"
)

func main() {
	defer func() {
		log.Info("End of run, clean up resources")
		clean()
	}()
	admin.StartServer(9999)
}
