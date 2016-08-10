package login

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

var (
	TestLogger     bogusLogger = newBogusLogger()
	TestMiddleWare Service
)

type bogusLogger struct {
}

func newBogusLogger() bogusLogger {
	return bogusLogger{}
}

func (bl bogusLogger) Log(v ...interface{}) error {
	_, err := fmt.Println(v)
	return err
}

func TestLoggingMiddleWare(t *testing.T) {
	TestMiddleWare = LoggingMiddleware(TestLogger)(TestService)
}

func TestLoginMiddleWare(t *testing.T) {
	u, err := TestMiddleWare.Login("test", "test")
	spew.Dump(u)
	spew.Dump(err)
	spew.Dump(TestMiddleWare)
}
