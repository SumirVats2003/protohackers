package internal

import (
	"bufio"
	"io"
	"log"
	"net"
)

func Serve() {
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
		query, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				log.Println("Client disconnected:", c.RemoteAddr())
				return
			}
			log.Println("Read error:", err)
			return
		}

		log.Println(string(query))
	}
}
