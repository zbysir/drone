#!/bin/sh

# clone the extras project.
set -e
set -x

# 由于要运行在alpine, 所以需要加GOOS=linux GOARCH=arm64 CGO_ENABLED=0
# 由于drone-server使用到了sqlite3, sqlite3使用到了gcc, 所以需要-extldflags "-static"
# -ldflags "-extldflags -static" at the end makes sure C code is statically linked so resulting binary truly has no dependencies even for C code.
GOPATH=Z:\\go_path\\src GOOS=linux GOARCH=amd64 CGO_ENABLED=0  go build -ldflags '-extldflags "-static" -X github.com/drone/drone/version.VersionDev=build.'0.01 -o release/drone-server github.com/drone/drone/cmd/drone-server
# GOOS=linux GOARCH=amd64 CGO_ENABLED=0         go build -ldflags '-X github.com/drone/drone/version.VersionDev=build.'0.01 -o release/drone-agent             github.com/drone/drone/cmd/drone-agent
# GOOS=linux GOARCH=arm64 CGO_ENABLED=0         go build -ldflags '-X github.com/drone/drone/version.VersionDev=build.'0.01 -o release/linux/arm64/drone-agent github.com/drone/drone/cmd/drone-agent
# GOOS=linux GOARCH=arm   CGO_ENABLED=0 GOARM=7 go build -ldflags '-X github.com/drone/drone/version.VersionDev=build.'0.01 -o release/linux/arm/drone-agent   github.com/drone/drone/cmd/drone-agent
