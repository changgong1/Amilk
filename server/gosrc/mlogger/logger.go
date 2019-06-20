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

// 日志等级和描述映射关系
var LevelMap = map[string]int{
	"EMER": LevelEmergency,
	"ALRT": LevelAlert,
	"CRIT": LevelCritical,
	"EROR": LevelError,
	"WARN": LevelWarning,
	"INFO": LevelInformational,
	"DEBG": LevelDebug,
	"TRAC": LevelTrace,
}

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
var defaultLogger *mLogger
type mLogger struct {
	LevelStr string
	Level int
}
// SetMlogger set level
func SetMlogger(level levelType) {
	defaultLogger.LevelStr = level
	if l, ok := LevelMap[level]; ok {
		defaultLogger.Level = l
	}
}

func (m *mLogger) Eror(f string, v ...interface{}) {
	if m.Level < LevelError && m.LevelStr != "" {
		return 
	}
	log.Printf(f, v...)
}

func (m *mLogger) Debg(f string, v ...interface{}) {
	if m.Level < LevelDebug && m.LevelStr != "" {
		return 
	}
	log.Printf(f, v...)
}
// Eror define a Debg
func Error(f string,v ...interface{}){
	defaultLogger.Eror(f, v...)
}

// Debg define a Debg
func Debug(f string,v ...interface{}){
	defaultLogger.Debg(f, v...)
}

// init a mLogger 
func init(){
	defaultLogger = new(mLogger)
}

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