package dockerops

import (
	"fmt"
	"os/exec"
)

// 启动 Docker 容器
func StartContainer(containerName string, port string, image string) {
	cmd := exec.Command("docker", "run", "-d", "--name", containerName, "-p", port, image)
	err := cmd.Run()
	if err != nil {
		fmt.Println("启动容器失败:", err)
		return
	}
	fmt.Println("容器已启动:", containerName)
}

// 停止 Docker 容器
func StopContainer(containerName string) {
	cmd := exec.Command("docker", "stop", containerName)
	err := cmd.Run()
	if err != nil {
		fmt.Println("停止容器失败:", err)
		return
	}
	fmt.Println("容器已停止:", containerName)
}

// 删除 Docker 容器
func RemoveContainer(containerName string) {
	cmd := exec.Command("docker", "rm", containerName)
	err := cmd.Run()
	if err != nil {
		fmt.Println("删除容器失败:", err)
		return
	}
	fmt.Println("容器已删除:", containerName)
}
