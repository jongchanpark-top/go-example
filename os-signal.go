package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

func test_signal() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGKILL)

	outer:
	for {
		select {
		case sig:=<- signals:
			log.Info().Msg(sig.String())
			break outer
		}
	}
	
	panic("Exiting")
}
