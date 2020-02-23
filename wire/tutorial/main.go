package main

import "fmt"

/*
	Entire implement code for wire tutorial
	https://github.com/google/wire/blob/master/JJ/README.md
*/

type Message struct {
	msg string
}

type Greeter struct {
	Message Message
}

func (g Greeter) Greet() Message {
	return g.Message
}

type Event struct {
	Greeter Greeter
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// NewMessage is the constructor of Message.
func NewMessage(msg string) Message {
	return Message{msg: msg}
}

// NewGreeter is the constructor of Greeter.
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// NewEvent is the constructor of Event
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}


func main() {
	// before using wire
	// message := NewMessage("hello, I'm Rancho")
	// greeter := NewGreeter(message)
	// event := NewEvent(greeter)

	// now with wire
	event := InitializeEvent("hello, I'm Rancho")
	event.Start()

}
