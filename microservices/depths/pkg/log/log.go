package log

import (
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
)

type LogLevel byte

const (
	PANIC LogLevel = iota + 1
	ERROR
	WARNING
	INFO
	DEBUG
)

var logger = New(DEBUG, os.Stderr)

type Logger struct {
	mu        sync.Mutex
	out       io.Writer
	buf       []byte
	isDiscard int32
	level     int32
}

func New(level LogLevel, out io.Writer) *Logger {
	log := &Logger{
		out:   out,
		level: int32(level),
	}
	if out == io.Discard {
		log.isDiscard = 1
	}
	return log
}

func (log *Logger) SetLevel(level LogLevel) {
	log.mu.Lock()
	defer log.mu.Unlock()
	atomic.StoreInt32(&log.level, int32(level))
}

func (log *Logger) SetOutput(w io.Writer) {
	log.mu.Lock()
	defer log.mu.Unlock()
	log.out = w
	isDiscard := int32(0)
	if w == io.Discard {
		isDiscard = 1
	}
	atomic.StoreInt32(&log.isDiscard, isDiscard)
}

func (log *Logger) Print(v ...interface{}) {
	logger.Output(DEBUG, fmt.Sprint(v...))
}
func (log *Logger) Printf(format string, v ...interface{}) {
	logger.Output(DEBUG, fmt.Sprintf(format, v...))
}
func (log *Logger) Println(v ...interface{}) {
	logger.Output(DEBUG, fmt.Sprint(v...))
}
