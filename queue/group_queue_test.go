/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2023-12-20 14:41
 * @Description:
 */

package queue

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestNewGroupQueue(t *testing.T) {

	groupQueue := NewMapGroupQueue(5, 3*time.Second, 3*time.Second)
	defer groupQueue.Clear()

	// set config
	groupQueue.SetConfig("queue", 3, time.Second*5, time.Second*6)

	//
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("info %v", i)
		log.Println("add", msg)
		<-groupQueue.GetQueue("queue", JobTask).Put(msg)
	}

	log.Println("success")
	select {}
}

func JobTask(data []interface{}) {
	for _, d := range data {
		log.Println("task", d)
	}
}
