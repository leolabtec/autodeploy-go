#!/bin/bash

set -e

REPO_RAW="https://raw.githubusercontent.com/leolabtec/autodeploy/main"
INSTALL_DIR="/opt/autodeploy"
GO_VERSION="1.18"  # 根据需要更新版本

# ========== 0. 检查依赖项 ==========
check_dep() {
  if ! command -v "$1" &>/dev/null; then
    echo "[-] 缺少必要依赖：$1，请先安装后再运行本脚本。"
    exit 1
  fi
}

echo "[+] 正在检查并准备系统关键依赖..."
check_dep go
check_dep docker
check_dep docker-compose
check_dep crontab

# ========== 1. 初始化目录结构 ==========
echo "[+] 创建主目录 $INSTALL_DIR..."
mkdir -p "$INSTALL_DIR"
cd "$INSTALL_DIR"

# ========== 2. 拉取所有 Go 文件 ==========
echo "[+] 拉取所有 Go 文件..."
for file in main.go docker_ops.go caddy.go wordpress.go halo.go backup.go delete.go restore.go uninstall.go shortcut.go; do
  curl -sS "$REPO_RAW/$file" -o "$INSTALL_DIR/$file"
done

# ========== 3. 拉取和启动 Caddy 容器 ==========
echo "[+] 启动 Caddy 容器..."
docker rm -f caddy 2>/dev/null || true
docker run -d \
  --name caddy \
  --restart=unless-stopped \
  --network host \
  -v /home/dockerdata/docker_caddy/Caddyfile:/etc/caddy/Caddyfile \
  -v /home/dockerdata/docker_caddy:/data \
  -v /home/dockerdata/docker_caddy:/config \
  caddy:2.7.6
echo "[✓] Caddy 已启动"

# ========== 4. 设置巡检任务 ==========
echo "[+] 设置 Caddy 巡检任务..."
CRON_JOB="*/5 * * * * $INSTALL_DIR/go run $INSTALL_DIR/monitor.go"
if crontab -l 2>/dev/null | grep -F "$CRON_JOB" > /dev/null; then
  echo "[i] 巡检任务已存在"
else
  (crontab -l 2>/dev/null; echo "$CRON_JOB") | crontab -
  echo "[✓] 已添加巡检任务"
fi

# ========== 5. 启动 main.go ==========
echo "[✓] 环境部署完成，正在启动 AutoDeploy 主菜单..."
go run main.go
