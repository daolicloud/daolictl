package utils

import (
	"io"
	"os"
)

// StdStreams returns the standard streams (stdin, stdout, stderr).
func StdStreams() (stdIn io.ReadCloser, stdOut, stdErr io.Writer) {
	return os.Stdin, os.Stdout, os.Stderr
}
