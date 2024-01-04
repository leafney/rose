/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2023-12-19 16:21
 * @Description:
 */

package xlog

import (
	"fmt"
	"log"
	"testing"
)

func TestNewXLog(t *testing.T) {

	xlog := NewXLog(false).
		//SetDebug(true).
		//SetPrefix("").
		//SetPrefix("[hello]").
		//SetLevel(ErrorLevel).
		SetEnable(true)

	fmt.Println("↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓")
	xlog.Debug("aaa", "bbb", "ccc")
	xlog.Debugln("aaa", "bbb", "ccc")
	xlog.Info("hello1")
	xlog.Infof("hello2 %v", 12345)
	xlog.Infoln("world", "oooo")
	xlog.Errorln("haha", "ohoh")
	fmt.Println("↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑")
	log.Print("aaa", "bbb", "ccc")
	log.Println("aaa", "bbb", "ccc")
	log.Print("hello1")
	log.Printf("hello2 %v", 12345)
	log.Println("world", "oooo")
	fmt.Println("↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑")
}
