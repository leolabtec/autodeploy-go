#!/bin/bash

set -e

REPO_RAW="https://github.com/leolabtec/autodeploy-go/releases/latest/download"
INSTALL_DIR="/opt/autodeploy-go"
BIN_NAME="autodeploy-linux-amd64"  # 依据系统选择对应的二进制文件

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

# ========== 1. 初始化目录结构 ==========
echo "[+] 创建主目录 $INSTALL_DIR..."
mkdir -p "$INSTALL_DIR"
cd "$INSTALL_DIR"

# ========== 2. 下载已编译的可执行文件 ==========
echo "[+] 下载已编译的可执行文件..."
curl -sSL "$REPO_RAW/$BIN_NAME" -o "$INSTALL_DIR/$BIN_NAME"

# 赋予执行权限
chmod +x "$INSTALL_DIR/$BIN_NAME"

# ========== 3. 启动 AutoDeploy 项目 ==========
echo "[+] 启动 AutoDeploy 项目..."
$INSTALL_DIR/$BIN_NAME

echo "[✓] AutoDeploy 启动成功！"
