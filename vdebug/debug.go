package vdebug

import (
	"io"
	"log"
	"os"
	"runtime/debug"
	"time"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/sunshinexcode/gotk/vstr"
)

func D(data ...interface{}) {
	Debug(data...)
}

func Debug(data ...interface{}) {
	log.New(io.MultiWriter([]io.Writer{os.Stdout}...), "", 0).Print(vstr.S("%c[0;39;32mDEBUG - %s - data:%v%c[0m", 0x1B, time.Now().Format("2006-01-02 15:04:05.000"), data, 0x1B))
}

// Dump dumps a variable to stdout with more manually readable.
func Dump(values ...any) {
	g.Dump(values)
}

// Stack stack information
func Stack() string {
	return string(debug.Stack())
}
