package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Body string `json:"body"`
}

func handleMessage(write http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	msg := Message{}
	err := decoder.Decode(&msg)

	if err != nil {
		fmt.Printf("error")
		return
	}

	fmt.Printf("Received msg: %s\n", msg.Body)
}
