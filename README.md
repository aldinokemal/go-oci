# Go-OCI: OCI Container Registry Tool

## Overview

Go-OCI is a simple CLI tool for working with OCI (Open Container Initiative) registries. It allows you to:

- Push Docker images to OCI registries
- Create and push multi-architecture manifest lists

## Installation

```bash
go install github.com/aldinokemal/go-oci@latest
```

Or [download a pre-built release](https://github.com/aldinokemal/go-oci/releases)

## Commands

### Push Docker Image to OCI Registry

```bash
go-oci image:push <image:tag> [flags]

# Example
go-oci image:push localhost:5000/my-application:linux-arm64 --insecure=true
```

Flags:

- `--insecure, -i`: Allow HTTP connections (default: false)
- `--verbose, -v`: Enable verbose logging (default: false)

### Push Existing Manifest to OCI Registry

```bash
go-oci manifest:push <image:tag> [flags]

# Example
go-oci manifest:push localhost:5000/my-application:latest --insecure=true
```

Flags:

- `--insecure, -i`: Allow HTTP connections (default: false)
- `--verbose, -v`: Enable verbose logging (default: false)

### Create Multi-Architecture Manifest and Push to OCI Registry

```bash
go-oci manifest:create <image:tag> [flags]

# Example
go-oci manifest:create localhost:5000/my-application:latest \
        --amend localhost:5000/my-application:linux-amd64 \
        --amend localhost:5000/my-application:linux-arm64 \
        --insecure=true \
        --push=true
```

Flags:

- `--amend, -a`: Add images to the manifest (required, can be specified multiple times)
- `--insecure, -i`: Allow HTTP connections (default: false)
- `--push, -p`: Push the manifest after creation (default: false)
- `--verbose, -v`: Enable verbose logging (default: false)

## Example Workflow

- Build multi-architecture Docker images:

```bash
# Build AMD64 image
docker buildx build \
        --tag localhost:5000/my-application:linux-amd64 \
        --platform linux/amd64 \
        --load \
        -f ./Dockerfile \
        .

# Build ARM64 image
docker buildx build \
        --tag localhost:5000/my-application:linux-arm64 \
        --platform linux/arm64 \
        --load \
        -f ./Dockerfile \
        .
```

- Push images to the registry:

```bash
go-oci image:push localhost:5000/my-application:linux-amd64 --insecure=true
go-oci image:push localhost:5000/my-application:linux-arm64 --insecure=true
```

- Create and push a multi-architecture manifest:

```bash
go-oci manifest:create localhost:5000/my-application:latest \
        --amend localhost:5000/my-application:linux-amd64 \
        --amend localhost:5000/my-application:linux-arm64 \
        --insecure=true \
        --push=true
```

## Screenshot

- Multi-architecture images in a Zot OCI registry:

![upload to zot registry](https://github.com/aldinokemal/go-oci/assets/14232125/dc5ead48-ffbe-43da-a651-83932f695cdc)
