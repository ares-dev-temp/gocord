package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Message struct {
	Body string `json:"body"`
}

func sendMessage(url string, data Message) {
	jsonData, _ := json.Marshal(data)

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

	var createdMsg Message
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&createdMsg)
	if err != nil {
		return
	}
}

func main() {
	url := "http://localhost:8080/api/message"
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("GoCord > ")

		scanner.Scan()
		text_out := scanner.Text()

		msg := Message{Body: text_out}

		sendMessage(url, msg)

		//fmt.Printf("output >> %s\n", text_out)
	}

}
