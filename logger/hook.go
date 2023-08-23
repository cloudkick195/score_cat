package logger

import (
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

type CustomHook struct{}

func (hook CustomHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook CustomHook) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, 15)
	n := runtime.Callers(10, pc)
	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, more := frames.Next()
		if !more {
			break
		}
		if !isLogrusFile(frame.File) {
			entry.Data["file"] = frame.File
			entry.Data["line"] = frame.Line
			break
		}
	}
	return nil
}

func isLogrusFile(file string) bool {
	return filepath.Base(file) == "logrus_entry.go" ||
		filepath.Base(file) == "logger.go"
}
