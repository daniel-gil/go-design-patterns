package singleton

type Singleton interface {
	// Increase adds one to the counter
	Increase() int
}

type singleton struct {
	counter int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

func (s *singleton) Increase() int {
	s.counter++
	return s.counter
}
