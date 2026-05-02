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
		dataStore.Store[key] = value
	} else {
		value, ok := dataStore.Store[key]

		if ok {
			log.Printf("writing %v to the connection", value)
			c.Write([]byte(value))
		}
	}
}
