package cmd

import (
	"io"
	"os"

	"github.com/aldinokemal/go-oci/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	verbose  bool
	push     bool
	insecure bool
	amend    []string
)

var rootCmd = &cobra.Command{
	Use:   "go-oci",
	Short: "A simple tool to push OCI images to a registry",
	Long:  `A simple tool to push OCI images to a registry`,
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	cobra.OnInitialize(setLogLevel)
}

func setLogLevel() {
	if utils.IsRunningFromGoRun() {
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		f, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			logrus.SetOutput(os.Stdout)
			logrus.Errorf("failed to open log file: %v", err)
		} else {
			logrus.SetOutput(io.MultiWriter(os.Stdout, f))
		}
	}
	logrus.Debugf("is running from go run: %v", utils.IsRunningFromGoRun())
}

func Execute() error {
	return rootCmd.Execute()
}
