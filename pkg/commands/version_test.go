package commands

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestRunVersion(t *testing.T) {
	out := capture(RunVersion)
	if out != "v0.1.0" {
		t.Errorf("unexpected output: %s", out)
	}
}

func capture(f func()) string {
	out := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = out
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
