package transport

import (
	"io"
	"net"
)

func UdpWriter(address string) (io.WriteCloser, error) {
	addr, err := net.ResolveUDPAddr("udp", address) // "127.0.0.1:12201"
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
