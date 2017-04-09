package main

import (
	//"encoding/json"
	"log"
	"net/http"

	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
	"time"
)

var accessToken = "EAASFsqn9tSMBAKcLwO4iPBRkSZBWHNL0ubveADhrE3w8OLA2gnENsZCYWvOjKe1ZBpqT7y20YVcI4JBnIHJndtcMStJ8fd7PJfOxYKXKe5gxqZCqBtOVVUmaOZCeegHc9YKTPstfhHsMTs6eBr3PFBw1Qj7RiZBkyz99HN4SAvlwZDZD"
var verifyToken = os.Getenv("sang_2201")

const FacebookEndPoint = "https://graph.facebook.com/v2.6/me/messages"

type ReceivedMessage struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}

type Entry struct {
	// ID        int64       `json:"id"`
	// Time      int64       `json:"time"`
	Messaging []Messaging `json:"messaging"`
}

type Messaging struct {
	Sender    Sender    `json:"sender"`
	Recipient Recipient `json:"recipient"`
	// Timestamp int64     `json:"timestamp"`
	Message Message `json:"message"`
}

type Sender struct {
	ID string `json:"id"`
}

type Recipient struct {
	ID string `json:"id"`
}

type Message struct {
	MID  string `json:"mid"`
	Seq  int64  `json:"seq"`
	Text string `json:"text"`
}

type Payload struct {
	TemplateType string  `json:"template_type"`
	Text         string  `json:"text"`
	Buttons      Buttons `json:"buttons"`
}

type Buttons struct {
	Type  string `json:"type"`
	Url   string `json:"url"`
	Title string `json:"title"`
}

type Attachment struct {
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}

type ButtonMessageBody struct {
	Attachment Attachment `json:"attachment"`
}

type ButtonMessage struct {
	Recipient         Recipient         `json:"recipient"`
	ButtonMessageBody ButtonMessageBody `json:"message"`
}

type SendMessage struct {
	Recipient Recipient `json:"recipient"`
	Message   struct {
		Text string `json:"text"`
	} `json:"message"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/webhook", webhookHandler)
	log.Fatal(http.ListenAndServe(":8080", router))

}

func webhookHandler1(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		verifyTokenAction(w, r)
	}
	if r.Method == "POST" {
		webhookPostAction(w, r)
	}
}
func webhookPostAction(w http.ResponseWriter, r *http.Request) {
	var receivedMessage ReceivedMessage
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}
	if err = json.Unmarshal(body, &receivedMessage); err != nil {
		log.Print(err)
	}
	messagingEvents := receivedMessage.Entry[0].Messaging
	for _, event := range messagingEvents {
		senderID := event.Sender.ID
		if &event.Message != nil && event.Message.Text != "" {
			message := getReplyMessage(event.Message.Text)
			sendTextMessage(senderID, message)
		}
	}
	fmt.Fprintf(w, "Success")

}

func getReplyMessage(receivedMessage string) string {
	var message string
	receivedMessage = strings.ToUpper(receivedMessage)
	log.Print(" Received message: " + receivedMessage)

	if strings.Contains(receivedMessage, "HELLO") {
		message = "Hi, my name is Annie. Nice to meet you"
	}

	return message
}

func sendTextMessage(senderID string, text string) {
	recipient := new(Recipient)
	recipient.ID = senderID
	sendMessage := new(SendMessage)
	sendMessage.Recipient = *recipient
	sendMessage.Message.Text = text
	sendMessageBody, err := json.Marshal(sendMessage)
	if err != nil {
		log.Print(err)
	}
	req, err := http.NewRequest("POST", FacebookEndPoint, bytes.NewBuffer(sendMessageBody))
	if err != nil {
		log.Print(err)
	}
	fmt.Println("%T1", req)
	fmt.Println("%T2", err)

	values := url.Values{}
	values.Add("access_token", accessToken)
	req.URL.RawQuery = values.Encode()
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{Timeout: time.Duration(30 * time.Second)}
	res, err := client.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer res.Body.Close()
	var result map[string]interface{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
	}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Print(err)
	}
	log.Print(result)
}

func verifyTokenAction(w http.ResponseWriter, req *http.Request) {
	mode := "subscribe"

	hubMode := req.URL.Query().Get("hub.mode")
	hubVerifyToken := req.URL.Query().Get("hub.verify_token")
	challenge := req.URL.Query().Get("hub.challenge")

	if hubMode == mode && hubVerifyToken == "sang_2201" {
		fmt.Println("Validating webhook")
		log.Print("Verify token success")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(challenge))
	} else {
		w.WriteHeader(http.StatusForbidden)
		log.Print("Error: verify token failed")
		fmt.Fprintf(w, "Failed validation. Make sure the validation tokens match.")
	}
}
