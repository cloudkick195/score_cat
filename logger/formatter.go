package logger

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

// Formatter cho logger
var Formatter = &logrus.TextFormatter{
	FullTimestamp:    true,
	CallerPrettyfier: callerPrettyfier,
	DisableColors:    true,
}

// In ra thông tin về file lỗi và dòng lỗi khi có lỗi xảy ra
func callerPrettyfier(f *runtime.Frame) (string, string) {
	_, filename := filepath.Split(f.File)
	return fmt.Sprintf("%s:%d", filename, f.Line), ""
}
