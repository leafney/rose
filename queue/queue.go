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
	mutex         sync.Mutex
	waitGroup     *sync.WaitGroup
	toNext        chan struct{}
	toNextSize    int
	toNextDelay   time.Duration
}

func NewGroupQueue(maxSize int, timeout, nextDelay time.Duration, handler func([]interface{})) *GroupQueue {
	return &GroupQueue{
		inputChannel:  make(chan interface{}, maxSize),
		outputChannel: make(chan []interface{}, 1),
		maxSize:       maxSize,
		timeout:       timeout,
		toNextDelay:   nextDelay,
		handler:       handler,
		waitGroup:     new(sync.WaitGroup),
		mutex:         sync.Mutex{},
		toNext:        make(chan struct{}, 1),
		toNextSize:    0,
	}
}

func (gq *GroupQueue) Start() {
	gq.mutex.Lock()
	defer gq.mutex.Unlock()

	// Start the timer
	if gq.ticker == nil {
		gq.ticker = time.NewTicker(gq.timeout)
	}

	go gq.processInput()
	go gq.processOutput()

}

func (gq *GroupQueue) processInput() {
	gq.waitGroup.Add(1)
	defer gq.waitGroup.Done()

	dataList := make([]interface{}, 0)
	for {
		select {
		case data, ok := <-gq.inputChannel:
			if !ok {
				gq.mutex.Lock()
				// When the input queue is closed, stop the timer.
				gq.ticker.Stop()
				// If there is currently unprocessed data, add it to the output queue and wait for processing. then close the output queue
				if len(dataList) > 0 {
					gq.outputChannel <- dataList
				}
				close(gq.outputChannel)
				gq.mutex.Unlock()
				return
			}

			gq.mutex.Lock()

			dataList = append(dataList, data)
			gq.toNextSize += 1

			// Allow next message to be received if conditions are not met
			if gq.toNextSize > 0 && gq.toNextSize != gq.maxSize {
				gq.toNext <- struct{}{}
			}

			if len(dataList) == gq.maxSize {
				gq.ticker.Stop()
				gq.outputChannel <- dataList
				dataList = make([]interface{}, 0)
			}

			gq.mutex.Unlock()
		case <-gq.ticker.C:
			if len(dataList) > 0 {
				gq.outputChannel <- dataList
				dataList = make([]interface{}, 0)
			}
		}
	}
}

func (gq *GroupQueue) processOutput() {
	gq.waitGroup.Add(1)
	defer gq.waitGroup.Done()

	for dataList := range gq.outputChannel {
		gq.toNextSize -= len(dataList)
		gq.handler(dataList)

		if gq.toNextSize == 0 {
			// Set how long to wait before executing the next round
			time.Sleep(gq.toNextDelay)

			// After waiting for message processing to complete, reset the timer
			gq.ticker.Reset(gq.timeout)
			gq.toNext <- struct{}{}
		}
	}
}

func (gq *GroupQueue) Put(data interface{}) <-chan struct{} {
	gq.inputChannel <- data

	return gq.toNext
}

func (gq *GroupQueue) Stop() {
	gq.mutex.Lock()
	defer gq.mutex.Unlock()

	gq.ticker.Stop()
	// When stopped, only the input queue is closed
	close(gq.inputChannel)

	gq.waitGroup.Wait()
}
