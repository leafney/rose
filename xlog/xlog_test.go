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

	loggg := NewXLog(false).SetDebug(true).SetPrefix("OKOK")

	fmt.Println("↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓")
	loggg.Debug("aaa", "bbb", "ccc")
	loggg.Debugln("aaa", "bbb", "ccc")
	loggg.Info("hello1")
	loggg.Infof("hello2 %v", 12345)
	loggg.Infoln("world")
	fmt.Println("↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑")
	log.Print("aaa", "bbb", "ccc")
	log.Println("aaa", "bbb", "ccc")
	log.Print("hello1")
	log.Printf("hello2 %v", 12345)
	log.Println("world")
	fmt.Println("↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑")
}
