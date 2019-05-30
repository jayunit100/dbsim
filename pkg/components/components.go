package components

type Message interface {
	GetMessage() string
}

type PublisherHandle interface {
	CanPublish() bool

}

type Topic interface {

	// Each subscriber is a function which succeeds/fails
	GetSubscribers() map[string]func(m Message) bool

	// Each publisher is a named entity which is registered
	// the boolean flag represents its last publish event time
	GetPublishers() map[string]int64
	AddPublisher(name string) bool
	AddSubscriber(name string, action func(m Message) bool)

	// Delete functionality
	DeleteSub(name string)
	DeletePub(name string)

	lazyInit()
}


// default implementation of our message queue.  Does nothing yet.
type DefaultTopic struct {
	Sub map[string] func(m Message) bool
	Pub map[string] int64
	Name string
}


func (m *DefaultTopic) lazyInit() {
	if m.Name == "" {
		panic("no name set for this topic !")
	}
	if m.Sub == nil {
		m.Sub = map[string]func(m Message) bool{}
		m.Pub = map[string]int64{}
	}
}

func (m *DefaultTopic) AddSubscriber(name string, action func(m Message) bool) {
	m.lazyInit()

	m.Sub[name] = action
}

func (m *DefaultTopic) DeletePub(name string) bool {
	delete (m.Pub, name)
	return true
}

func (m *DefaultTopic) DeleteSub(name string) bool {
	delete (m.Sub, name)
	return true
}

func (m *DefaultTopic) AddPublisher(name string) {
	m.lazyInit()
	m.Pub[name] = -1
}

func (m *DefaultTopic) GetSubscribers() map[string]func(m Message)bool  {
	return m.Sub
}

func (m *DefaultTopic) GetPublishers() map[string]int64 {
	m.lazyInit()
	return m.Pub
}

func (m *DefaultTopic) Start() {
	for {
		select {
			case msg1 := <-m.Event:
			    fmt.Println("received", msg1)
			case msg2 := <-m.Event:
			    fmt.Println("received", msg2)
		}
        }
}
