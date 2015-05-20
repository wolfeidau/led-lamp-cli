package main

import (
	"log"
	"net"
)

func handlePacket(l *net.UDPConn, data []byte, u *net.UDPAddr) {

	log.Printf("msg %x", data)

}

// ListenAndServe binds to the given address and serve requests forever.
func ListenAndServe(n, addr string) error {
	uaddr, err := net.ResolveUDPAddr(n, addr)
	if err != nil {
		return err
	}

	l, err := net.ListenUDP(n, uaddr)
	if err != nil {
		return err
	}

	buf := make([]byte, maxPktLen)
	for {
		nr, addr, err := l.ReadFromUDP(buf)
		if err == nil {
			tmp := make([]byte, nr)
			copy(tmp, buf)
			go handlePacket(l, tmp, addr)
		}
	}
}
