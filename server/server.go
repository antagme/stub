package server

import (
	"log"

	"github.com/antagme/stub/config"
	"github.com/miekg/dns"
)

var serverUdp *dns.Server
var serverTcp *dns.Server

// StartServer handles the servers initialization based on the DnsConfig Struct Variables
func StartServer(c config.DnsConfig) {
	if !c.EnableTCP && !c.EnableUDP {
		log.Fatal("Neither TCP or UDP server enabled. Exiting...")
	}
	if c.EnableTCP {
		serverTcp = &dns.Server{Addr: ":53", Net: "tcp"}
		go func() {
			err := serverTcp.ListenAndServe()
			if err != nil {
				log.Panic(err)
			}
		}()
		log.Print("Started TCP server. Listening TCP/53")
	}
	if c.EnableUDP {
		serverUdp = &dns.Server{Addr: ":53", Net: "udp"}
		go func() {
			err := serverUdp.ListenAndServe()
			if err != nil {
				log.Panic(err)
			}
		}()
		log.Print("Started UDP server. Listening UDP/53")
	}
}

func ShutdownServers() {
	shutdownServer(serverTcp)
	shutdownServer(serverUdp)
}

// ShutdownServers Shutdown the Server gracefully if it is working
func shutdownServer(s *dns.Server) {
	if s == nil {
		return
	}
	if err := s.Shutdown(); err != nil {
		log.Panicf("Failed to shutdown server %s", s.Net)
	}
}
