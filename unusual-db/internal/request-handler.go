package internal

import (
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
		c.Write([]byte(dataStore.Store[key]))
	}
}
