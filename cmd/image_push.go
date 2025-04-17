package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/aldinokemal/go-oci/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var imagePushCmd = &cobra.Command{
	Use:   "image:push",
	Short: "Push an OCI image to a registry",
	Run:   imagePushCmdRun,
}

func init() {
	rootCmd.AddCommand(imagePushCmd)
	imagePushCmd.PersistentFlags().BoolVarP(&insecure, "insecure", "i", false, "allow http  --insecure=true | -i=true")
}

func imagePushCmdRun(cmd *cobra.Command, args []string) {
	logrus.Debugf("args: %v", args)
	if len(args) < 1 {
		logrus.Fatalf("image name argument is required")
	}
	imageName := args[0]
	// image have to contain tag
	if !strings.Contains(imageName, ":") {
		logrus.Fatalf("image name must contain tag")
	}
	arrayImageName := strings.Split(imageName, ":")
	// tag is last element
	tagName := arrayImageName[len(arrayImageName)-1]
	// if tag contains /, it means it is a digest
	if strings.Contains(tagName, "/") {
		logrus.Fatalf("image name must contain tag")
	}

	tmpFolder, err := utils.CreateTmpFolder()
	if err != nil {
		logrus.Fatalf("failed to create tmp folder: %v", err)
	}

	logrus.Debugf("tmp folder: %s", tmpFolder)

	defer func() {
		logrus.Debugf("removing tmp folder: %s", tmpFolder)
		err := os.RemoveAll(tmpFolder)
		if err != nil {
			logrus.Errorf("failed to remove tmp folder: %v", err)
		}
	}()

	tmpPath := fmt.Sprintf("%s/%s.tar", tmpFolder, utils.RandomString(10))
	commandDockerSave := fmt.Sprintf("docker save %s -o %s", imageName, tmpPath)

	logrus.Debugf("saving image to tar: %s", commandDockerSave)

	if err = utils.RunCommand(commandDockerSave); err != nil {
		logrus.Fatalln("failed to save image to tar")
	}

	var orasCopyStrBuilder strings.Builder
	orasCopyStrBuilder.WriteString("oras cp --from-oci-layout ")
	if insecure {
		orasCopyStrBuilder.WriteString("--to-plain-http ")
	}
	orasCopyStrBuilder.WriteString(fmt.Sprintf("--from-oci-layout %s:%s %s", tmpPath, tagName, imageName))

	commandOrasCopy := orasCopyStrBuilder.String()
	logrus.Debugf("pushing image: %s", commandOrasCopy)

	if err = utils.RunCommand(commandOrasCopy); err != nil {
		logrus.Fatalf("failed to push image: %v", err)
	}

	logrus.Infof("image pushed successfully")
}
