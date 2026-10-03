package main

import (
	"encoding/json"
	"fmt"
	"gocord/internal/database"
	"net/http"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type Message struct {
	Body   string    `json:"body"`
	UserID uuid.UUID `json:"user_id"`
}

func requestData(request *http.Request, data interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(data)

	if err != nil {
		fmt.Printf("error requesting data: %s\n", err)
		return
	}
}

func respondWithJSON(write http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error decode: %s", err)
		return
	}
	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(code)
	write.Write(data)
}

func (cfg *Config) handleReset(write http.ResponseWriter, request *http.Request) {
	cfg.database.ResetDatabase(request.Context())
}

func (cfg *Config) handleMessage(write http.ResponseWriter, request *http.Request) {
	msg := Message{}
	requestData(request, &msg)

	//save chats
	chatlog_params := database.CreateChatlogParams{
		Message: msg.Body,
		UserID:  uuid.NullUUID{UUID: msg.UserID, Valid: true},
	}
	_, err := cfg.database.CreateChatlog(request.Context(), chatlog_params)

	if err != nil {
		fmt.Printf("issue saving chat log: %s\n", err)
	}
}

func (cfg *Config) handleCreateLogin(write http.ResponseWriter, request *http.Request) {
	type login_param struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	login := login_param{}
	requestData(request, &login)

	//add user to database
	user, err := cfg.database.CreateUser(request.Context(), login.Username)

	if err != nil {
		fmt.Printf("issue saving user login: %s\n", err)
		return
	}

	user_out := User{
		ID:       user.ID,
		Username: login.Username,
	}

	respondWithJSON(write, 201, user_out)
}
