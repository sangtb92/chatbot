package main

import (
	"fmt"
	"net"
	"encoding/gob"
)

func client() {
	c, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := "Hello server"
	fmt.Println("Sending : ", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	client()
}