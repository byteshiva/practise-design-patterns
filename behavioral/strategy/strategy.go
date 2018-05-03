package main

import (
	"fmt"
)

// Context define strategy
type Context struct {
	strategy func()
}

// Execute strategy
func (c *Context) Execute() {
	c.strategy()
}

// SetStrategy sets strategy
func (c *Context) SetStrategy(strategy func()) {
	c.strategy = strategy
}

func main() {
	concreteStrategyA := func() {
		fmt.Println("concreteStrategyA()")
	}

	concreteStrategyB := func() {
		fmt.Println("concreteStrategyB()")
	}

	concreteStrategyC := func() {
		fmt.Println("concreteStrategyC()")
	}

	context := Context{concreteStrategyA}
	context.Execute()
	context.SetStrategy(concreteStrategyB)
	context.Execute()
	context.SetStrategy(concreteStrategyC)
	context.Execute()
}
