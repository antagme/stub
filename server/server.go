package server

import (
	"errors"
	"net"
	"strings"
	"sync"
)

// Server defines the minimum contract our
// TCP and UDP server implementations must satisfy.
type Server interface {
	Run() error
	Close() error
}

// TCPServer holds the structure of our TCP
// implementation.
type TCPServer struct {
	addr   string
	server net.Listener
	externalServer string
	waitGroup sync.WaitGroup
}

// UDPServer holds the necessary structure for our
// UDP server.
type UDPServer struct {
	addr   string
	server *net.UDPConn
	externalServer string
	waitGroup sync.WaitGroup
}

// NewServer creates a new Server using given protocol
// and addr.
func NewServer(protocol, addr string, externalServer string, waitGroup sync.WaitGroup) (Server, error) {
	switch strings.ToLower(protocol) {

	case "tcp":
		return &TCPServer{
			addr: addr,
			externalServer: externalServer,
			waitGroup: waitGroup,
		}, nil

	case "udp":
		return &UDPServer{
			addr: addr,
			externalServer: externalServer,
			waitGroup: waitGroup,
		}, nil
	}

	return nil, errors.New("Invalid protocol given")
}
