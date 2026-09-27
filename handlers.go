package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct{
	Username string `json:"username"`
}

type Message struct {
	Body string `json:"body"`
}

func requestData( request *http.Request, T any ) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(T)

	if err != nil {
		fmt.Printf("error requesting data: %s\n", err)
		return
	}
}

func (cfg *Config) handleMessage(write http.ResponseWriter, request *http.Request) {
	msg := Message{}
	requestData( request, &msg )

	//save chats
	_, err := cfg.database.CreateChatlog( request.Context(), msg.Body )

	if err != nil{
		fmt.Printf( "issue saving chat log: %s\n", err )
	}
}

func (cfg *Config) handleCreateLogin( write http.ResponseWriter, request *http.Request ){
	type login_param struct{
		Username string `json:username`
		Password string `json:password`
	}

	login := login_param{}
	requestData( request, &login )

	//add user to database
	_, err := cfg.database.CreateUser( request.Context(), login.Username )

	if err != nil{
		fmt.Printf( "issue saving user login: %s\n", err )
		return
	}

	fmt.Printf( "User logged in: %s", login.Username )
}
