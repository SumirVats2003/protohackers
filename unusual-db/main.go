package main

import (
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", ":8080")

	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	log.Println("Listening on port :8080")

	buf := make([]byte, 1024)

	for {
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Error reading:", err)
			continue
		}

		log.Printf("Received %d bytes from %s: %s\n", n, remoteAddr, string(buf[:n]))
	}
}
