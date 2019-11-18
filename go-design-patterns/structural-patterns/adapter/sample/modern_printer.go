package main

import "fmt"

type ModernPrinter interface {
	PrintStored() string
}

// the Adapter pattern is a good design
// when two interfaces that are incompatible, but which must work together

// keep in mind that the Adapter pattern must ideally just provide
// the way to use the old LegacyPrinter and nothing else
type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg string
}

func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}

	return
}
