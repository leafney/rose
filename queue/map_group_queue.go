/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2023-06-08 15:42
 * @Description:
 */

package queue

import (
	"sync"
	"time"
)

type MapGroupQueue struct {
	maxSize           int
	timeout           time.Duration
	nextDelay         time.Duration
	groupQueueMapLock sync.Mutex
	groupQueueMap     map[string]*GroupQueue
	handlerConfigMap  map[string]*Config
}

type Config struct {
	maxSize   int
	timeout   time.Duration
	nextDelay time.Duration
}

func NewMapGroupQueue(maxSize int, timeout, nextDelay time.Duration) *MapGroupQueue {
	return &MapGroupQueue{
		maxSize:          maxSize,
		timeout:          timeout,
		nextDelay:        nextDelay,
		groupQueueMap:    make(map[string]*GroupQueue),
		handlerConfigMap: make(map[string]*Config),
	}
}

func (c *MapGroupQueue) SetConfig(handlerName string, maxSize int, timeout, nextDelay time.Duration) {
	c.groupQueueMapLock.Lock()
	defer c.groupQueueMapLock.Unlock()

	c.handlerConfigMap[handlerName] = &Config{
		maxSize:   maxSize,
		timeout:   timeout,
		nextDelay: nextDelay,
	}
}

func (c *MapGroupQueue) SetConfigMaxSize(handlerName string, maxSize int) {
	c.groupQueueMapLock.Lock()
	defer c.groupQueueMapLock.Unlock()

	c.handlerConfigMap[handlerName] = &Config{
		maxSize:   maxSize,
		timeout:   c.timeout,
		nextDelay: c.nextDelay,
	}
}

func (c *MapGroupQueue) SetConfigTimeout(handlerName string, nextDelay time.Duration) {
	c.groupQueueMapLock.Lock()
	defer c.groupQueueMapLock.Unlock()

	c.handlerConfigMap[handlerName] = &Config{
		maxSize:   c.maxSize,
		timeout:   c.timeout,
		nextDelay: nextDelay,
	}
}

func (c *MapGroupQueue) SetConfigNextDelay(handlerName string, timeout time.Duration) {
	c.groupQueueMapLock.Lock()
	defer c.groupQueueMapLock.Unlock()

	c.handlerConfigMap[handlerName] = &Config{
		maxSize:   c.maxSize,
		timeout:   timeout,
		nextDelay: c.nextDelay,
	}
}

func (c *MapGroupQueue) GetQueue(handlerName string, handlerFunc func([]interface{})) *GroupQueue {
	c.groupQueueMapLock.Lock()
	defer c.groupQueueMapLock.Unlock()

	if groupQueue, ok := c.groupQueueMap[handlerName]; ok {
		return groupQueue
	}

	// Adjust queue totals and timeouts individually
	theMaxSize := c.maxSize
	theTimeout := c.timeout
	theNextDelay := c.nextDelay

	if config, ok := c.handlerConfigMap[handlerName]; ok {
		theMaxSize = config.maxSize
		theTimeout = config.timeout
		theNextDelay = config.nextDelay
	}

	groupQueue := NewGroupQueue(theMaxSize, theTimeout, theNextDelay, handlerFunc)
	groupQueue.Start()

	c.groupQueueMap[handlerName] = groupQueue

	return groupQueue
}

func (c *MapGroupQueue) RmvQueue(handlerName string) {
	c.groupQueueMapLock.Lock()
	defer c.groupQueueMapLock.Unlock()

	delete(c.groupQueueMap, handlerName)
}

func (c *MapGroupQueue) Clear() {
	for _, queue := range c.groupQueueMap {
		queue.Stop()
	}
}
