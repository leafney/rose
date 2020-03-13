package snowflake

import (
	"fmt"
	"os"
	"strconv"
)

var worker *IdWorker

/// Params: Given the workerId, 0 < workerId < 1024
/// Param 0 will get workerId from ENV with `WORKER_ID`
func InitWorker(workerId int64) (err error) {
	if workerId == 0 {
		envs := os.Getenv("WORKER_ID")
		if len(envs) > 0 {
			workerId, _ = strconv.ParseInt(envs, 10, 64)
		}
	}

	fmt.Printf("SnowFlake use workerId: [%d]\n", workerId)
	worker, err = NewIdWorker(workerId)
	return err
}

// 获取NextId
func GetNextId() (ts int64, err error) {
	return worker.NextId()
}
