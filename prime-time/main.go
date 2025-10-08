package main

import (
	"bufio"
	"fmt"
	"io"
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

		log.Println("--------------------------Handling new Connection---------------------------")
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	reader := bufio.NewReader(c)
	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
				return
			}
		}

		request := string(bytes)
		response := PrimeHandler(request)
		line := fmt.Sprintf("%s\n", response)

		_, err = c.Write([]byte(line))
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
