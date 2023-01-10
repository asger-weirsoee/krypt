#!/bin/bash

mkdir -p ./build
rm -r ./build/*

platforms=("linux/amd64")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name="./build/krypt"
	env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name krypt.go
	fyne-cross windows -arch=amd64 -silent
	mv ./fyne-cross/bin/* ./build
done

