package main

import (
	"fmt"
	"golang.org/x/net/websocket"
)

func httpHandler(ws *websocket.Conn) {
	fmt.Println("Hello", ws)
	data := ws.Request().URL.Query().Get("data") // json
	fmt.Println("data:", data)
}