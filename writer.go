/* stdlogtoapex is a package that provides an implementation of
/* io.Writer that can be used to redirect log output from the standard
/* libraries log package via github.com/apex/log.
*/
package stdlogtoapex

import (
	"log"

	alog "github.com/apex/log"
)

// Writer is a struct that implements the io.Writer interface and can,
// therefore, be passed into log.SetOutput.  Writer must always be
// constructed by called NewWriter.
type Writer struct {
	prefixLen int
}

func (w *Writer) stripDatePrefix(p []byte) string {
	if w.prefixLen > len(p) {
		// Safer to return the old value than trim all.
		return string(p)
	}
	return string(p[w.prefixLen:])
}

// Write implements the io.Writer interface for Writer.  Log messages
// output via this method will have their date and time information
// stripped (apex log will have its own) .
func (w *Writer) Write(p []byte) (n int, err error) {
	msg := w.stripDatePrefix(p)
	alog.Info(msg)
	return len(p), nil
}

// NewWriter creates a new Writer that can be passed to the SetOutput
// function in the log package from the standard library.
func NewWriter() *Writer {
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
