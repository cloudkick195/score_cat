package logger

import (
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger cho file app.log
var Log = logrus.New()

func Init() {
	// Thiết lập output cho logger
	Log.SetOutput(createRollingLogger("logs/app", "app.log"))

	// Thiết lập formatter cho logger
	Log.SetFormatter(Formatter)
	Log.AddHook(CustomHook{})
}

// Tạo rolling logger
func createRollingLogger(logFolder, logFile string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filepath.Join(logFolder, logFile),
		MaxSize:    5, // Megabytes
		MaxBackups: 3,
		MaxAge:     7, // Days
		LocalTime:  true,
	}
}

func Error(args ...interface{}) {
	// Lấy thông tin về caller
	_, file, line, ok := runtime.Caller(1)
	if ok {
		// Tạo một đối tượng Entry mới với thông tin về caller
		entry := Log.WithFields(logrus.Fields{
			"file": file,
			"line": line,
		})
		entry.Error(args...)
	} else {
		Log.Error(args...)
	}
}
