package utils

import (
	"io"
	"log"
	"text/template"
	"zjhw.com/oneci/config"
)

func Render(temp string, target io.Writer, attr *config.AppInfo) {
	if t, err := template.New(attr.APP).Parse(temp); err == nil {
		t.Execute(target, attr)
	} else {
		log.Fatalf("**** 渲染模板失败, %s, %v", string(temp), err)
	}
}
