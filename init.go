package main

import (
	"github.com/feifeifeimoon/mrunc/pkg/process"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
)

func init() {

	if len(os.Args) > 1 && os.Args[1] == "Init" {

		runtime.GOMAXPROCS(0)
		runtime.LockOSThread()
		log.Info("init process running..")

		log.SetLevel(log.DebugLevel)
		log.SetFormatter(&log.JSONFormatter{})
		// todo log fd
		//pipeFd, err := strconv.Atoi(os.Getenv("INIT_PIPE_PD"))
		//if err != nil {
		//	panic(err)
		//}

		p := process.NewInitProcess()
		p.StartInitialization()

		os.Exit(0)
	}

}
