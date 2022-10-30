package main

import (
	"log"
	"sync"
	"time"
)

type Task interface {
	Handler(interface{}) (bool, error)
	GetID() string
}

type task struct {
	wg *sync.WaitGroup
	id string
}

func NewTask(id string, wg *sync.WaitGroup) Task {
	return &task{
		wg: wg,
		id: id,
	}
}

func (task *task) Handler(i interface{}) (bool, error) {
	defer task.wg.Done()

	time.Sleep(3 * time.Second)

	log.Printf("task Handler: %s", task.GetID())

	return true, nil
}

func (task *task) GetID() string {
	return task.id
}
