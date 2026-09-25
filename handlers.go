package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Body string `json:"body"`
}

func (cfg *Config) handleMessage(write http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	msg := Message{}
	err := decoder.Decode(&msg)

	if err != nil {
		fmt.Printf("error")
		return
	}

	//save chats
	_, err = cfg.database.CreateChatlog( request.Context(), msg.Body )

	if err != nil{
		fmt.Printf( "issue saving chat: ", err, "\n" )
	}
}
