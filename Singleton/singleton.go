package creational

import (
	"sync"
)

type Singleton struct {
	count int
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func (s *Singleton) Increment() int {
	s.count++
	return s.count
}
