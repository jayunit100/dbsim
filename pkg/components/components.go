package components

type Message interface {

}

type MQ interface {
	// Each subscriber is a function which succeeds/fails
	GetSubscriptions() map[string]func(m Message) bool
	// Each publisher is a named entity which is registered
	// the boolean flag represents its last publish event time
	GetPublishers() map[string]int64
}

// default implementation of our message queue.  Does nothing yet.
type DefaultMQ struct {

}


func (m *DefaultMQ) GetSubscriptions() map[string]func(m Message)bool  {
	return nil
}

func (m *DefaultMQ) GetPublishers() map[string]int64 {
	return nil
}

