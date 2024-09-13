package hello_test

import (
	"bytes"
	"io"
	"menta/hello"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	oldStdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	hello.Hello("Hello, World!")

	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	os.Stdout = oldStdout

	assert.Equal(t, "Hello, World!\n", buf.String())
}
