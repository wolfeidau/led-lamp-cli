package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {

	saddr, err := net.ResolveUDPAddr("udp", "192.168.0.57:2534")

	if err != nil {
		log.Fatalf("unable to resolve address: %s", err)
	}

	laddr, err := net.ResolveUDPAddr("udp", ":0")

	if err != nil {
		log.Fatalf("unable to resolve address: %s", err)
	}

	conn, err := net.DialUDP("udp", laddr, saddr)

	defer conn.Close()

	for i := 0; i < 10; i++ {
		msg := strconv.Itoa(i)

		buf := []byte(msg)
		_, err := conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
	}
}
