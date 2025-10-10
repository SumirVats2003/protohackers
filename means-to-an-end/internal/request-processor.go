package internal

import (
	"encoding/binary"
	"log"
	"math"
	"net"
)

func ProcessRequest(command string, bytes []byte, conn net.Conn) {
	a, b := parseCommand(bytes)
	dataStore := InitDataStore()
	if command == "I" {
		log.Printf("writing at timestamp %v price %v\n", a, b)
		dataStore.Insert(a, b)
	} else {
		result := dataStore.GetAvg(a, b)
		log.Printf("got average from timestamp %v to %v value %f\n", a, b, result)

		bits := math.Float64bits(result)
		byteArrAlt := make([]byte, 8)
		binary.BigEndian.PutUint64(byteArrAlt, bits)
		conn.Write(byteArrAlt)
	}
}

func parseCommand(bytes []byte) (int32, int32) {
	first := binary.BigEndian.Uint32(bytes[:4])
	second := binary.BigEndian.Uint32(bytes[4:])
	return int32(first), int32(second)
}
