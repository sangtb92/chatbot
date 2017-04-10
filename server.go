package main

import (
	"net"
	"fmt"
	"encoding/gob"
)

func server()  {
	ln, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	fmt.Println("Server start listening on port 8000")
	for  {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleServerConnection(c)
	}
}
func handleServerConnection(conn net.Conn) {
	var msg string
	err := gob.NewDecoder(conn).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received: ", msg)
	}
	conn.Close()
}

func main() {
	server()
}