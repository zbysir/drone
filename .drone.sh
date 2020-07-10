#!/bin/sh

# clone the extras project.
set -e
set -x

# 注意! 不能设置环境变量 CGO_ENABLED=0, 否则sqlite3无法被import. 目前win上可能无法编译此项目, 请在linux上编译.
# 由于需要运行在alpine上, 所以需要-extldflags "-static"
# -ldflags "-extldflags -static" at the end makes sure C code is statically linked so resulting binary truly has no dependencies even for C code.
go build -ldflags '-extldflags "-static" -X github.com/drone/drone/version.VersionDev=build.0.02' -o release/drone-server github.com/drone/drone/cmd/drone-server
# GOOS=linux GOARCH=amd64 CGO_ENABLED=0         go build -ldflags '-X github.com/drone/drone/version.VersionDev=build.'0.01 -o release/drone-agent             github.com/drone/drone/cmd/drone-agent
# GOOS=linux GOARCH=arm64 CGO_ENABLED=0         go build -ldflags '-X github.com/drone/drone/version.VersionDev=build.'0.01 -o release/linux/arm64/drone-agent github.com/drone/drone/cmd/drone-agent
# GOOS=linux GOARCH=arm   CGO_ENABLED=0 GOARM=7 go build -ldflags '-X github.com/drone/drone/version.VersionDev=build.'0.01 -o release/linux/arm/drone-agent   github.com/drone/drone/cmd/drone-agent
