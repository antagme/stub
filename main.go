package main

import (
	"flag"
	"fmt"
	"github.com/antagme/stub/internal/server"
	"log"
	"sync"
)

const defaultServerAddr = "1.1.1.1:853"
const defaultListenAddr = ":53"
const defaultProtocol = "all"

func main() {
	var waitGroup sync.WaitGroup
	flagListenAddr := flag.String("listen", defaultListenAddr, "DoT Provider ip:port")
	flagServerAddr := flag.String("server", defaultServerAddr, "Internal port to run DNS Proxy: :port")
	flagPrintHelp := flag.Bool("help", false, "print this help")
	flagProtocol := flag.String("protocol", defaultProtocol, "tcp / udp / all")

	flag.Usage = func() {
		fmt.Println("stub is a simple DNS over TLS proxy.")
		fmt.Println( "Usage: stub\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if *flagPrintHelp {
		flag.Usage()
		return
	}


	listenAddr := *flagListenAddr
	serverAddr := *flagServerAddr
	protocol   := *flagProtocol

	switch {

	case protocol == "tcp":

		waitGroup.Add(1)
		tcp, err := server.NewServer("tcp", listenAddr , serverAddr, waitGroup)
		if err != nil {
			log.Println("error starting TCP server")
			return
		}
		go tcp.Run()
		waitGroup.Wait()

	case protocol == "udp":

		waitGroup.Add(1)
		udp, err := server.NewServer("udp", listenAddr, serverAddr, waitGroup)

		if err != nil {
			log.Println("error starting UDP server")
			return
		}

		go udp.Run()
		waitGroup.Wait()

	case protocol == "all":

		waitGroup.Add(2)
		tcp, err := server.NewServer("tcp", listenAddr , serverAddr, waitGroup)

		if err != nil {
			log.Println("error starting TCP server")
			return
		}

		udp, err := server.NewServer("udp", listenAddr, serverAddr, waitGroup)

		if err != nil {
			log.Println("error starting UDP server")
			return
		}
		go tcp.Run()
		go udp.Run()
		waitGroup.Wait()

	default:
		flag.PrintDefaults()
		return
	}
}