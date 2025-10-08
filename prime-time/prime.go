package main

import (
	"encoding/json"
	"errors"
	"log"
	"math"
)

type Request struct {
	Method string
	Number float64
}

type Response struct {
	Method  string `json:"method"`
	IsPrime bool   `json:"prime"`
}

type MalformedResponse struct {
	Foo string
}

func PrimeHandler(request string) []byte {
	log.Println("handling request: ", request)
	req, err := getValidRequest(request)
	if err != nil {
		mr := MalformedResponse{"hello"}
		res, _ := json.Marshal(mr)
		return []byte(res)
	}

	res := getProcessedResponse(req)
	jsonResponse, err := json.Marshal(res)
	if err != nil {
		mr := MalformedResponse{"hello"}
		res, _ := json.Marshal(mr)
		return []byte(res)
	}
	return []byte(jsonResponse)
}

func getValidRequest(request string) (Request, error) {
	var requestObject map[string]any
	json.Unmarshal([]byte(request), &requestObject)

	method := requestObject["method"]
	number := requestObject["number"]

	if method == nil || method != "isPrime" || number == nil {
		return Request{}, errors.New("Malformed Request")
	}

	var req Request
	json.Unmarshal([]byte(request), &req)
	return req, nil
}

func getProcessedResponse(req Request) Response {
	var isPrime bool
	if req.Number == math.Trunc(req.Number) && req.Number >= 0 {
		num := int(req.Number)
		isPrime = num%2 == 0
	} else {
		isPrime = false
	}
	return Response{
		Method:  req.Method,
		IsPrime: isPrime,
	}
}
