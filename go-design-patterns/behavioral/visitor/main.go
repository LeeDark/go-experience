package main

import (
	"os"

	"github.com/LeeDark/go-experience/go-design-patterns/behavioral/visitor/visitor"
)

func main() {
	msg := visitor.MessageA{
		Msg:    "Hello World",
		Output: os.Stdout,
	}

	visitor1 := &visitor.MessageVisitor{}
	visitor2 := &visitor.MsgFieldVisitorPrinter{}
	msg.Accept(visitor1)
	msg.Accept(visitor2)
	msg.Print()
}
