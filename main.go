// main.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "AutoDeploy",
		Usage: "自动化部署 WordPress 和 Halo 博客",
		Commands: []*cli.Command{
			{
				Name:    "wordpress",
				Usage:   "部署 WordPress 站点",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "domain",
						Usage:    "指定 WordPress 站点的域名",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					domain := c.String("domain")
					return deployWordPress(domain)
				},
			},
			{
				Name:    "halo",
				Usage:   "部署 Halo 博客",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "domain",
						Usage:    "指定 Halo 博客的域名",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					domain := c.String("domain")
					return deployHalo(domain)
				},
			},
		},
	}

	// 运行应用程序
	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
