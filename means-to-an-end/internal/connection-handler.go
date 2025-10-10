package internal

import (
	"bufio"
	"io"
	"log"
	"net"
)

func HandleConnection(c net.Conn) {
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

		command := string(query)
		if command != "I" && command != "Q" {
			log.Println("Invalid Command")
			return
		}

		bytes := make([]byte, 8)
		_, err = io.ReadFull(reader, bytes)
		if err != nil {
			log.Println("Malformed Request")
			return
		}

		ProcessRequest(command, bytes, c)
	}
}
