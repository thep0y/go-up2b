#!/bin/bash

set -ex

echo '打包 mac 版'
go build -o go-up2b_macOS_amd64 main.go

echo '打包 windows 版'
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o go-up2b_windows_x64.exe main.go

echo '打包 linux 版'
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-up2b_linux_amd64 main.go

echo '打包完成'
