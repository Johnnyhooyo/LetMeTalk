#!/bin/bash

# 脚本退出时显示错误
set -e

# 获取用户输入
read -r -p "Please enter server IP: " SERVER_IP
read -r -p "Please enter server port: " SERVER_PORT
read -r -p "Please enter your name: " USER_NAME

# 编译项目
echo "Building the Go project..."
go build -o myclient .

# 运行项目
echo "Running the Go project..."
./myclient "$SERVER_IP" "$SERVER_PORT" "$USER_NAME"
