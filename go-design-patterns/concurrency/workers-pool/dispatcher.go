package main

import "time"

type Dispatcher interface {
	LaunchWorker(w WorkerLauncher)
	MakeRequest(Request)
	Stop()
}

type dispatcher struct {
	inCh chan Request
}

func NewDispatcher(b int) Dispatcher {
	return &dispatcher{
		inCh: make(chan Request, b),
	}
}

func (d *dispatcher) LaunchWorker(w WorkerLauncher) {
	w.LaunchWorker(d.inCh)
}

func (d *dispatcher) MakeRequest(r Request) {
	select {
		case d.inCh <- r:
		case <-time.After(5 * time.Second):
			return
	}
}

func (d *dispatcher) Stop() {
	close(d.inCh)
}
