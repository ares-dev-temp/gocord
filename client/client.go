package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type User struct {
	Username string `json:"username"`
}

type Message struct {
	Body string `json:"body"`
}

func request( url string, T any ){
	jsonData, _ := json.Marshal(T)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Printf("erorr creating request: %s\n", err)
		return
	}

	// setting a header on a the new request
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("failed to speak/connect to server: %s\n", err)
		return
	}

	defer res.Body.Close()

	//var createdMsg Message
	var createdData any
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&createdData)
	if err != nil {
		return
	}
}

func main() {
	url_dir := "http://localhost:8080/api"
	scanner := bufio.NewScanner(os.Stdin)

	//prompts user to login a username
	fmt.Print("Enter username: ")

	scanner.Scan()
	text_out := scanner.Text()
	username := User{Username: text_out}
	request( url_dir + "/login", username )

	//prompts user to send messages
	for {
		fmt.Print("GoCord > ")

		scanner.Scan()
		text_out := scanner.Text()
		msg := Message{Body: text_out}
		request(url_dir + "/message", msg)
	}

}
