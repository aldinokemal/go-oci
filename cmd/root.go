package cmd

import "github.com/spf13/cobra"

var (
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
}

func Execute() error {
	return rootCmd.Execute()
}
