package halo

import (
	"fmt"
	"autodeploy-go/dockerops"
)

// 创建 Halo 博客
func CreateHaloSite() {
	// 创建 Halo 博客所需的容器
	fmt.Println("[+] 创建 Halo 博客...")
	// 启动 Docker 容器
	dockerops.StartContainer("halo", "8090:80", "halohub/halo:latest")
}
