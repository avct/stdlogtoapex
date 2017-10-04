package stdlogtoapex

import (
	"log"

	alog "github.com/apex/log"
)

type Writer struct {
	prefixLen int
}

func (w *Writer) stripDatePrefix(p []byte) string {
	return string(p[w.prefixLen:])
}

func (w *Writer) Write(p []byte) (n int, err error) {
	msg := w.stripDatePrefix(p)
	alog.Info(msg)
	return len(p), nil
}

func NewWriter(level alog.Level, handler alog.Handler) *Writer {
	var prefixLen int

	flags := log.Flags()
	if flags&log.Ldate != 0 {
		prefixLen += 11
	}
	if flags&log.Lmicroseconds != 0 {
		prefixLen += 16
	} else {
		if flags&log.Ltime != 0 {
			prefixLen += 9
		}
	}
	writer := &Writer{prefixLen: prefixLen}
	return writer
}
