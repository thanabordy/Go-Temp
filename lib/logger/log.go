package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/spf13/viper"
)

// Priority level
type Priority int

// Log Level
const (
	LogEmerg Priority = iota
	LogAlert
	LogCrit
	LogErr
	LogWarning
	LogNotice
	LogInfo
	LogDebug
)

var (
	logLevel = map[string]Priority{
		"emerg":   LogEmerg,
		"alert":   LogAlert,
		"crit":    LogCrit,
		"err":     LogErr,
		"warning": LogWarning,
		"notice":  LogNotice,
		"info":    LogInfo,
		"debug":   LogDebug,
	}
	logName = map[Priority]string{
		LogEmerg:   "emerg",
		LogAlert:   "alert",
		LogCrit:    "crit",
		LogErr:     "err",
		LogWarning: "warning",
		LogNotice:  "notice",
		LogInfo:    "info",
		LogDebug:   "debug",
	}
	logVal = logstruct{}
)

type logstruct struct {
	mu     sync.Mutex
	logOut io.Writer
}

func init() {
	SetOutput(os.Stderr)
}

// SetOutput sets the output destination for the logger.
func SetOutput(w io.Writer) {
	logVal.mu.Lock()
	defer logVal.mu.Unlock()
	logVal.logOut = w
}

func getOutputLogLevel() Priority {
	levelStr := viper.GetString("APP_LOG_LEVEL")
	if level, ok := logLevel[levelStr]; ok {
		return level
	}
	return LogWarning
}

// Output Priority
func Output(level Priority, calldepth int, v ...interface{}) {
	Outputf(level, calldepth+1, "%s", fmt.Sprint(v...))
}

// Outputf Priority
func Outputf(level Priority, calldepth int, format string, v ...interface{}) {
	if level <= getOutputLogLevel() {
		now := time.Now()
		year, month, day := now.Date()
		hour, min, sec := now.Clock()
		zone, _ := now.Zone()
		logDate := fmt.Sprintf("%d/%02d/%02d %d:%d:%d %s", year, int(month), day, hour, min, sec, zone)
		pc, file, line, ok := runtime.Caller(calldepth)
		_, _, _ = file, line, ok
		funca := runtime.FuncForPC(pc).Name()
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				file = file[i+1:]
				break
			}
		}
		logSrc := fmt.Sprintf("%s:%d(%s)", file, line, funca)
		prefix := fmt.Sprintf("%s, %s", logDate, logSrc)
		str := fmt.Sprintf("[%s]: %s [%s] %s", prefix, viper.GetString("APP_NAME"), logName[level], fmt.Sprintf(format, v...))
		fmt.Fprintln(logVal.logOut, str)
	}
}

// Emerg LogEmerg by log.Print
func Emerg(v ...interface{}) {
	Output(LogEmerg, 2, v...)
}

// Emergf LogEmerg by log.Printf
func Emergf(format string, v ...interface{}) {
	Outputf(LogEmerg, 2, format, v...)
}

// Alert LogAlert by log.Print
func Alert(v ...interface{}) {
	Output(LogAlert, 2, v...)
}

// Alertf LogAlert by log.Printf
func Alertf(format string, v ...interface{}) {
	Outputf(LogAlert, 2, format, v...)
}

// Crit LogCrit by log.Print
func Crit(v ...interface{}) {
	Output(LogCrit, 2, v...)
}

// Critf LogCrit by log.Printf
func Critf(format string, v ...interface{}) {
	Outputf(LogCrit, 2, format, v...)
}

// Err LogErr by log.Print
func Err(v ...interface{}) {
	Output(LogErr, 2, v...)
}

// Errf LogErr by log.Printf
func Errf(format string, v ...interface{}) {
	Outputf(LogErr, 2, format, v...)
}

// Warning LogWarning by log.Print
func Warning(v ...interface{}) {
	Output(LogWarning, 2, v...)
}

// Warningf LogWarning by log.Printf
func Warningf(format string, v ...interface{}) {
	Outputf(LogWarning, 2, format, v...)
}

// Notice LogNotice by log.Print
func Notice(v ...interface{}) {
	Output(LogNotice, 2, v...)
}

// Noticef LogNotice by log.Printf
func Noticef(format string, v ...interface{}) {
	Outputf(LogNotice, 2, format, v...)
}

// Info LogInfo by log.Print
func Info(v ...interface{}) {
	Output(LogInfo, 2, v...)
}

// Infof LogInfo by log.Printf
func Infof(format string, v ...interface{}) {
	Outputf(LogInfo, 2, format, v...)
}

// Debug LogDebug by log.Print
func Debug(v ...interface{}) {
	Output(LogDebug, 2, v...)
}

// Debugf LogDebug by log.Printf
func Debugf(format string, v ...interface{}) {
	Outputf(LogDebug, 2, format, v...)
}
