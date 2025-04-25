#!/bin/bash

set -e

REPO_RAW="https://raw.githubusercontent.com/leolabtec/autodeploy-go/main"
INSTALL_DIR="/opt/autodeploy-go"

# ========== 0. 检查依赖项 ==========
check_dep() {
  if ! command -v "$1" &>/dev/null; then
    echo "[-] 缺少必要依赖：$1，请先安装后再运行本脚本。"
    exit 1
  fi
}

echo "[+] 正在检查并准备系统关键依赖..."
check_dep docker
check_dep docker-compose
check_dep go

# ========== 1. 初始化目录结构 ==========
echo "[+] 创建主目录 $INSTALL_DIR..."
mkdir -p "$INSTALL_DIR"
cd "$INSTALL_DIR"

# ========== 2. 安装 Go 项目依赖 ==========
echo "[+] 获取 Go 项目代码..."

# 拉取项目代码
git clone https://github.com/leolabtec/autodeploy-go.git .

# ========== 3. 安装 Go 依赖 ==========
echo "[+] 安装 Go 项目依赖..."
go mod tidy

# ========== 4. 启动主程序 ==========
echo "[+] 启动 AutoDeploy 主程序..."
go run main.go

echo "[✓] AutoDeploy 启动成功！"
