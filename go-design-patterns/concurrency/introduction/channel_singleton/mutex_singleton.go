package channel_singleton

import "sync"

type mutex_singleton struct {
	count int
	sync.RWMutex
}

var mutex_instance mutex_singleton
func GetMutexInstance() *mutex_singleton {
	return &mutex_instance
}

func (s *mutex_singleton) AddOne() {
	s.Lock()
	defer s.Unlock()
	s.count++
}

func (s *mutex_singleton) GetCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}
