package main

import (
	"fmt"
	"os"
	"os/exec"
	"autodeploy-go/backup"
	"autodeploy-go/delete"
	"autodeploy-go/restore"
	"autodeploy-go/wordpress"
	"autodeploy-go/halo"
	"autodeploy-go/utils"
)

func showMenu() {
	fmt.Println("\n📋 请选择操作：")
	fmt.Println("1. 创建 WordPress 站点")
	fmt.Println("2. 创建 Halo 博客")
	fmt.Println("3. 备份所有站点数据")
	fmt.Println("4. 删除已部署站点")
	fmt.Println("5. 恢复备份环境")
	fmt.Println("6. 卸载部署系统")
	fmt.Println("7. 配置系统")
	fmt.Println("0. 退出")
}

func main() {
	for {
		showMenu()

		var choice string
		fmt.Print("请输入编号: ")
		_, err := fmt.Scanf("%s", &choice)
		if err != nil {
			fmt.Println("输入无效，退出程序")
			return
		}

		switch choice {
		case "1":
			wordpress.CreateWordPressSite()
		case "2":
			halo.CreateHaloSite()
		case "3":
			backup.CreateBackup()
		case "4":
			delete.DeleteSite()
		case "5":
			restore.RestoreBackup()
		case "6":
			// 卸载系统操作
			utils.UninstallSystem()
		case "7":
			// 配置系统操作
			utils.ConfigureSystem()
		case "0":
			fmt.Println("👋 再见！")
			return
		default:
			fmt.Println("无效选择，请重新选择！")
		}
	}
}
