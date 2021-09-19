package server

import (
	"bufio"
	"github.com/antagme/stub/internal/dns"
	"log"
	"net"
)


// Run starts the TCP Server and create a new goroutine in every TCP Request
func (t *TCPServer) Run() (err error) {
	t.server, err = net.Listen("tcp", t.addr)

	if err != nil {
		return err
	}
	defer t.Close()

	for {
		c, err := t.server.Accept()
		if err != nil {
			log.Println("Accept()")
			log.Fatal(err)
		}
		go t.handleConnection(c)
	}

	t.waitGroup.Done()
	return
}

// Close shuts down the TCP Server
func (t *TCPServer) Close() (err error) {
	return t.server.Close()
}

// handleConnection is used to accept connections on
// the TCPServer and handle each of them
func (t *TCPServer) handleConnection(conn net.Conn) (err error) {
	defer conn.Close()
	for {
		lenDNSPacket := 1024
		netData := make([]byte, lenDNSPacket)
		_, err := bufio.NewReader(conn).Read(netData)
		if err != nil {
			break
		}

		// now that we have netData, let's send them to endpoint
		oc, err := dns.TCPDNS(t.externalServer, netData)

		_, err = conn.Write(oc)
		if err != nil {
			return err
		}
		log.Println("sent results to", conn.RemoteAddr())

	}
	return
}
