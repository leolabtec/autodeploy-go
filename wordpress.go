// wordpress.go
package main

import (
	"fmt"
	"os"
	"os/exec"
)

// 部署 WordPress
func deployWordPress(domain string) error {
	// 创建 Docker Compose 文件
	composeFile := fmt.Sprintf("/home/dockerdata/%s/docker-compose.yml", domain)
	file, err := os.Create(composeFile)
	if err != nil {
		return fmt.Errorf("创建 Docker Compose 文件失败: %v", err)
	}
	defer file.Close()

	// 写入 WordPress 配置
	composeContent := fmt.Sprintf(`
version: '3'
services:
  wordpress:
    image: wordpress:latest
    ports:
      - "8080:80"
    environment:
      WORDPRESS_DB_HOST: db:3306
      WORDPRESS_DB_NAME: wordpress
      WORDPRESS_DB_USER: root
      WORDPRESS_DB_PASSWORD: example
    networks:
      - wp_network
  db:
    image: mariadb:latest
    environment:
      MYSQL_ROOT_PASSWORD: example
    networks:
      - wp_network
networks:
  wp_network:
    driver: bridge
`)
	_, err = file.WriteString(composeContent)
	if err != nil {
		return fmt.Errorf("写入 Docker Compose 文件失败: %v", err)
	}

	// 启动 WordPress 容器
	err = startDockerContainer("wordpress", "wordpress:latest", "8080:80")
	if err != nil {
		return fmt.Errorf("启动 WordPress 容器失败: %v", err)
	}

	fmt.Printf("WordPress 站点部署完成: http://%s:8080\n", domain)
	return nil
}
