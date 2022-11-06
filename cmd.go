package rose

import (
	"bytes"
	"log"
	"os/exec"
	"time"
)

func ExecShell(s string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

// ExecDurationTime 获取函数执行时间
// use: defer tools.ExecDurationTime()()
func ExecDurationTime() func() {
	t := time.Now()
	return func() {
		log.Printf("[Tips] The current function execution time is (%v) \n", time.Since(t))
	}
}
