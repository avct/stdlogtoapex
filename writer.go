package stdlogtoapex

import alog "github.com/apex/log"

type Writer struct {
}

func (w *Writer) Write(p []byte) (n int, err error) {
	msg := string(p)
	alog.Info(msg)
	return len(p), nil
}

func NewWriter(handler alog.Handler) *Writer {
	writer := &Writer{}
	return writer
}
