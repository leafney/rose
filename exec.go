package rose

import (
	"bytes"
	"context"
	"log"
	"os/exec"
	"strings"
	"time"
)

const (
	Bash   = "/bin/bash"
	SplitC = "-c"
)

func ExecCmdRun(name string, arg ...string) error {
	return ExecCmdRunCtx(context.Background(), name, arg...)
}

func ExecCmdRunCtx(ctx context.Context, name string, arg ...string) error {
	cmd := exec.CommandContext(ctx, name, arg...)
	return cmd.Run()
}

func ExecCmdOut(name string, arg ...string) (string, error) {
	return ExecCmdOutCtx(context.Background(), name, arg...)
}

func ExecCmdOutCtx(ctx context.Context, name string, arg ...string) (string, error) {
	cmd := exec.CommandContext(ctx, name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return trimOut(out.Bytes()), err
}

func ExecCmdOutErr(name string, arg ...string) (string, string, error) {
	return ExecCmdOutErrCtx(context.Background(), name, arg...)
}

func ExecCmdOutErrCtx(ctx context.Context, name string, arg ...string) (string, string, error) {
	cmd := exec.CommandContext(ctx, name, arg...)
	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut // 标准输出
	cmd.Stderr = &stdErr // 标准错误
	err := cmd.Run()
	return trimOut(stdOut.Bytes()), trimOut(stdErr.Bytes()), err
}

// ExecCmdBashOut
// "/bin/bash", "-c", "command"
func ExecCmdBashOut(command string) (string, error) {
	return ExecCmdBashOutCtx(context.Background(), command)
}

// ExecCmdBashOutCtx
// "/bin/bash", "-c", "command"
func ExecCmdBashOutCtx(ctx context.Context, command string) (string, error) {
	cmd := exec.CommandContext(ctx, Bash, SplitC, command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return trimOut(out.Bytes()), err
}

// ExecCmdBashOutErr
// "/bin/bash", "-c", "command"
func ExecCmdBashOutErr(command string) (string, string, error) {
	return ExecCmdBashOutErrCtx(context.Background(), command)
}

// ExecCmdBashOutErrCtx
// "/bin/bash", "-c", "command"
func ExecCmdBashOutErrCtx(ctx context.Context, command string) (string, string, error) {
	cmd := exec.CommandContext(ctx, Bash, SplitC, command)
	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut // 标准输出
	cmd.Stderr = &stdErr // 标准错误
	err := cmd.Run()
	return trimOut(stdOut.Bytes()), trimOut(stdErr.Bytes()), err
}

func ExecCmdCombinedOut(name string, arg ...string) (string, error) {
	return ExecCmdCombinedOutCtx(context.Background(), name, arg...)
}

// ExecCmdCombinedOutCtx CombinedOutput
func ExecCmdCombinedOutCtx(ctx context.Context, name string, arg ...string) (string, error) {
	cmd := exec.CommandContext(ctx, name, arg...)
	outBytes, err := cmd.CombinedOutput()

	outStr := trimOut(outBytes)
	return outStr, err
}

// trimOut Clean the output and remove special characters
func trimOut(input []byte) string {
	return strings.TrimSpace(string(input))
}

// ----------------

// ExecDurationTime 获取函数执行时间
// use: defer tools.ExecDurationTime()()
func ExecDurationTime() func() {
	t := time.Now()
	return func() {
		log.Printf("[Tips] The current function execution time is (%v) \n", time.Since(t))
	}
}