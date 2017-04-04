package main

import (
	"github.com/abhinavdahiya/go-messenger-bot"
	"net/http"
	"log"
)

const (
	ACCESS_TOKEN = "EAASFsqn9tSMBAKcLwO4iPBRkSZBWHNL0ubveADhrE3w8OLA2gnENsZCYWvOjKe1ZBpqT7y20YVcI4JBnIHJndtcMStJ8fd7PJfOxYKXKe5gxqZCqBtOVVUmaOZCeegHc9YKTPstfhHsMTs6eBr3PFBw1Qj7RiZBkyz99HN4SAvlwZDZD"
	VERIFY_TOKEN = "sang_2201"
	APPLICATION_ID = "73f0f3025f61649f88ca3d04c7edba50"
)


func main()  {
	bot := mbotapi.NewBotAPI(ACCESS_TOKEN, VERIFY_TOKEN, APPLICATION_ID)
	callbacks, mux := bot.SetWebhook("/webhook")
	go http.ListenAndServe(":8080", mux)
	for callback := range callbacks {
		log.Printf("[%#v] %s", callback.Sender, callback.Message.Text)

		msg := mbotapi.NewMessage(callback.Message.Text)
		bot.Send(callback.Sender, msg, mbotapi.RegularNotif)
	}
}