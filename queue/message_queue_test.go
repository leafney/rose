/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2024-01-18 15:43
 * @Description:
 */

package queue

import (
	"fmt"
	"testing"
	"time"
)

func TestNewMessageQueue(t *testing.T) {
	queue := NewMessageQueue(2)
	queue.Consume(func(c string) {
		t.Logf("Received message on default topic: %s\n", c)
	})

	queue.Consume(func(cc string) {
		t.Logf("Received message on default topic2: %s\n", cc)
	})

	queue.Consume(func(cc string) {
		t.Logf("Received message on default topic3: %s\n", cc)
	}, "test")

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("message %v", i)
		if i%2 == 0 {
			queue.Publish(msg, "test")
		} else {
			queue.Publish(msg)
		}
	}

	time.Sleep(6 * time.Second)
}
