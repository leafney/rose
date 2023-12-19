/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2023-07-14 16:36
 * @Description:
 */

package xlog

import (
	"fmt"
	"github.com/leafney/rose"
	"log"
	"os"
)

type LogLevel string

const (
	DebugLevel LogLevel = "DEBUG"
	InfoLevel  LogLevel = "INFO"
	ErrorLevel LogLevel = "ERROR"
	FatalLevel LogLevel = "FATAL"
)

type Log struct {
	debug  bool
	enable bool
	logger *log.Logger
}

func NewXLog(debug bool) *Log {
	return &Log{
		debug:  debug,
		enable: true,
		logger: log.New(os.Stdout, "[XLog] ", log.LstdFlags|log.Lmsgprefix),
	}
}

func (c *Log) SetDebug(debug bool) *Log {
	c.debug = debug
	return c
}

func (c *Log) SetEnable(enable bool) *Log {
	c.debug = enable
	return c
}

func (c *Log) SetPrefix(prefix string) *Log {
	if !rose.StrIsEmpty(prefix) {
		if prefix[len(prefix)-1:] != " " {
			c.logger.SetPrefix(prefix + " ")
		} else {
			c.logger.SetPrefix(prefix)
		}
	}
	return c
}

func (c *Log) SetFlags(flag int) *Log {
	c.logger.SetFlags(flag)
	return c
}

func (c *Log) logf(level LogLevel, format string, v ...any) {
	if c.enable && (c.debug || level != DebugLevel) {
		msg := fmt.Sprintf(format, v...)
		c.logger.Printf("[%s]: %s", level, msg)
	}
	if level == FatalLevel {
		os.Exit(1)
	}
}

func (c *Log) Debug(v ...any) {
	if c.debug {
		c.logf(DebugLevel, "%s", fmt.Sprint(v...))
	}
}

func (c *Log) Debugf(format string, v ...any) {
	if c.debug {
		c.logf(DebugLevel, format, v...)
	}
}

func (c *Log) Debugln(v ...any) {
	if c.debug {
		c.logf(DebugLevel, "%s", fmt.Sprintln(v...))
	}
}

func (c *Log) Info(v ...any) {
	c.logf(InfoLevel, "%s", fmt.Sprint(v...))
}

func (c *Log) Infof(format string, v ...any) {
	c.logf(InfoLevel, format, v...)
}

func (c *Log) Infoln(v ...any) {
	c.logf(InfoLevel, "%s", fmt.Sprintln(v...))
}

func (c *Log) Error(v ...any) {
	c.logf(ErrorLevel, "%s", fmt.Sprint(v...))
}

func (c *Log) Errorf(format string, v ...any) {
	c.logf(ErrorLevel, format, v...)
}

func (c *Log) Errorln(v ...any) {
	c.logf(ErrorLevel, "%s", fmt.Sprintln(v...))
}

func (c *Log) Fatal(v ...any) {
	c.logf(FatalLevel, "%s", fmt.Sprint(v...))
}

func (c *Log) Fatalf(format string, v ...any) {
	c.logf(FatalLevel, format, v...)
}

func (c *Log) Fatalln(v ...any) {
	c.logf(FatalLevel, "%s", fmt.Sprintln(v...))
}

//
//func (c *Log) Println(v ...any) {
//	if c.debug {
//		c.logger.Println(v...)
//	}
//}
//
//func (c *Log) Printf(format string, v ...any) {
//	if c.debug {
//		c.logger.Printf(format, v...)
//	}
//}
//
//func (c *Log) Printfln(format string, v ...any) {
//	if c.debug {
//		msg := fmt.Sprintf(format, v...)
//		c.logger.Printf("%s\n", msg)
//	}
//}
