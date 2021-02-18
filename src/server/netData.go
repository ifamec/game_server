package main

import (
	"fmt"
	"golang.org/x/net/websocket"
)

type NetDataConn struct {
	Connection *websocket.Conn
	StrMd5     string
}

func (n *NetDataConn) PullFromClient() {

	for {
		var content string
		err := websocket.Message.Receive(n.Connection, &content)
		if err != nil || len(content) == 0 {
			fmt.Println("Websocket.Message.Receive Error -", err)
			break
		}

		fmt.Println(content)

	}
	return
}
