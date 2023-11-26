/*
 * Copyright 2023 Attains Cloud, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 *
 * visit: https://cloud.attains.cn
 *
 */

package logger

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
)

// FileWithLineNum return the file name and line number of the current file
func FileWithLineNum() string {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		return file + ":" + strconv.FormatInt(int64(line), 10)
	}

	return ""
}

// Colors
const (
	Reset    = "\033[0m"
	Red      = "\033[31m"
	Green    = "\033[32m"
	Magenta  = "\033[35m"
	Cyan     = "\033[33m"
	BlueBold = "\033[34;1m"
)

// LogLevel log level
type LogLevel int

const (
	// Silent silent log level
	Silent LogLevel = iota + 1

	// Warn warn log level
	Warn
	// Error error log level
	Error
	// Info info log level
	Info
	// Debug debug log level
	Debug
)

// Writer log writer interface
type Writer interface {
	Printf(string, ...interface{})
}

// Config logger config
type Config struct {
	Colorful bool
	LogLevel LogLevel
}

// Interface logger interface
type Interface interface {
	LogMode(LogLevel) Interface
	Info(context.Context, string, ...interface{})
	Warn(context.Context, string, ...interface{})
	Error(context.Context, string, ...interface{})
	Debug(context.Context, string, ...interface{})
}

var (
	// Discard logger will print any log to io.Discard
	Discard = New(log.New(ioutil.Discard, "", log.LstdFlags), Config{})
	// Default Default logger
	Default = New(log.New(os.Stdout, "\r\n", log.LstdFlags), Config{
		LogLevel: Debug,
		Colorful: true,
	})
)

// New initialize logger
func New(writer Writer, config Config) Interface {
	var (
		infoStr  = "%s\n[info] "
		warnStr  = "%s\n[warn] "
		errStr   = "%s\n[error] "
		debugStr = "%s\n[debug] "
	)

	if config.Colorful {
		infoStr = Green + "%s\n" + Reset + Green + "[info] " + Reset
		warnStr = BlueBold + "%s\n" + Reset + Magenta + "[warn] " + Reset
		errStr = Magenta + "%s\n" + Reset + Red + "[error] " + Reset
		debugStr = Cyan + "%s\n" + Reset + Cyan + "[debug] " + Reset
	}

	return &logger{
		Writer:   writer,
		Config:   config,
		infoStr:  infoStr,
		warnStr:  warnStr,
		errStr:   errStr,
		debugStr: debugStr,
	}
}

type logger struct {
	Writer
	Config
	infoStr, warnStr, errStr, debugStr string
}

// LogMode log mode
func (l *logger) LogMode(level LogLevel) Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Debug print debug
func (l *logger) Debug(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= Debug {
		l.Printf(l.debugStr+msg, append([]interface{}{FileWithLineNum()}, data...)...)
	}
}

// Info print info
func (l *logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= Info {
		l.Printf(l.infoStr+msg, append([]interface{}{FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (l *logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= Warn {
		l.Printf(l.warnStr+msg, append([]interface{}{FileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (l *logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= Error {
		l.Printf(l.errStr+msg, append([]interface{}{FileWithLineNum()}, data...)...)
	}
}
