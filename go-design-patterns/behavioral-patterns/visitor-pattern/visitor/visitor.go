package visitor

import "io"

type MessageA struct {
	Msg    string
	Output io.Writer
}

type MessageB struct {
	Msg    string
	Output io.Writer
}

type Visitor interface {
	VisitA(*MessageA)
	VisitB(*MessageB)
}

type Visitable interface {
	Accept(Visitor)
}

type MessageVisitor struct{}
