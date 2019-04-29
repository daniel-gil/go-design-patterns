package singleton

import "sync"

var (
	once     sync.Once
	instance *singleton
)

type Singleton interface {
	// Increase adds one to the counter
	Increase() int
}

type singleton struct {
	counter int
}

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}

func (s *singleton) Increase() int {
	s.counter++
	return s.counter
}
