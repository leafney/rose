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
	"log"
	"os"
)

type Log struct {
	debug  bool
	logger *log.Logger
}

func NewXLog(debug bool) *Log {
	return &Log{
		debug:  debug,
		logger: log.New(os.Stdout, "[Log] ", log.LstdFlags|log.Lmsgprefix),
	}
}

func (c *Log) SetDebug(enable bool) {
	c.debug = enable
}

func (c *Log) SetPrefix(prefix string) {
	c.logger.SetPrefix(prefix)

}

func (c *Log) SetFlags(flag int) {
	c.logger.SetFlags(flag)
}

func (c *Log) Println(v ...any) {
	if c.debug {
		c.logger.Println(v...)
	}
}

func (c *Log) Printf(format string, v ...any) {
	if c.debug {
		c.logger.Printf(format, v...)
	}
}

func (c *Log) Printfn(format string, v ...any) {
	if c.debug {
		msg := fmt.Sprintf(format, v...)
		c.logger.Printf("%s\n", msg)
	}
}
