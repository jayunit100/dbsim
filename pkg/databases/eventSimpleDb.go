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

type Storable interface {
	Id() int
}

type SimpleDBApi interface {
	StoreData(*Storable) int
	TotalRecords() int
	GetData(i int) interface{}
}

type SimpleDB struct {
	Events  map[int][]byte
	Counter int
}

func (s *SimpleDB) lazyInit() {
	if s.Events == nil {
		s.Events = map[int][]byte{}
		s.Counter = 0
	}
}

func (s *SimpleDB) StoreData(e Storable) bool {
	s.lazyInit()

	b, _ := json.Marshal(e)
	s.Events[s.Counter] = b
	log.Infof("Total bytes %v", len(s.Events))
	s.Counter++
	return true
}

// DATA API

func (s *SimpleDB) TotalRecords() int {
	return len(s.Events)
}

func (s *SimpleDB) GetData(i int) []byte {
	if s.Events[i] == nil {
		log.Infof("Key miss @ %v", i)
		return nil
	}
	return s.Events[i]
}
