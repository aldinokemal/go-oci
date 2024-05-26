## Push Image to OCI Registry


### Requirements
- [docker](https://docs.docker.com/get-docker/)
- [oras](https://oras.land/)

### Features
- Push Image to OCI Registry
- Push Manifest Multi-Architecture Docker Image to OCI Registry

### Installation
```bash
go install github.com/aldinokemal/go-oci@latest
```

### Command: Push Image from Local to Zot OCI Registry
`go-oci image:push <image:tag>`

```bash
go-oci image:push localhost:5000/go-whatsapp-web-multidevice:linux-arm64
```

### Command: Push Manifest from Local to Zot OCI Registry
```bash
go-oci manifest:push localhost:5000/go-whatsapp-web-multidevice:latest --insecure
```


### Command: Create Manifest and Push to Zot OCI Registry
```bash
go-oci manifest:create --insecure=true --push=true raspberrypi.tail1d23a.ts.net:5000/go-whatsapp-web-multidevice:latest \
        --amend raspberrypi.tail1d23a.ts.net:5000/go-whatsapp-web-multidevice:linux-amd64 \
        --amend raspberrypi.tail1d23a.ts.net:5000/go-whatsapp-web-multidevice:linux-arm64
```

### Reference:
```bash
docker buildx build \
        --tag raspberrypi.tail1d23a.ts.net:5000/go-whatsapp-web-multidevice:linux-amd64 \
        --platform linux/amd64 \
        --load \
        --progress plain \
        -f ./docker/golang.Dockerfile \
        .

docker buildx build \
        --tag raspberrypi.tail1d23a.ts.net:5000/go-whatsapp-web-multidevice:linux-arm64 \
        --platform linux/arm64 \
        --load \
        --progress plain \
        -f ./docker/golang.Dockerfile \
        .


docker manifest create --insecure raspberrypi.tail1d23a.ts.net:5000/go-whatsapp-web-multidevice:latest \
        --amend raspberrypi.tail1d23a.ts.net:5000/go-whatsapp-web-multidevice:linux-amd64 \
        --amend raspberrypi.tail1d23a.ts.net:5000/go-whatsapp-web-multidevice:linux-arm64
```

