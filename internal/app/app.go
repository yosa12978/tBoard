package app

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	api "github.com/yosa12978/tBoard/internal/pkg/web"
)

func Run() {
	listener, err := net.Listen("tcp", os.Getenv("ADDR"))
	if err != nil {
		panic(err)
	}
	go api.Listen(listener)
	out := make(chan os.Signal, 1)
	signal.Notify(out, os.Interrupt, syscall.SIGTERM)
	<-out
}
