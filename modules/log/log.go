// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package log is a wrapper of logs for short calling name.
package log

import (
	"fmt"

	"github.com/martini-contrib/render"

	"github.com/gogits/logs"

	"github.com/gogits/gogs/modules/base"
)

var logger *logs.BeeLogger

func init() {
	logger = logs.NewLogger(10000)
	logger.SetLogger("console", "")
}

func Trace(format string, v ...interface{}) {
	logger.Trace(format, v...)
}

func Info(format string, v ...interface{}) {
	logger.Info(format, v...)
}

func Error(format string, v ...interface{}) {
	logger.Error(format, v...)
}

func Warn(format string, v ...interface{}) {
	logger.Warn(format, v...)
}

func Critical(format string, v ...interface{}) {
	logger.Critical(format, v...)
}

func Handle(status int, title string, data base.TmplData, r render.Render, err error) {
	data["ErrorMsg"] = err
	Error("%s: %v", title, err)
	r.HTML(status, fmt.Sprintf("status/%d", status), data)
}