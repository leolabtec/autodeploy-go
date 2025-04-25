// halo.go
package main

import (
	"fmt"
	"os"
	"os/exec"
)

// 部署 Halo 博客
func deployHalo(domain string) error {
	// 创建 Docker Compose 文件
	composeFile := fmt.Sprintf("/home/dockerdata/%s/docker-compose.yml", domain)
	file, err := os.Create(composeFile)
	if err != nil {
		return fmt.Errorf("创建 Docker Compose 文件失败: %v", err)
	}
	defer file.Close()

	// 写入 Halo 配置
	composeContent := fmt.Sprintf(`
version: '3'
services:
  halo:
    image: halohub/halo:2.15
    ports:
      - "8081:8080"
    environment:
      HALO_PORT: 8080
    networks:
      - halo_network
networks:
  halo_network:
    driver: bridge
`)
	_, err = file.WriteString(composeContent)
	if err != nil {
		return fmt.Errorf("写入 Docker Compose 文件失败: %v", err)
	}

	// 启动 Halo 容器
	err = startDockerContainer("halo", "halohub/halo:2.15", "8081:8080")
	if err != nil {
		return fmt.Errorf("启动 Halo 容器失败: %v", err)
	}

	fmt.Printf("Halo 博客站点部署完成: http://%s:8081\n", domain)
	return nil
}
