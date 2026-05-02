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
		log.Printf("writing %v to the connection", dataStore.Store[key])
		c.Write([]byte(dataStore.Store[key]))
	}
}
