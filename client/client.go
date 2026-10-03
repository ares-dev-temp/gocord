package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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

func request(url string, data interface{}) *http.Response {
	jsonData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Printf("erorr creating request: %s\n", err)
		return nil
	}

	// setting a header on a the new request
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("failed to speak/connect to server: %s\n", err)
		return nil
	}

	return res
}

func login(url_dir, username string) User {
	user := User{Username: username}
	res := request(url_dir+"/login", user)

	defer res.Body.Close()

	var userData User
	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(&userData)

	if err != nil {
		fmt.Printf("error decoding user\n")
		return User{}
	}

	return userData
}

func send_message(url_dir, message string, user User) Message {
	msg := Message{Body: message, UserID: user.ID}
	res := request(url_dir+"/message", msg)

	defer res.Body.Close()

	var msgData Message
	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(&msgData)
	if err != nil {
		return Message{}
	}

	return msgData
}

func main() {
	url_dir := "http://localhost:8080/api"

	//reset database
	request(url_dir+"/reset", struct{}{})

	scanner := bufio.NewScanner(os.Stdin)

	//prompts user to login a username
	fmt.Print("Enter username: ")

	scanner.Scan()
	text_out := scanner.Text()
	user_info := login(url_dir, text_out)

	//prompts user to send messages
	for {
		fmt.Print("GoCord > ")

		scanner.Scan()
		text_out := scanner.Text()
		send_message(url_dir, text_out, user_info)
	}

}
