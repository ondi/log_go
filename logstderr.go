//
//
//

package log

import "os"
import "fmt"
import "time"

type Stderr_t struct {
	datetime func() string
}

func NewStderr(datetime string) Writer {
	self := &Stderr_t{}
	if len(datetime) > 0 {
		datetime += " "
		self.datetime = func() string {return time.Now().Format(datetime)}
	} else {
		self.datetime = func() string {return ""}
	}
	return self
}

func (self * Stderr_t) Write(level string, format string, args ...interface{}) (err error) {
	_, err = fmt.Fprintf(os.Stderr, self.datetime() + level + " " + format + "\n", args...)
	return
}

type Stdout_t struct {
	datetime func() string
}

func NewStdout(datetime string) Writer {
	self := &Stdout_t{}
	if len(datetime) > 0 {
		datetime += " "
		self.datetime = func() string {return time.Now().Format(datetime)}
	} else {
		self.datetime = func() string {return ""}
	}
	return self
}

func (self * Stdout_t) Write(level string, format string, args ...interface{}) (err error) {
	_, err = fmt.Fprintf(os.Stdout, self.datetime() + level + " " + format + "\n", args...)
	return
}
