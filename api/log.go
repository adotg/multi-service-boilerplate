package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/zput/zxcTool/ztLog/zt_formatter"
)

func GetLogFormatter() *zt_formatter.ZtFormatter {
	var formatter = &zt_formatter.ZtFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}
	return formatter
}
