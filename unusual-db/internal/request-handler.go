package internal

import (
	"log"
	"net"
	"strings"
)

func HandleRequest(c *net.UDPConn, dataStore DataStore, request string) {
	key, value, success := strings.Cut(request, "=")

	if success {
		if key == "version" {
			return
		}
		log.Printf("inserting %v : %v to dataStore", key, value)
		dataStore.Store[key] = value
		log.Printf("%v : %v", key, dataStore.Store[key])
	} else {
		value, ok := dataStore.Store[key]

		if ok {
			log.Printf("writing %v to the connection", value)
			c.Write([]byte(value))
		}
	}
}
