package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aldinokemal/go-oci/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"oras.land/oras-go/v2"
	"oras.land/oras-go/v2/content/oci"
	"oras.land/oras-go/v2/registry"
	"oras.land/oras-go/v2/registry/remote"
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

	// Parse the reference to extract registry, repository, and tag
	ref, err := registry.ParseReference(imageName)
	if err != nil {
		logrus.Fatalf("failed to parse reference: %v", err)
	}

	ctx := context.Background()

	// Create an OCI store from the saved tar file
	store, err := oci.NewFromTar(ctx, tmpPath)
	if err != nil {
		logrus.Fatalf("failed to create OCI store from tar: %v", err)
	}

	// Get repository name with registry
	regAddr := ref.Registry
	repoName := ref.Repository
	tagName = ref.Reference

	// Create a registry client
	repo, err := remote.NewRepository(fmt.Sprintf("%s/%s", regAddr, repoName))
	if err != nil {
		logrus.Fatalf("failed to create repository: %v", err)
	}

	// Configure HTTP options
	if insecure {
		repo.PlainHTTP = true
	}

	// Copy from OCI store to remote repository
	desc, err := oras.Copy(ctx, store, tagName, repo, tagName, oras.DefaultCopyOptions)
	if err != nil {
		logrus.Fatalf("failed to push image: %v", err)
	}

	logrus.Infof("image pushed successfully: %s@%s", imageName, desc.Digest)
}
