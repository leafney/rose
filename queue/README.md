## queue


### Upgrade

Support the waiting time interval for the next round of execution

For detail [queue_test.go](queue_test.go)

### GroupQueue

**Group Queue**

----

### MapGroupQueue

**Named Group Queue**

> 场景示例：nsq消息队列，将接收到的消息先暂存到一个指定容量的缓存channel中，如果channel满了则立即执行，遍历处理数据。如果channel不满则一直等待，直到达到超时时间后执行。

> Scenario example: nsq message queue, temporarily store received messages in a cache channel with a specified capacity, and execute immediately if the channel is full, and traverse and process data. If the channel is not full, wait until the timeout is reached.

----

mapGroupQueue init:

```go
	GroupQueue := queue.NewMapGroupQueue(10, 30*time.Second)
	defer groupQueue.Clear()
```

nsq `HandlerMsg`:

```go

func (c *TaskWorker) HandlerMsg(msg *rnsq.XMessage) error {
	
	data := msg.ToString()

	// 接收消息并添加到任务组
	if rose.RandInt(10)%2 == 0 {
		log.Info("[Bot] 接收到nsq消息", zap.String("msg", data), zap.String("func", "JobQueue"))
        
		// 修改初始设置
		GroupQueue.SetConfig("queue", 5, time.Second*20)
		GroupQueue.GetQueue("queue", JobQueue).Put(data)
	} else {
		log.Info("[Bot] 接收到nsq消息", zap.String("msg", data), zap.String("func", "JobTask"))

		GroupQueue.GetQueue("task", JobTask).Put(data)
	}

	return nil
}


```

`Jobs`:

```go
func JobQueue(dataList []interface{}) {
	for _, data := range dataList {
		// 处理数据
		if v, ok := data.(string); ok {
			log.Info("[JobQueue] 处理数据", zap.String("data", v))
		}
	}
}

func JobTask(dataList []interface{}) {
	for _, data := range dataList {
		// 处理数据
		if v, ok := data.(string); ok {
			log.Info("[JobTask] 处理数据", zap.String("data", v))
		}
	}
}
```

result:

```text
2023/06/08 15:47:22 [Nsq] ConsumeConcurrent success
{"level":"info","time":"2023-06-08 15:47:43.424","msg":"[Bot] 接收到nsq消息","msg":"msg-0","func":"JobQueue"}
{"level":"info","time":"2023-06-08 15:47:45.425","msg":"[Bot] 接收到nsq消息","msg":"msg-1","func":"JobQueue"}
{"level":"info","time":"2023-06-08 15:47:47.426","msg":"[Bot] 接收到nsq消息","msg":"msg-2","func":"JobQueue"}
{"level":"info","time":"2023-06-08 15:47:49.428","msg":"[Bot] 接收到nsq消息","msg":"msg-3","func":"JobQueue"}
{"level":"info","time":"2023-06-08 15:47:51.430","msg":"[Bot] 接收到nsq消息","msg":"msg-4","func":"JobTask"}
{"level":"info","time":"2023-06-08 15:47:53.431","msg":"[Bot] 接收到nsq消息","msg":"msg-5","func":"JobTask"}
{"level":"info","time":"2023-06-08 15:47:55.433","msg":"[Bot] 接收到nsq消息","msg":"msg-6","func":"JobQueue"}
{"level":"info","time":"2023-06-08 15:47:55.433","msg":"[JobQueue] 处理数据","data":"msg-0"}
{"level":"info","time":"2023-06-08 15:47:55.433","msg":"[JobQueue] 处理数据","data":"msg-1"}
{"level":"info","time":"2023-06-08 15:47:55.433","msg":"[JobQueue] 处理数据","data":"msg-2"}
{"level":"info","time":"2023-06-08 15:47:55.433","msg":"[JobQueue] 处理数据","data":"msg-3"}
{"level":"info","time":"2023-06-08 15:47:55.433","msg":"[JobQueue] 处理数据","data":"msg-6"}
{"level":"info","time":"2023-06-08 15:47:57.435","msg":"[Bot] 接收到nsq消息","msg":"msg-7","func":"JobQueue"}
{"level":"info","time":"2023-06-08 15:47:59.436","msg":"[Bot] 接收到nsq消息","msg":"msg-8","func":"JobQueue"}
{"level":"info","time":"2023-06-08 15:48:01.438","msg":"[Bot] 接收到nsq消息","msg":"msg-9","func":"JobTask"}
{"level":"info","time":"2023-06-08 15:48:03.440","msg":"[Bot] 接收到nsq消息","msg":"msg-10","func":"JobQueue"}
{"level":"info","time":"2023-06-08 15:48:05.441","msg":"[Bot] 接收到nsq消息","msg":"msg-11","func":"JobQueue"}
{"level":"info","time":"2023-06-08 15:48:07.443","msg":"[Bot] 接收到nsq消息","msg":"msg-12","func":"JobQueue"}
{"level":"info","time":"2023-06-08 15:48:07.443","msg":"[JobQueue] 处理数据","data":"msg-7"}
{"level":"info","time":"2023-06-08 15:48:07.443","msg":"[JobQueue] 处理数据","data":"msg-8"}
{"level":"info","time":"2023-06-08 15:48:07.443","msg":"[JobQueue] 处理数据","data":"msg-10"}
{"level":"info","time":"2023-06-08 15:48:07.443","msg":"[JobQueue] 处理数据","data":"msg-11"}
{"level":"info","time":"2023-06-08 15:48:07.443","msg":"[JobQueue] 处理数据","data":"msg-12"}
{"level":"info","time":"2023-06-08 15:48:09.444","msg":"[Bot] 接收到nsq消息","msg":"msg-13","func":"JobTask"}
{"level":"info","time":"2023-06-08 15:48:11.447","msg":"[Bot] 接收到nsq消息","msg":"msg-14","func":"JobQueue"}
{"level":"info","time":"2023-06-08 15:48:13.449","msg":"[Bot] 接收到nsq消息","msg":"msg-15","func":"JobTask"}
{"level":"info","time":"2023-06-08 15:48:15.452","msg":"[Bot] 接收到nsq消息","msg":"msg-16","func":"JobTask"}
{"level":"info","time":"2023-06-08 15:48:17.454","msg":"[Bot] 接收到nsq消息","msg":"msg-17","func":"JobTask"}
{"level":"info","time":"2023-06-08 15:48:19.456","msg":"[Bot] 接收到nsq消息","msg":"msg-18","func":"JobQueue"}
{"level":"info","time":"2023-06-08 15:48:21.436","msg":"[JobTask] 处理数据","data":"msg-4"}
{"level":"info","time":"2023-06-08 15:48:21.436","msg":"[JobTask] 处理数据","data":"msg-5"}
{"level":"info","time":"2023-06-08 15:48:21.436","msg":"[JobTask] 处理数据","data":"msg-9"}
{"level":"info","time":"2023-06-08 15:48:21.436","msg":"[JobTask] 处理数据","data":"msg-13"}
{"level":"info","time":"2023-06-08 15:48:21.436","msg":"[JobTask] 处理数据","data":"msg-15"}
{"level":"info","time":"2023-06-08 15:48:21.436","msg":"[JobTask] 处理数据","data":"msg-16"}
{"level":"info","time":"2023-06-08 15:48:21.436","msg":"[JobTask] 处理数据","data":"msg-17"}
{"level":"info","time":"2023-06-08 15:48:21.458","msg":"[Bot] 接收到nsq消息","msg":"msg-19","func":"JobTask"}
{"level":"info","time":"2023-06-08 15:48:23.460","msg":"[Bot] 接收到nsq消息","msg":"msg-20","func":"JobTask"}
{"level":"info","time":"2023-06-08 15:48:25.462","msg":"[Bot] 接收到nsq消息","msg":"msg-21","func":"JobTask"}
{"level":"info","time":"2023-06-08 15:48:27.450","msg":"[JobQueue] 处理数据","data":"msg-14"}
{"level":"info","time":"2023-06-08 15:48:27.450","msg":"[JobQueue] 处理数据","data":"msg-18"}
{"level":"info","time":"2023-06-08 15:48:27.467","msg":"[Bot] 接收到nsq消息","msg":"msg-22","func":"JobTask"}
{"level":"info","time":"2023-06-08 15:48:29.468","msg":"[Bot] 接收到nsq消息","msg":"msg-23","func":"JobQueue"}
{"level":"info","time":"2023-06-08 15:48:31.470","msg":"[Bot] 接收到nsq消息","msg":"msg-24","func":"JobTask"}
{"level":"info","time":"2023-06-08 15:48:33.472","msg":"[Bot] 接收到nsq消息","msg":"msg-25","func":"JobQueue"}
{"level":"info","time":"2023-06-08 15:48:35.474","msg":"[Bot] 接收到nsq消息","msg":"msg-26","func":"JobTask"}
{"level":"info","time":"2023-06-08 15:48:47.452","msg":"[JobQueue] 处理数据","data":"msg-23"}
{"level":"info","time":"2023-06-08 15:48:47.452","msg":"[JobQueue] 处理数据","data":"msg-25"}
{"level":"info","time":"2023-06-08 15:48:51.440","msg":"[JobTask] 处理数据","data":"msg-19"}
{"level":"info","time":"2023-06-08 15:48:51.440","msg":"[JobTask] 处理数据","data":"msg-20"}
{"level":"info","time":"2023-06-08 15:48:51.440","msg":"[JobTask] 处理数据","data":"msg-21"}
{"level":"info","time":"2023-06-08 15:48:51.440","msg":"[JobTask] 处理数据","data":"msg-22"}
{"level":"info","time":"2023-06-08 15:48:51.440","msg":"[JobTask] 处理数据","data":"msg-24"}
{"level":"info","time":"2023-06-08 15:48:51.440","msg":"[JobTask] 处理数据","data":"msg-26"}

```

----
