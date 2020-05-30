package tests

import (
	"io/ioutil"
	"os"
)

// CaptureOutput will get output and convert it to string
func CaptureOutput(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	return string(out)
}
