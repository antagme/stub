package server

import (
	"errors"
	"github.com/antagme/stub/internal/dns"
	"log"
	"net"
)

// Run starts the UDP server.
func (u *UDPServer) Run() (err error) {
	laddr, err := net.ResolveUDPAddr("udp", u.addr)
	if err != nil {
		return errors.New("could not resolve UDP addr")
	}

	u.server, err = net.ListenUDP("udp", laddr)
	if err != nil {
		return errors.New("could not listen on UDP")
	}

	return u.handleConnections()
}

// handleConnections Read UDP Requests and handle them in separated goroutines
func (u *UDPServer) handleConnections() error {
	var err error
	for {
		buf := make([]byte, 2048)
		n, conn, err := u.server.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
			break
		}
		if conn == nil {
			continue
		}

		go u.handleConnection(conn, buf[:n])
	}
	u.waitGroup.Done()
	return err
}

// handleConnection Handles a UDP Request, send it to DoT Server, receive the response and send it to the client
func (u *UDPServer) handleConnection(addr *net.UDPAddr, cmd []byte) {
	resp, err := dns.DNSUDP(u.externalServer, cmd)
	if err != nil {
		log.Println("error", err)
	}

	_, err = u.server.WriteToUDP(resp, addr)

	if err != nil {
		log.Println("error", err)
	}

	log.Println("sent results to", addr)

}

// Close ensures that the UDPServer is shut down gracefully.
func (u *UDPServer) Close() error {
	return u.server.Close()
}
