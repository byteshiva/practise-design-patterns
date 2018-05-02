package main

import (
	"fmt"
)

// Command interface has commond method Execute
type Command interface {
	Execute()
}

// ConcreteCommandA struct has pointer type Receiver
type ConcreteCommandA struct {
	receiver *Receiver
}

// Execute Method performs Action on Command A
func (c *ConcreteCommandA) Execute() {
	c.receiver.Action("CommandA")
}

// ConcreteCommandB struct has pointer type Receiver
type ConcreteCommandB struct {
	receiver *Receiver
}

// Execute Method performs Action on Command A
func (c *ConcreteCommandB) Execute() {
	c.receiver.Action("CommandB")
}

// Receiver struct definition
type Receiver struct{}

// Action merthod prints
func (r *Receiver) Action(msg string) {
	fmt.Println(msg)
}

// Invoker struct definition
type Invoker struct {
	history []Command
}

// StoreAndExecute loops thru arrays of command and execute one after another.
func (i *Invoker) StoreAndExecute(cmd Command) {
	i.history = append(i.history, cmd)
	for i, cmd := range i.history {
		fmt.Printf("history %d: ", i)
		cmd.Execute()
	}
}

func main() {
	receiver := new(Receiver)
	commandA := &ConcreteCommandA{receiver}
	commandB := &ConcreteCommandB{receiver}
	invoker := new(Invoker)
	invoker.StoreAndExecute(commandA)
	invoker.StoreAndExecute(commandB)
}
