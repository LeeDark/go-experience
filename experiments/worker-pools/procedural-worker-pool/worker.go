package main

import "log"

type WorkerLauncher interface {
	LaunchWorker(taskChannel chan Task, d Dispatcher)
}

type worker struct {
	id         int
	dispatcher Dispatcher
}

func NewWorker(id int) *worker {
	worker := &worker{
		id:         id,
		dispatcher: nil,
	}
	return worker
}

func (w *worker) LaunchWorker(taskChannel chan Task, d Dispatcher) {
	w.dispatcher = d
	w.handleTask(taskChannel)
}

func (w *worker) handleTask(taskChannel <-chan Task) {
	go func() {
		for task := range taskChannel {
			// param for worker, or param for some data, result
			log.Printf("worker handleTask: %s", task.GetID())
			_, _ = task.Handler(w)
		}
	}()
}
