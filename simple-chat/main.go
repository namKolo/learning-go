package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", handleWebsocket())
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func handleWebsocket() http.HandlerFunc {
	h := newHub()
	go h.run()
	return wsHandler(h)
}
