package dns

import (
	"crypto/tls"
	"log"
)

func TCPDNS(serverAddr string, query []byte) ([]byte, error) {
	lenDNSPacket := 1024

	conn, err := tls.Dial("tcp", serverAddr, &tls.Config{ InsecureSkipVerify: true})

	if err != nil {
		return nil, err
	}

	conn.Write(query)
	resp := make([]byte, lenDNSPacket)
	bytesReceived, err := conn.Read(resp)

	log.Println("sent query on to server", serverAddr)
	conn.Close()

	log.Println("received response from server", serverAddr, bytesReceived, "bytes")

	return resp, nil
}
