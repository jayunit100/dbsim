package databases

/**
* This is a simple database written in go.
* It stores any transaction by creating an integer key for it.
* It expects clients to be able to remember the ID of the transaction.
 */
import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type QueueFullPolicy int

const (
	Delete QueueFullPolicy = 0
	Block  QueueFullPolicy = 1
)

type SimpleMessageQueueAPI interface {
	StoreData(*Storable) int
	TotalRecords() int
	GetData(i int) interface{}
}

// similar to simple DB but we store in a Queue.
type SimpleMessageQueue struct {
	MaxSize         int
	Events          [][]byte
	QueueFullPolicy QueueFullPolicy
	Writes          chan Storable
}

func (s *SimpleMessageQueue) lazyInit() {
	if s.Events == nil {
		s.Events = [][]byte{}
	}
	if s.MaxSize == 0 {
		log.Info("Max size was 0 ! Setting to default 10")
		s.MaxSize = 10
		s.QueueFullPolicy = Delete
	}

	// Allows us to block while mq is full, one at a time though.
	s.Writes = make(chan Storable, 1)
}

func (s *SimpleMessageQueue) StoreData(e Storable) bool {
	s.lazyInit()

	log.Infof("Writing ...")
	s.Writes <- e
	log.Infof("Done ...")

	if s.TotalRecords() >= s.MaxSize && s.QueueFullPolicy == Block {
		log.Info("Queue is full, cant store %v", e)
		return false
	}
	b, _ := json.Marshal(e)
	s.Events = append(s.Events, b)
	log.Infof("Total bytes %v", len(s.Events))
	return true
}

// DATA API

func (s *SimpleMessageQueue) TotalRecords() int {
	return len(s.Events)
}

func (s *SimpleMessageQueue) GetData(i int) []byte {
	if len(s.Events) < i+1 {
		log.Infof("Key miss @ %v", i)
		return nil
	}
	return s.Events[i]
}
