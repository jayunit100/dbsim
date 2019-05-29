package components

import "testing"
import log "github.com/sirupsen/logrus"
import c "github.com/jayunit100/kafka-sim/pkg/components"

func TestMessageQueue(t *testing.T) {
	log.Info("-")

	// Topics need a default name.
	mq := c.DefaultTopic{
		Name: "horses",
	}

	mq.AddSubscriber("a", func(m c.Message) bool {
		return true
	})

	// Confirm that the subscriber is there.
	if (1 != len(mq.GetSubscribers())) {
		t.Fail()
	}


}
