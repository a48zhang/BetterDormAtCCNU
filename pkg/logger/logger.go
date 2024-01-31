package logger

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
)

func init() {
	f, err := os.Create("gin.log")
	if err != nil {
		FatalLog(err.Error())
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	InfoLog("Ready.")
}

type DebugInfo struct {
	FuncName string
	FileName string
	Line     int
}

func GetCaller(count ...int) DebugInfo {
	if count == nil {
		count = []int{1}
	}
	pc, fl, ln, _ := runtime.Caller(count[0])
	runtime.FuncForPC(pc).Name()
	return DebugInfo{
		FuncName: runtime.FuncForPC(pc).Name(),
		FileName: fl,
		Line:     ln,
	}
}

/*
1: "Info",
2: "Debug",
3: "Warning",
4: "Error",
5: "Fatal",
*/
func writeLog(s string, tp int, pc ...int) {
	tpm := map[int]string{
		0: "Log",
		1: "Info",
		2: "Debug",
		3: "Warning",
		4: "Error",
		5: "Fatal",
		6: "Panic",
		7: "Undefined",
		8: "Reversed",
	}
	if tpm[tp] == "" {
		tp = 7
	}
	typ := "[" + tpm[tp] + "] "
	if pc == nil {
		pc = []int{2}
	}

	v := GetCaller(pc[0])
	//from := v.FileName + " " + v.FuncName + " " + strconv.Itoa(v.Line)
	from := v.FuncName + " " + strconv.Itoa(v.Line)
	write, err := gin.DefaultWriter.Write([]byte(typ + from + s + "\n"))
	if err != nil {
		ErrorLog(err)
		ErrorLog("Write " + strconv.Itoa(write) + " bytes.")
		return
	}
}

func InfoLog(s ...string) {
	msg := " "
	for _, v := range s {
		msg += v + " "
	}
	writeLog(msg, 1, 3)
}

func DebugLog(s ...string) {
	msg := " "
	for _, v := range s {
		msg += v + " "
	}
	writeLog(msg, 2, 3)
}

type T interface {
}

func ErrorLog(s T) {
	if _, ok := s.(string); ok {
		writeLog(s.(string), 4, 3)
	} else if _, ok := s.(error); ok {
		writeLog(s.(error).Error(), 4, 3)
	}
}

func FatalLog(s ...string) {
	msg := " "
	for _, v := range s {
		msg += v + " "
	}
	writeLog(msg, 5, 3)
	log.Fatalln("[Fatal]" + msg)
}
