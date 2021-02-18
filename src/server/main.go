package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"runtime"
)

// main http server
func main() {
	// implement multicore
	runtime.GOMAXPROCS(runtime.NumCPU())

	// http.HandleFunc("/", httpHandler)
	http.Handle("/game", websocket.Handler(httpHandler))
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println("Error -", err)
	}

}
