package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"
	"github.com/chobito/n26_test/config"
	"github.com/chobito/n26_test/handler"
	"github.com/chobito/n26_test/server"
	"github.com/miekg/dns"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM)
	signal.Notify(signalChan, syscall.SIGINT)

	cfg := config.Config{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal("Failed to parse config. Exiting...")
	}
	c := new(dns.Client)
	c.Net = "tcp-tls"
	c.Dialer = &net.Dialer{
		Timeout: cfg.UpstreamTimeout,
	}
	h := handler.NewHandler(c, cfg)
	server.StartServers(cfg)
	dns.Handle(".", h)

	sig := <-signalChan
	log.Printf("Received signal: %q, shutting down..", sig.String())
	server.ShutdownServers()
}
