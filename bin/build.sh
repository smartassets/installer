#!/bin/bash -eux

function build() {
    local version=$1
    local platform=$2
    local arch=$3
    local name=$4

    echo calling to build for $platform $arch
    GOOS=$platform GOARCH=$arch go build \
        -ldflags "-X main.Version=${version}" \
        -o ${name}
}

function moveArtifactsToBuildFolder() {
    local folder=$1
    mv $ASSETS_NAME_WIN_32 $folder
    mv $ASSETS_NAME_WIN_64 $folder
    mv $ASSETS_NAME_LINUX_32 $folder
    mv $ASSETS_NAME_LINUX_64 $folder
    mv "$ASSETS_NAME_OSX" "$folder"
}

BUILD_FOLDER=build
ASSETS_NAME_WIN_32=assets.win32
ASSETS_NAME_WIN_64=assets.win64
ASSETS_NAME_LINUX_32=assets.linux32
ASSETS_NAME_LINUX_64=assets.linux64
ASSETS_NAME_OSX=assets.osx

version=${1}
echo "This will be the version $version"
build $version linux 386 $ASSETS_NAME_LINUX_32
build $version linux amd64 $ASSETS_NAME_LINUX_64
build $version windows 386 $ASSETS_NAME_WIN_32
build $version windows amd64 $ASSETS_NAME_WIN_64
build $version darwin amd64 $ASSETS_NAME_OSX

mkdir -p $BUILD_FOLDER
moveArtifactsToBuildFolder $BUILD_FOLDER
cd $BUILD_FOLDER