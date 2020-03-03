package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/carldanley/nats-listener/src/config"
	"github.com/nats-io/nats.go"
)

var signalChannel chan os.Signal

func main() {
	signalChannel = make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT)

	cfg, _ := config.GetConfig()
	conn, err := nats.Connect(
		cfg.Server,
		nats.ReconnectWait(time.Second*10),
		nats.DisconnectErrHandler(func(conn *nats.Conn, err error) {
			fmt.Printf("Disconnected from NATS: %s\n", cfg.Server)
		}),
		nats.ReconnectHandler(func(conn *nats.Conn) {
			fmt.Printf("Reconnected to NATS: %s\n", conn.ConnectedUrl())
		}),
	)

	defer conn.Close()

	if err != nil {
		fmt.Printf("%v; %s\n", err, cfg.Server)
		os.Exit(1)
	}

	conn.Subscribe(cfg.Subject, func(m *nats.Msg) {
		fmt.Printf("RCVD: SUB: %s MSG: %s\n", string(m.Subject), string(m.Data))
	})

	<-signalChannel
}
