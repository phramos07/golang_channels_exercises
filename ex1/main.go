package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
	MsgChan chan *Message // this is where user receives message
}

func (u *User) SendMessage(centralChan chan<- *Message) {
	for {
		message := &Message{
			FromUserID: u.ID,
			Content:    fmt.Sprintf("Hello from user %d\n", u.ID),
		}

		userSleep()
		// TODO : send message to channel
		centralChan <- message
	}
}

func (u *User) ReceiveMessage() {
	for {
		message := <-u.MsgChan // broadcast func will write to that channel
		fmt.Printf("\nUser %d received message from user %d: %s", u.ID, message.FromUserID, message.Content)
	}
}

func userSleep() {
	// simulate the passing of time (between 5 and 10 seconds)
	rand.Seed(time.Now().UnixNano())
	min := 5
	max := 10
	// rand.Intn(n) -> [0,n) -> [5, 11)
	// max - min + 1 == 6
	// intn(6) -> [0, 6) (from 0 to 5)
	// + min (5) -> from 5 to 10
	sleepDuration := time.Duration(rand.Intn(max-min+1) + min)
	time.Sleep(sleepDuration * time.Second)
}

type Message struct {
	FromUserID int
	Content    string
}

// This function will listen to the central channel and send every message
// received on the central channel to every user
func broadcast(users []*User, centralChan <-chan *Message) {
	// listens to channel and sends the message to other users.
	// Don't send the same message to the same user that created it

	// TODO
	for {
		select {
		case message := <-centralChan:
			for _, user := range users {
				if user.ID != message.FromUserID {
					user.MsgChan <- message
				}
			}
		}
	}
}

func main() {

	// example of central channel
	centralChan := make(chan *Message)

	// TODO:
	// - Create N_USERS

	users := make([]*User, N_USERS)
	for i := 0; i < N_USERS; i++ {
		users[i] = &User{ID: i + 1, MsgChan: make(chan *Message)}
		// - dispatch goroutines to send/receive channels // using the word 'go'
		go users[i].SendMessage(centralChan)
		go users[i].ReceiveMessage()
	}

	// - dispatch goroutine to broadcast messages that arrive to the central channel
	go broadcast(users, centralChan)

	// - test
	for {
	}
}
