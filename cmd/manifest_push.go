package cmd

import (
	"fmt"
	"github.com/aldinokemal/go-oci/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var manifestPushCmd = &cobra.Command{
	Use:   "manifest:push",
	Short: "Push docker manifest to a registry using oras",
	Run:   manifestPushCmdRun,
}

func init() {
	rootCmd.AddCommand(manifestPushCmd)
	manifestPushCmd.PersistentFlags().BoolVarP(&insecure, "insecure", "i", false, "allow http  --insecure=true | -i=true")
}

func manifestPushCmdRun(cmd *cobra.Command, args []string) {
	dockerManifest := args[0]
	tmpFolder, err := utils.CreateTmpFolder()
	if err != nil {
		logrus.Fatalf("failed to create tmp folder: %v", err)
	}
	defer func() {
		logrus.Debugf("removing tmp folder: %s", tmpFolder)
		err := os.RemoveAll(tmpFolder)
		if err != nil {
			logrus.Errorf("failed to remove tmp folder: %v", err)
		}
	}()

	manifestPath := fmt.Sprintf("%s/%s", tmpFolder, "manifest.json")
	exportManifestCmd := fmt.Sprintf("docker manifest inspect %s > %s", dockerManifest, manifestPath)

	if err = utils.RunCommand(exportManifestCmd); err != nil {
		logrus.Fatalf("failed to export manifest: %v", err)
	}

	// push manifest using oras
	var orasPushManifestStrBuilder strings.Builder

	orasPushManifestStrBuilder.WriteString("oras manifest push ")
	if insecure {
		orasPushManifestStrBuilder.WriteString("--plain-http ")
	}
	orasPushManifestStrBuilder.WriteString(fmt.Sprintf("%s %s", dockerManifest, manifestPath))

	orasPushManifestCmd := orasPushManifestStrBuilder.String()
	logrus.Debugf("pushing manifest: %s", orasPushManifestCmd)

	if err = utils.RunCommand(orasPushManifestCmd); err != nil {
		logrus.Fatalf("failed to push manifest: %v", err)
	}

	logrus.Infof("manifest pushed successfully")
}
