package rose

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
	"time"
)

func ExecShell(str string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", str)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

// ExecCmd execute shell command
// "/bin/sh", "-c", "command"
func ExecCmd(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	outBytes, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	// Clean the output and remove special characters
	outStr := strings.TrimSpace(string(outBytes))
	return outStr, nil
}

// ExecDurationTime 获取函数执行时间
// use: defer tools.ExecDurationTime()()
func ExecDurationTime() func() {
	t := time.Now()
	return func() {
		log.Printf("[Tips] The current function execution time is (%v) \n", time.Since(t))
	}
}
