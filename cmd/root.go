package cmd

import (
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
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func Execute() error {
	return rootCmd.Execute()
}
