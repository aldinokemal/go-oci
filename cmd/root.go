package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "go-oci",
	Short: "A simple tool to push OCI images to a registry",
	Long:  `A simple tool to push OCI images to a registry`,
}

func Execute() error {
	return rootCmd.Execute()
}
