package cmd

import (
	"strings"

	"github.com/aldinokemal/go-oci/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var manifestCreateCmd = &cobra.Command{
	Use:   "manifest:create",
	Short: "Create a manifest for a docker image multi-arch",
	Run:   manifestCreateCmdRun,
}

func init() {
	rootCmd.AddCommand(manifestCreateCmd)
	manifestCreateCmd.PersistentFlags().BoolVarP(&insecure, "insecure", "i", false, "allow http  --insecure=true | -i=true")
	manifestCreateCmd.PersistentFlags().BoolVarP(&push, "push", "p", false, "push the manifest after creation")
	manifestCreateCmd.PersistentFlags().StringSliceVarP(&amend, "amend", "a", nil, "amend the manifest with the specified platform")
}

func manifestCreateCmdRun(cmd *cobra.Command, args []string) {
	if len(amend) == 0 {
		logrus.Fatalf("amend flag is required")
	}

	dockerManifest := args[0]

	// remove current manifest if exists
	if err := utils.RunCommand("docker manifest rm " + dockerManifest); err != nil {
		logrus.Warnf("no manifest: %v", err)
	}

	var dockerManifestStrBuilder strings.Builder
	dockerManifestStrBuilder.WriteString("docker manifest create ")
	if insecure {
		dockerManifestStrBuilder.WriteString("--insecure ")
	}

	dockerManifestStrBuilder.WriteString(dockerManifest + " ")

	for _, platform := range amend {
		dockerManifestStrBuilder.WriteString("--amend ")
		dockerManifestStrBuilder.WriteString(platform)
		dockerManifestStrBuilder.WriteString(" ")
	}

	logrus.Debugf("creating manifest: %s", dockerManifestStrBuilder.String())

	if err := utils.RunCommand(dockerManifestStrBuilder.String()); err != nil {
		logrus.Fatalf("failed to create manifest: %v", err)
	}

	if push {
		manifestPushCmdRun(cmd, []string{dockerManifest})
	}
}
