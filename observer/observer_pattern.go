package observer

type Observer interface {
	OnMessage(data string)
}

type Subject struct {
	observers []Observer
}

func (s *Subject) RegisterObserver(obs Observer) {
	s.observers = append(s.observers, obs)
}

func (s *Subject) NotifyAllObservers(data string) {
	for _, observer := range s.observers {
		observer.OnMessage(data)
	}
}
