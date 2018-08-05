package tricks

import (
	"io"
	"testing"
	"os"
)

type WriterOnly struct {
	io.Writer
}

type WriteCloser struct {
	io.WriteCloser
}

// Cannot convert WriterOnly to io.WriteCloser even though it holds WriteCloser
// it leads to this kind of workarounds: https://github.com/newrelic/go-agent/blob/a923caa59930cbb1b27d2a9fb2d5663d9e5c637c/internal_response_writer.go
func TestConversion(t *testing.T) {
	var wc io.WriteCloser = WriteCloser{WriteCloser: os.Stdout}
	var wo io.Writer = WriterOnly{Writer: wc}

	_, ok := wo.(WriteCloser)
	if ok {
		t.Error("could covert WriterOnly to WriteCloser")
	}
	_, ok = wo.(WriterOnly).Writer.(WriteCloser)
	if !ok {
		t.Error("could not convert Writer to WriteCloser")
	}
}

// see: https://golang.org/doc/faq#nil_error
func TestNilIsNotNil(t *testing.T) {
	Err := func() error {
		var err *os.PathError = nil
		return err
	}
	err := Err()
	t.Logf("%#v", err)

	if err == nil {
		t.Errorf("%+v must be a nil", err)
	}

	shouldPanic(t, func() { err.Error() }, "runtime error: invalid memory address or nil pointer dereference")
}
