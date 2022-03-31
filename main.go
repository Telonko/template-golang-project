package main

import (
	log "github.com/sirupsen/logrus"
	"homespace/cmd"
	"os"
	"os/signal"
	"syscall"
)

var version string // set by the compiler

func main() {
	cmd.Execute()

	log.Debugln("pid", os.Getpid())

	ch := make(chan os.Signal)
	signal.Ignore(syscall.SIGPIPE, syscall.SIGALRM)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGABRT, syscall.SIGKILL)
	<-ch

	log.Info("Process is closed")
}
