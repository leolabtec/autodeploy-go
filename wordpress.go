package wordpress

import (
	"fmt"
	"autodeploy-go/dockerops"
)

// 创建 WordPress 站点
func CreateWordPressSite() {
	// 创建 WordPress 所需的容器
	fmt.Println("[+] 创建 WordPress 站点...")
	// 启动 Docker 容器
	dockerops.StartContainer("wordpress", "8080:80", "wordpress:latest")
}
