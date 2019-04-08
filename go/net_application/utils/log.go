//处理日志文件

package utils

import (	
	"log"
	"time"
	"os"
	"strings"
	"fmt"
	"runtime"
)

var loghandler *log.Logger
var MsgTypes map[string]int = make(map[string]int, 4)

func init() {
	//日志记录类型
	MsgTypes["INFO"] = 1;
	MsgTypes["WARNING"] = 1;
	MsgTypes["DEBUG"] = 1;
	MsgTypes["ERROR"] = 1;	

	currentYear := time.Now().Year()
	currentMonth := time.Now().Month()
	currentDay := time.Now().Day()
	currentDate := fmt.Sprintf("%0.2d_%0.2d_%0.2d", currentYear, currentMonth, currentDay)	


	filename := "../log/" + currentDate  + ".log"
	logFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	runtime.SetFinalizer(logFile, (*os.File).Close)

	
	loghandler = log.New(logFile, "INFO ", log.Ldate|log.Ltime|log.Llongfile)
}


//输出日志
func LogMessage(msgType string, args ...interface{}){
	msgType = strings.ToUpper(msgType)
	prefix := "WARNING"
	_, ok := MsgTypes[msgType]
	if ok == true {
		prefix = msgType
	}

	loghandler.SetPrefix(prefix + " ")
	loghandler.Println(args...)
}

//错误输出
func OutputError(err error) {
	if err != nil {
		log.Fatalln("Fatal error:", err)
	}
}
