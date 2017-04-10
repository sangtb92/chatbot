package main

import (
	"github.com/abhinavdahiya/go-messenger-bot"
	"log"
	"net/http"
	"github.com/dyatlov/gostardict/stardict"
	//"github.com/derekparker/delve/pkg/dwarf/reader"
	//"fmt"
	"fmt"
	"strings"
)

const (
	ACCESS_TOKEN   = "EAASFsqn9tSMBAKcLwO4iPBRkSZBWHNL0ubveADhrE3w8OLA2gnENsZCYWvOjKe1ZBpqT7y20YVcI4JBnIHJndtcMStJ8fd7PJfOxYKXKe5gxqZCqBtOVVUmaOZCeegHc9YKTPstfhHsMTs6eBr3PFBw1Qj7RiZBkyz99HN4SAvlwZDZD"
	VERIFY_TOKEN   = "sang_2201"
	APPLICATION_ID = "73f0f3025f61649f88ca3d04c7edba50"
)

func main() {
	bot := mbotapi.NewBotAPI(ACCESS_TOKEN, VERIFY_TOKEN, APPLICATION_ID)
	callbacks, mux := bot.SetWebhook("/webhook")
	go http.ListenAndServe(":8080", mux)
	for callback := range callbacks {
		if callback.IsMessage()  {
			//log.Printf("[%#v] %s", callback.Sender, callback.Message.Text)
			msg1 := goStartDict(callback.Message.Text)
			//fmt.Println(msg1)
			msg := mbotapi.NewMessage(msg1)
			bot.Send(callback.Sender, msg, mbotapi.RegularNotif)
		}


	}
}

func goStartDict(msg string) string {
	var result string
	// init dictionary with path to dictionary files and name of dictionary
	dict, err := stardict.NewDictionary("/home/sangnd/Downloads/VietAnh/AnhViet", "star_anhviet")

	if err != nil {
		panic(err)
	}
	senses := dict.Translate(msg) // get translations
	for i, sense := range senses {
		log.Println("Sense", i)
		for j, tran := range sense.Parts {
			log.Println(j)
			result += string(tran.Data)
		}
	}
	fmt.Println(result)
	rest := strings.Split(result, "*");
	fmt.Println(len(rest))
	for r := range rest {
		fmt.Println(string(r))
	}
	return result
}
