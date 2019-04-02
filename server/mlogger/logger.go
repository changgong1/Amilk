package mlogger

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

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
