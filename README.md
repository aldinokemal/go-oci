# Push Image to OCI Registry

## Production Ready
[Download Release](https://github.com/aldinokemal/go-oci/releases)

## Requirements (for development)

- [docker](https://docs.docker.com/get-docker/)
- [oras](https://oras.land/)

## Features

- Push Image to OCI Registry
- Push Manifest Multi-Architecture Docker Image to OCI Registry

## Installation

```bash
go install github.com/aldinokemal/go-oci@latest
```

## Command: Push Image from Local to Zot OCI Registry

`go-oci image:push <image:tag>`

```bash
go-oci image:push localhost:5000/go-whatsapp-web-multidevice:linux-arm64 \
        --insecure=true
```

## Command: Push Manifest from Local to Zot OCI Registry

```bash
go-oci manifest:push localhost:5000/go-whatsapp-web-multidevice:latest \
        --insecure=true
```

## Command: Create Manifest and Push to Zot OCI Registry

```bash
go-oci manifest:create localhost:5000/go-whatsapp-web-multidevice:latest \
        --amend localhost:5000/go-whatsapp-web-multidevice:linux-amd64 \
        --amend localhost:5000/go-whatsapp-web-multidevice:linux-arm64 \
        --insecure=true \
        --push=true 
```

## Reference

```bash
docker buildx build \
        --tag localhost:5000/go-whatsapp-web-multidevice:linux-amd64 \
        --platform linux/amd64 \
        --load \
        --progress plain \
        -f ./docker/golang.Dockerfile \
        .

docker buildx build \
        --tag localhost:5000/go-whatsapp-web-multidevice:linux-arm64 \
        --platform linux/arm64 \
        --load \
        --progress plain \
        -f ./docker/golang.Dockerfile \
        .

docker manifest create --insecure localhost:5000/go-whatsapp-web-multidevice:latest \
        --amend localhost:5000/go-whatsapp-web-multidevice:linux-amd64 \
        --amend localhost:5000/go-whatsapp-web-multidevice:linux-arm64
```

## Picture

- Uploaded multi arch image to zot registry

![upload to zot registry](https://github.com/aldinokemal/go-oci/assets/14232125/dc5ead48-ffbe-43da-a651-83932f695cdc)
