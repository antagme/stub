package main

import (
	"github.com/antagme/stub/server"
	"log"
	"sync"
)

const defaultServerAddr = "1.1.1.1:853"
const defaultListenAddr = ":53"

func main() {
	var waitGroup sync.WaitGroup

	// Start the new server
	tcp, err := server.NewServer("tcp", defaultListenAddr , defaultServerAddr, waitGroup)
	if err != nil {
		log.Println("error starting TCP server")
		return
	}

	udp, err := server.NewServer("udp", defaultListenAddr, defaultServerAddr, waitGroup)
	if err != nil {
		log.Println("error starting UDP server")
		return
	}

	waitGroup.Add(2)

	// Run the servers in goroutines to stop blocking
	go tcp.Run()
	go udp.Run()

	waitGroup.Wait()
}