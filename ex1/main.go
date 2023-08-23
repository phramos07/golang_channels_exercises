package main

import "fmt"

const (
	N_USERS = 3
)

/*
Implement a messaging application with the following components:

User Struct: Each user has a unique ID and a channel for sending and receiving messages.
Central Channel: A central channel where all users send their messages.
Message Struct: A simple struct containing the sender's ID and the message content.
Your task is to create a program that demonstrates the communication between multiple users and the central channel using Go channels and the select statement.

Steps:

- Define the User struct with an ID and a channel for sending messages.
- Create a central channel for sending messages.
- Implement a function that simulates a user sending a message to the central channel.
- Implement a function that simulates the central channel distributing messages to all users.
- Use the select statement to coordinate sending and receiving messages.
*/

type User struct {
	ID      int
	MsgChan chan *Message
}

func (u *User) SendMessage(centralChan chan<- *Message) {
	message := &Message{
		FromUserID: u.ID,
		Content:    fmt.Sprintf("Message from user %d", u.ID),
	}

	// TODO : send message to channel
}

func (u *User) ReceiveMessage() {
	// TODO: listen to user channel and print message received
}

type Message struct {
	FromUserID int
	Content    string
}

// This function will listen to the central channel and send every message
// received on the central channel to every user
func handleCentralChannel(users []*User, centralChan <-chan *Message) {
	// listens to channel and sends to the message to other users.
	// Don't send the same message to the same user that created it

	// TODO

	// ...
}

func main() {

	// example of central channel
	centralChan := make(chan *Message)

	// example of how to write to a channel. Each user will write a message to a channel
	centralChan <- &Message{}

	// TODO:
	// - Create N_USERS
	// - dispatch goroutines to send/receive channels
	// - dispatch goroutine to distribute messages with central channel handler
	// - test
}
