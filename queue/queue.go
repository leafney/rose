/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2023-06-08 15:41
 * @Description:
 */

package queue

import (
	"sync"
	"time"
)

type GroupQueue struct {
	inputChannel  chan interface{}
	outputChannel chan []interface{}
	maxSize       int
	timeout       time.Duration
	handler       func([]interface{})
	ticker        *time.Ticker
	tickerChannel <-chan time.Time
	mutex         sync.Mutex
	waitGroup     sync.WaitGroup
}

func NewGroupQueue(maxSize int, timeout time.Duration, handler func([]interface{})) *GroupQueue {
	return &GroupQueue{
		inputChannel:  make(chan interface{}, maxSize),
		outputChannel: make(chan []interface{}, maxSize),
		maxSize:       maxSize,
		timeout:       timeout,
		handler:       handler,
		ticker:        time.NewTicker(timeout),
		tickerChannel: nil,
		mutex:         sync.Mutex{},
	}
}

func (gq *GroupQueue) Start() {
	gq.waitGroup.Add(1)
	go func() {
		defer gq.waitGroup.Done()

		var dataList []interface{}
		for {
			select {
			case data, ok := <-gq.inputChannel:
				if !ok {
					// When the input queue is closed, stop the timer.
					gq.mutex.Lock()
					gq.ticker.Stop()
					gq.mutex.Unlock()

					// If there is currently unprocessed data, add it to the output queue and wait for processing. then close the output queue
					if len(dataList) > 0 {
						gq.outputChannel <- dataList
						dataList = nil
					}
					close(gq.outputChannel)
					return
				}

				dataList = append(dataList, data)
				if len(dataList) == gq.maxSize {
					// When the queue is full, reset the timeout
					gq.mutex.Lock()
					gq.ticker.Stop()
					gq.ticker = time.NewTicker(gq.timeout)
					gq.tickerChannel = gq.ticker.C
					gq.mutex.Unlock()

					gq.outputChannel <- dataList
					dataList = nil
				}
			case <-gq.getTickerChannel():
				if len(dataList) > 0 {
					gq.outputChannel <- dataList
					dataList = nil
				}
			}
		}
	}()

	gq.waitGroup.Add(1)
	go func() {
		defer gq.waitGroup.Done()

		for dataList := range gq.outputChannel {
			gq.handler(dataList)
		}
	}()
}

func (gq *GroupQueue) Put(data interface{}) {
	gq.inputChannel <- data
}

func (gq *GroupQueue) Stop() {
	gq.mutex.Lock()
	gq.ticker.Stop()
	gq.mutex.Unlock()

	// When stopped, only the input queue is closed
	close(gq.inputChannel)

	gq.waitGroup.Wait()
}

func (gq *GroupQueue) getTickerChannel() <-chan time.Time {
	gq.mutex.Lock()
	defer gq.mutex.Unlock()

	if gq.tickerChannel == nil {
		gq.tickerChannel = gq.ticker.C
	}
	return gq.tickerChannel
}
