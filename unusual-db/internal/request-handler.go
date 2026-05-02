package internal

import (
	"log"
	"net"
	"strings"
)

func HandleRequest(c *net.UDPConn, addr *net.UDPAddr, dataStore DataStore, request string) {
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

		var response string
		if ok {
			response = key + "=" + value
		} else {
			response = key
		}
		log.Printf("writing %q to %s", response, addr.String())
		n, err := c.WriteToUDP([]byte(response), addr)
		if err != nil {
			log.Printf("WriteToUDP error: %v", err)
		} else {
			log.Printf("sent %d bytes to %s", n, addr.String())
		}
	}
}
