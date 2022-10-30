package main

import (
	"log"
	"time"
)

type Dispatcher interface {
	LaunchWorker(w WorkerLauncher)
	PushTask(task Task)
	Stop()
}

type dispatcher struct {
	taskChannel chan Task
	timeout     time.Duration
}

func NewDispatcher(taskBufferSize int, timeout time.Duration) Dispatcher {
	return &dispatcher{
		taskChannel: make(chan Task, taskBufferSize),
		timeout:     timeout,
	}
}

func (d *dispatcher) LaunchWorker(w WorkerLauncher) {
	w.LaunchWorker(d.taskChannel, d)
}

func (d *dispatcher) PushTask(task Task) {
	select {
	case d.taskChannel <- task:
		log.Printf("PushTask Run: %s", task.GetID())
	case <-time.After(d.timeout):
		log.Printf("PushTask Timeout: %s", task.GetID())
		return
	}
}

func (d *dispatcher) Stop() {
	close(d.taskChannel)
}
