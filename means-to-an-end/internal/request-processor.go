package internal

import (
	"encoding/binary"
)

func ProcessRequest(command string, bytes []byte) {
	a, b := parseCommand(bytes)
	dataStore := InitDataStore()
	if command == "I" {
		dataStore.Insert(a, b)
	} else {
		GetAvg(a, b)
	}
}

func parseCommand(bytes []byte) (int32, int32) {
	first := binary.BigEndian.Uint32(bytes[:4])
	second := binary.BigEndian.Uint32(bytes[4:])
	return int32(first), int32(second)
}
