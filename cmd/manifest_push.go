package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aldinokemal/go-oci/utils"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"oras.land/oras-go/v2"
	"oras.land/oras-go/v2/content/memory"
	"oras.land/oras-go/v2/registry"
	"oras.land/oras-go/v2/registry/remote"
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

	// Read the manifest file
	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		logrus.Fatalf("failed to read manifest file: %v", err)
	}

	// Parse the manifest
	var manifestContent interface{}
	if err := json.Unmarshal(manifestBytes, &manifestContent); err != nil {
		logrus.Fatalf("failed to parse manifest: %v", err)
	}

	// Create a new memory store
	store := memory.New()

	// Create a context
	ctx := context.Background()

	// Create descriptor for the manifest
	manifestDesc := ocispec.Descriptor{
		MediaType: ocispec.MediaTypeImageManifest,
		Digest:    "",
		Size:      int64(len(manifestBytes)),
		Annotations: map[string]string{
			ocispec.AnnotationTitle: "manifest.json",
		},
	}

	// Push to the memory store
	if err := store.Push(ctx, manifestDesc, strings.NewReader(string(manifestBytes))); err != nil {
		logrus.Fatalf("failed to push manifest to memory store: %v", err)
	}

	// Parse the reference
	ref, err := registry.ParseReference(dockerManifest)
	if err != nil {
		logrus.Fatalf("failed to parse reference: %v", err)
	}

	// Create a repository client
	repo, err := remote.NewRepository(fmt.Sprintf("%s/%s", ref.Registry, ref.Repository))
	if err != nil {
		logrus.Fatalf("failed to create repository client: %v", err)
	}

	// Set HTTP options
	if insecure {
		repo.PlainHTTP = true
	}

	// Copy the manifest to the remote repository
	desc, err := oras.Copy(ctx, store, manifestDesc.Annotations[ocispec.AnnotationTitle], repo, ref.Reference, oras.DefaultCopyOptions)
	if err != nil {
		logrus.Fatalf("failed to push manifest: %v", err)
	}

	logrus.Infof("manifest pushed successfully: %s@%s", dockerManifest, desc.Digest)
}
