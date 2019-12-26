package log

import "time"
import "testing"

func Example_log() {
	logger := NewEmpty()
	logger.AddOutput("stdout", 0, NewStdout(""))
	log_file, _ := NewFile("/tmp/test.log", "", 1024, 10)
	logger.AddOutput("file", 0, log_file)
	log_http := NewHttp(DefaultTransport(time.Second), 10, 1, "http://localhost", Convert, nil)
	logger.AddOutput("http", 0, log_http)
	SetLogger(logger)
	Debug("lalala")
	Debug("bububu")
/* Output:
DEBUG lalala
DEBUG bububu
*/
}

func TestLog1(t * testing.T) {

}
