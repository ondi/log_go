//
//
//

package log

import (
	"io"
	"path/filepath"
	"runtime"
	"strconv"
)

func FileLine(skip int, limit int) (path string, line int) {
	var next_line int
	var next_path string
	_, path, line, ok := runtime.Caller(skip)
	for i := skip + 1; i < limit; i++ {
		if _, next_path, next_line, ok = runtime.Caller(i); !ok {
			return
		}
		if filepath.Dir(path) != filepath.Dir(next_path) {
			return next_path, next_line
		}
	}
	return
}

type DT_t struct {
	Layout string
}

func NewDt(layout string) Formatter {
	return &DT_t{Layout: layout}
}

func (self *DT_t) FormatLog(out io.Writer, m LogMsg_t) (n int, err error) {
	var b [64]byte
	if n, err = out.Write(m.Level.Ts.AppendFormat(b[:0], self.Layout)); n > 0 {
		io.WriteString(out, " ")
	}
	return
}

type FileLine_t struct{}

func NewFileLine() Formatter {
	return &FileLine_t{}
}

func (self *FileLine_t) FormatLog(out io.Writer, m LogMsg_t) (n int, err error) {
	if n, err = io.WriteString(out, filepath.Base(m.Level.File)); n > 0 {
		io.WriteString(out, ":")
		io.WriteString(out, strconv.FormatInt(int64(m.Level.Line), 10))
		io.WriteString(out, " ")
	}
	return
}

type GetLogContext_t struct{}

func NewGetLogContext() Formatter {
	return &GetLogContext_t{}
}

func (self *GetLogContext_t) FormatLog(out io.Writer, m LogMsg_t) (n int, err error) {
	if v := GetLogContext(m.Ctx); v != nil {
		if n, err = io.WriteString(out, v.ContextName()); n > 0 {
			io.WriteString(out, " ")
		}
	}
	return
}
