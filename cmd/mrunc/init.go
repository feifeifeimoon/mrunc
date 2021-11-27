package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
)

func init() {

	if len(os.Args) > 1 && os.Args[1] == "Init" {

		runtime.GOMAXPROCS(0)
		runtime.LockOSThread()
		log.Info("init process running..")

		os.Exit(0)
	}

}
