package visitor

import "io"

type MessageA struct {
	Msg    string
	Output io.Writer
}

func (m *MessageA) Accept(v Visitor) {
	//Do nothing
}

func (m *MessageA) Print() {
	//Do nothing
}

type MessageB struct {
	Msg    string
	Output io.Writer
}

func (m *MessageB) Accept(v Visitor) {
	//Do nothing
}

func (m *MessageB) Print() {
	//Do nothing
}

type Visitor interface {
	VisitA(*MessageA)
	VisitB(*MessageB)
}

type Visitable interface {
	Accept(Visitor)
}

type MessageVisitor struct{}

func (mf *MessageVisitor) VisitA(m *MessageA) {
	//Do nothing
}

func (mf *MessageVisitor) VisitB(m *MessageB) {
	//Do nothing
}
