/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2024-01-18 15:37
 * @Description:
 */

package queue

import "log"

type MessageQueue struct {
	queues           map[string]chan string
	defaultTopic     string
	defaultQueueSize int
	tempCache        chan string
}

func NewMessageQueue(tempCacheSize int) *MessageQueue {
	mq := &MessageQueue{
		queues:           make(map[string]chan string),
		defaultTopic:     "default",
		defaultQueueSize: 1024,
		tempCache:        make(chan string, tempCacheSize), // If the default buffer size is exceeded, a temporary buffer is used to save it to prevent loss
	}
	mq.queues[mq.defaultTopic] = make(chan string, mq.defaultQueueSize)
	return mq
}

func (mq *MessageQueue) Publish(content string, topics ...string) {
	if len(topics) == 0 {
		topics = append(topics, mq.defaultTopic)
	}

	for _, topic := range topics {
		if _, ok := mq.queues[topic]; !ok {
			mq.queues[topic] = make(chan string, mq.defaultQueueSize)
		}

		select {
		case mq.queues[topic] <- content:
		default:
			select {
			case mq.tempCache <- content:
			default:
				log.Printf("Temp cache is full, message dropped: %s", content)
			}
		}
	}

}

func (mq *MessageQueue) Consume(handler func(content string), topics ...string) {
	if len(topics) == 0 {
		topics = append(topics, mq.defaultTopic)
	}

	for _, topic := range topics {
		if _, ok := mq.queues[topic]; !ok {
			mq.queues[topic] = make(chan string, mq.defaultQueueSize)
		}

		queue := mq.queues[topic]
		go func(q chan string) {
			for {
				select {
				case msg := <-q:
					handler(msg)

					select {
					case msg := <-mq.tempCache:
						log.Printf("Message retrieved from temp cache: %v", msg)
						q <- msg
					default:

					}
				}
			}
		}(queue)
	}
}
