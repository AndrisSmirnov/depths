package log

import "fmt"

func Default() *Logger { return logger }

func SetLevel(level LogLevel) {
	logger.SetLevel(level)
}

func Info(v ...any) {
	logger.Output(INFO, fmt.Sprint(v...))
}
func Infof(format string, v ...any) {
	logger.Output(INFO, fmt.Sprintf(format, v...))
}

func Debug(v ...any) {
	logger.Output(DEBUG, fmt.Sprint(v...))
}
func Debugf(format string, v ...any) {
	logger.Output(DEBUG, fmt.Sprintf(format, v...))
}

func Warning(v ...any) {
	logger.Output(WARNING, fmt.Sprint(v...))
}
func Warningf(format string, v ...any) {
	logger.Output(WARNING, fmt.Sprintf(format, v...))
}

func Error(v ...any) {
	logger.Output(ERROR, fmt.Sprint(v...))
}
func Errorf(format string, v ...any) {
	logger.Output(ERROR, fmt.Sprintf(format, v...))
}

func Panic(v ...any) {
	s := fmt.Sprint(v...)
	logger.Output(PANIC, fmt.Sprint(s))
	panic(s)
}
func Panicf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	logger.Output(PANIC, fmt.Sprint(s))
	panic(s)
}
