package log

import (
	"fmt"
	"time"
)

func Info(message ...interface{}) {
	fmt.Printf(" %s %c[1;40;32m%s%c[0m %s\n", nowTime(), 0x1B, "INFO", 0x1B, message)
}

func Waring(message ...interface{})  {
	fmt.Printf(" %s %c[1;40;33m%s%c[0m %s\n", nowTime(), 0x1B, "Waring", 0x1B, message)
}

func Debug(message ...interface{}) {
	fmt.Printf(" %s %c[1;40;34m%s%c[0m %s\n", nowTime(), 0x1B, "DEBUG", 0x1B, message)
}

func Error(message ...interface{})  {
	fmt.Printf(" %s %c[1;40;31m%s%c[0m %s\n", nowTime(), 0x1B, "ERROR", 0x1B, message)
}

func nowTime() string {
	timeStr:=time.Now().Format("2006-01-02 15:04:05")
	return timeStr
}
