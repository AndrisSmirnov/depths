package log

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func (log *Logger) Output(level LogLevel, s string) error {
	if atomic.LoadInt32(&log.isDiscard) != 0 {
		return nil
	}
	if atomic.LoadInt32(&log.level) < int32(level) {
		return nil
	}
	log.mu.Lock()
	defer log.mu.Unlock()
	log.buf = log.buf[:0]
	log.formatHeader(levelToString(level), &log.buf)
	log.buf = append(log.buf, s...)
	log.buf = append(log.buf, '\n')
	_, err := log.out.Write(log.buf)
	return err
}

func (log *Logger) formatHeader(level string, buf *[]byte) {
	t := time.Now()
	year, month, day := t.Date()
	itoa(buf, year, 4)
	*buf = append(*buf, '.')
	itoa(buf, int(month), 2)
	*buf = append(*buf, '.')
	itoa(buf, day, 2)
	*buf = append(*buf, ' ')
	hour, min, sec := t.Clock()
	itoa(buf, hour, 2)
	*buf = append(*buf, ':')
	itoa(buf, min, 2)
	*buf = append(*buf, ':')
	itoa(buf, sec, 2)
	*buf = append(*buf, ' ')

	levelString := fmt.Sprint("[", level, "] \t")
	*buf = append(*buf, levelString...)

	_, file, line, ok := runtime.Caller(3)
	if !ok {
		return
	}
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short
	*buf = append(*buf, file...)
	*buf = append(*buf, ':')
	itoa(buf, line, -1)
	*buf = append(*buf, ' ')
}

func levelToString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "debug"
	case INFO:
		return "info"
	case WARNING:
		return "warning"
	case ERROR:
		return "error"
	case PANIC:
		return "PANIC"
	default:
		return "..."
	}
}

func itoa(buf *[]byte, i int, wid int) {
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}
