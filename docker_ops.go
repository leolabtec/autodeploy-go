// docker_ops.go
package main

import (
	"fmt"
	"os/exec"
)

// 启动 Docker 容器
func startDockerContainer(containerName, imageName string, portMapping string) error {
	cmd := exec.Command("docker", "run", "-d", "--name", containerName, "-p", portMapping, imageName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("启动容器失败: %v", err)
	}
	fmt.Printf("容器启动成功: %s\n", output)
	return nil
}

// 停止 Docker 容器
func stopDockerContainer(containerName string) error {
	cmd := exec.Command("docker", "stop", containerName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("停止容器失败: %v", err)
	}
	fmt.Printf("容器已停止: %s\n", output)
	return nil
}
