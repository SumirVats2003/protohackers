package main

import (
	"log"
	"net"

	"github.com/SumirVats2003/protohackers/means-to-an-end/internal"
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
		go internal.HandleConnection(conn)
	}
}
