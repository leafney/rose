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

func ExecCmdRunDir(dir string, name string, arg ...string) error {
	return ExecCmdRunDirCtx(context.Background(), dir, name, arg...)
}

func ExecCmdRunCtx(ctx context.Context, name string, arg ...string) error {
	return ExecCmdRunDirCtx(ctx, "", name, arg...)
}

func ExecCmdRunDirCtx(ctx context.Context, dir string, name string, arg ...string) error {
	cmd := exec.CommandContext(ctx, name, arg...)
	if !StrIsEmpty(dir) {
		cmd.Dir = dir
	}
	return cmd.Run()
}

// -------------

func ExecCmdOut(name string, arg ...string) (out string, err error) {
	return ExecCmdOutCtx(context.Background(), name, arg...)
}

func ExecCmdOutDir(dir string, name string, arg ...string) (out string, err error) {
	return ExecCmdOutDirCtx(context.Background(), dir, name, arg...)
}

func ExecCmdOutCtx(ctx context.Context, name string, arg ...string) (out string, err error) {
	return ExecCmdOutDirCtx(ctx, "", name, arg...)
}

func ExecCmdOutDirCtx(ctx context.Context, dir string, name string, arg ...string) (out string, err error) {
	cmd := exec.CommandContext(ctx, name, arg...)
	var stdOut bytes.Buffer
	cmd.Stdout = &stdOut
	if !StrIsEmpty(dir) {
		cmd.Dir = dir
	}
	err = cmd.Run()
	out = trimOut(stdOut.Bytes())
	return
}

// -------------

func ExecCmdOutB(name string, arg ...string) (out []byte, err error) {
	return ExecCmdOutBCtx(context.Background(), name, arg...)
}

func ExecCmdOutDirB(dir string, name string, arg ...string) (out []byte, err error) {
	return ExecCmdOutDirBCtx(context.Background(), dir, name, arg...)
}

func ExecCmdOutBCtx(ctx context.Context, name string, arg ...string) (out []byte, err error) {
	return ExecCmdOutDirBCtx(ctx, "", name, arg...)
}

func ExecCmdOutDirBCtx(ctx context.Context, dir string, name string, arg ...string) (out []byte, err error) {
	cmd := exec.CommandContext(ctx, name, arg...)
	var stdOut bytes.Buffer
	cmd.Stdout = &stdOut
	if !StrIsEmpty(dir) {
		cmd.Dir = dir
	}
	err = cmd.Run()
	out = trimOutB(stdOut.Bytes())
	return
}

// -------------

func ExecCmdOutErr(name string, arg ...string) (out string, fail string, err error) {
	return ExecCmdOutErrCtx(context.Background(), name, arg...)
}

func ExecCmdOutErrDir(dir string, name string, arg ...string) (out string, fail string, err error) {
	return ExecCmdOutErrDirCtx(context.Background(), dir, name, arg...)
}

func ExecCmdOutErrCtx(ctx context.Context, name string, arg ...string) (out string, fail string, err error) {
	return ExecCmdOutErrDirCtx(ctx, "", name, arg...)
}

func ExecCmdOutErrDirCtx(ctx context.Context, dir string, name string, arg ...string) (out string, fail string, err error) {
	cmd := exec.CommandContext(ctx, name, arg...)
	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut // 标准输出
	cmd.Stderr = &stdErr // 标准错误
	if !StrIsEmpty(dir) {
		cmd.Dir = dir
	}
	err = cmd.Run()
	out = trimOut(stdOut.Bytes())
	fail = trimOut(stdErr.Bytes())
	return
}

// -------------

func ExecCmdOutErrB(name string, arg ...string) (out []byte, fail []byte, err error) {
	return ExecCmdOutErrBCtx(context.Background(), name, arg...)
}

func ExecCmdOutErrDirB(dir string, name string, arg ...string) (out []byte, fail []byte, err error) {
	return ExecCmdOutErrDirBCtx(context.Background(), dir, name, arg...)
}

func ExecCmdOutErrBCtx(ctx context.Context, name string, arg ...string) (out []byte, fail []byte, err error) {
	return ExecCmdOutErrDirBCtx(ctx, "", name, arg...)
}

func ExecCmdOutErrDirBCtx(ctx context.Context, dir string, name string, arg ...string) (out []byte, fail []byte, err error) {
	cmd := exec.CommandContext(ctx, name, arg...)
	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut // 标准输出
	cmd.Stderr = &stdErr // 标准错误
	if !StrIsEmpty(dir) {
		cmd.Dir = dir
	}
	err = cmd.Run()
	out = trimOutB(stdOut.Bytes())
	fail = trimOutB(stdErr.Bytes())
	return
}

// -------------

// ExecCmdBashOut
// "/bin/bash", "-c", "command"
func ExecCmdBashOut(command string) (out string, err error) {
	return ExecCmdBashOutCtx(context.Background(), command)
}

// ExecCmdBashOutCtx
// "/bin/bash", "-c", "command"
func ExecCmdBashOutCtx(ctx context.Context, command string) (out string, err error) {
	cmd := exec.CommandContext(ctx, Bash, SplitC, command)
	var stdOut bytes.Buffer
	cmd.Stdout = &stdOut
	err = cmd.Run()
	out = trimOut(stdOut.Bytes())
	return
}

// -------------

func ExecCmdBashOutB(command string) (out []byte, err error) {
	return ExecCmdBashOutBCtx(context.Background(), command)
}

func ExecCmdBashOutBCtx(ctx context.Context, command string) (out []byte, err error) {
	cmd := exec.CommandContext(ctx, Bash, SplitC, command)
	var stdOut bytes.Buffer
	cmd.Stdout = &stdOut
	err = cmd.Run()
	out = trimOutB(stdOut.Bytes())
	return
}

// -------------

// ExecCmdBashOutErr
// "/bin/bash", "-c", "command"
func ExecCmdBashOutErr(command string) (out string, fail string, err error) {
	return ExecCmdBashOutErrCtx(context.Background(), command)
}

// ExecCmdBashOutErrCtx
// "/bin/bash", "-c", "command"
func ExecCmdBashOutErrCtx(ctx context.Context, command string) (out string, fail string, err error) {
	cmd := exec.CommandContext(ctx, Bash, SplitC, command)
	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut // 标准输出
	cmd.Stderr = &stdErr // 标准错误
	err = cmd.Run()
	out = trimOut(stdOut.Bytes())
	fail = trimOut(stdErr.Bytes())
	return
}

// -------------

func ExecCmdBashOutErrB(command string) (out []byte, fail []byte, err error) {
	return ExecCmdBashOutErrBCtx(context.Background(), command)
}

func ExecCmdBashOutErrBCtx(ctx context.Context, command string) (out []byte, fail []byte, err error) {
	cmd := exec.CommandContext(ctx, Bash, SplitC, command)
	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut // 标准输出
	cmd.Stderr = &stdErr // 标准错误
	err = cmd.Run()
	out = trimOutB(stdOut.Bytes())
	fail = trimOutB(stdErr.Bytes())
	return
}

// -------------

// ExecCmdCombinedOut CombinedOutput
func ExecCmdCombinedOut(name string, arg ...string) (out string, err error) {
	return ExecCmdCombinedOutCtx(context.Background(), name, arg...)
}

// ExecCmdCombinedOutDir CombinedOutput
func ExecCmdCombinedOutDir(dir string, name string, arg ...string) (out string, err error) {
	return ExecCmdCombinedOutDirCtx(context.Background(), dir, name, arg...)
}

// ExecCmdCombinedOutCtx CombinedOutput
func ExecCmdCombinedOutCtx(ctx context.Context, name string, arg ...string) (out string, err error) {
	return ExecCmdCombinedOutDirCtx(ctx, "", name, arg...)
}

// ExecCmdCombinedOutDirCtx CombinedOutput
func ExecCmdCombinedOutDirCtx(ctx context.Context, dir string, name string, arg ...string) (out string, err error) {
	cmd := exec.CommandContext(ctx, name, arg...)
	if !StrIsEmpty(dir) {
		cmd.Dir = dir
	}
	outBytes, err := cmd.CombinedOutput()
	out = trimOut(outBytes)
	return
}

// -------------

// trimOut Clean the output and remove special characters
func trimOut(input []byte) string {
	return strings.TrimSpace(string(input))
}

func trimOutB(input []byte) []byte {
	return bytes.TrimSpace(input)
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
