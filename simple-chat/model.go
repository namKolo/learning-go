package main

// Message sending message between client and server
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

type NumOfMessage struct {
	Count int `json:"count"`
}
