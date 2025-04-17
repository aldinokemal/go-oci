package main

import (
	"os"
	"runtime"

	"github.com/aldinokemal/go-oci/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("panic: %v", r)
			logrus.Errorf("stacktrace:\n%s", getStackTrace())
			// exit with error code
			// use os.Exit(1) to ensure non-zero exit
			os.Exit(1)
		}
	}()
	_ = cmd.Execute()
}

// getStackTrace returns the current stack trace as a string
func getStackTrace() string {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}
