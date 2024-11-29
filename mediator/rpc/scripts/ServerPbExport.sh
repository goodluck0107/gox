#!/bin/bash

rootDirect=$(pwd)

echo "执行目录 $rootDirect"

protoc  --proto_path=../proto --go_out=.. ../proto/*.proto

echo "按回车键任意键退出"
read