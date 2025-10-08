package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	reader := bufio.NewReader(c)
	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			log.Fatal(err)
			continue
		}

		request := string(bytes)
		response := PrimeHandler(request)
		c.Write(response)
	}
}
