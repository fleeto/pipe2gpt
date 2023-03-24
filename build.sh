#!/bin/bash

# 创建 bin 文件夹（如果不存在）
mkdir -p bin

# 编译代码并将二进制文件输出到 bin/pipe2gpt
go build -o bin/pipe2gpt cmd/main.go

echo "Build complete. Binary generated in bin/pipe2gpt"
