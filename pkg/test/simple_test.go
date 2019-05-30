package test

import (
	"testing"
	"time"

	"github.com/jayunit100/kafka-sim/pkg/databases"
	log "github.com/sirupsen/logrus"
)

type Dog struct {
	DogId int
}

func (d *Dog) Id() int {
	return d.DogId
}
func TestSimpleDatabase(t *testing.T) {
	t.Log("Simple DB")
	sdb := databases.SimpleDB{}
	// reminder interfaces are usually passed as pointers b/c!
	sdb.StoreData(&Dog{DogId: 0})
	sdb.StoreData(&Dog{DogId: 1})
	sdb.StoreData(&Dog{DogId: 2})
	if sdb.TotalRecords() != 3 {
		t.Fail()
	}
	for i := 0; i < 3; i++ {
		if sdb.GetData(i) == nil {
			t.Fail()
		}
	}
	if sdb.GetData(4) != nil {
		t.Fail()
	}
}

func TestMessageQueue(t *testing.T) {
	sdb := databases.SimpleMessageQueue{
		QueueFullPolicy: databases.Block,
		MaxSize:         3,
	}

	// We need an infinite loop for a consumer here, otherwise
	// we will get a deadlock when trying to write to the blocked queue.
	go func() {
		for {
			select {
			case newRecord := <-sdb.Writes:
				log.Infof("new record %v ", newRecord)
			default:
				log.Infof("no records...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// reminder interfaces are usually passed as pointers b/c!
	sdb.StoreData(&Dog{DogId: 0})
	sdb.StoreData(&Dog{DogId: 1})
	sdb.StoreData(&Dog{DogId: 2})
	if sdb.TotalRecords() != 3 {
		t.Fail()
	}

	// This should fail
	success := sdb.StoreData(&Dog{DogId: 3})
	if success || sdb.TotalRecords() != 3 {
		t.Fatalf("%v %v", success, sdb.TotalRecords())
	}
	sdb.QueueFullPolicy = databases.Block

	for i := 0; i < 3; i++ {
		if sdb.GetData(i) == nil {
			t.Fail()
		}
	}
	if sdb.GetData(4) != nil {
		t.Fail()
	}
}
