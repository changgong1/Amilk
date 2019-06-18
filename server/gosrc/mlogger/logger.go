package mlogger

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

// 日志等级，从0-7，日优先级由高到低
const (
	LevelEmergency     = iota // 系统级紧急，比如磁盘出错，内存异常，网络不可用等
	LevelAlert                // 系统级警告，比如数据库访问异常，配置文件出错等
	LevelCritical             // 系统级危险，比如权限出错，访问异常等
	LevelError                // 用户级错误
	LevelWarning              // 用户级警告
	LevelInformational        // 用户级信息
	LevelDebug                // 用户级调试
	LevelTrace                // 用户级基本输出
)

// 日志记录等级字段
var levelPrefix = [LevelTrace + 1]string{
	"EMER",
	"ALRT",
	"CRIT",
	"EROR",
	"WARN",
	"INFO",
	"DEBG",
	"TRAC",
}
type levelType = string

// LogerDebgPrint is Print msg
func LogerDebgPrint(level levelType, formating string, args ...interface{}) {
	if level != levelPrefix[LevelDebug] {
		return
	}
	LogerPrint("[DEBUG]" + formating, args...)
}

// LogerPrint is Print msg
func LogerPrint(formating string, args ...interface{}) {
	filename, line, funcname := "???", 0, "???"
	pc, filename, line, ok := runtime.Caller(2)
	// fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))
	if ok {
		funcname = runtime.FuncForPC(pc).Name()      // main.(*MyStruct).foo
		funcname = filepath.Ext(funcname)            // .foo
		funcname = strings.TrimPrefix(funcname, ".") // foo

		filename = filepath.Base(filename) // /full/path/basename.go => basename.go
	}

	log.Printf("%s:%d:%s: %s\n", filename, line, funcname, fmt.Sprintf(formating, args...))
}