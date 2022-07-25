package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "repo",
				Aliases: []string{"r"},
				Usage:   "数据存储的代码仓库",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "新增数据存储的代码仓库",
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:  "default",
								Value: true,
								Usage: "是否是默认的代码仓库",
							},
							&cli.StringFlag{
								Name:        "url",
								Usage:       "代码仓库的地址。默认应具有代码仓库的读写权限",
								Required:    true,
								DefaultText: "git@github.com:zhangyucumt/password-tool.git",
							},
							&cli.StringFlag{
								Name:        "name",
								Usage:       "代码仓库的名称标识。",
								Required:    true,
								DefaultText: "repo1",
							},
						},
						Action: func(cCtx *cli.Context) error {
							fmt.Println("new task template: ", cCtx.Args().First())
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "移除数据存储的代码仓库",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "name",
								Usage:       "代码仓库的名称标识。",
								DefaultText: "repo1",
							},
						},
						Action: func(cCtx *cli.Context) error {
							fmt.Println("removed task template: ", cCtx.Args().First())
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "数据存储的代码仓库列表",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("removed task template: ", cCtx.Args().First())
							return nil
						},
					},
					{
						Name:  "set-default",
						Usage: "设置某个代码仓库为默认的数据存储仓库",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "name",
								Usage:       "代码仓库的名称标识。",
								Required:    true,
								DefaultText: "repo1",
							},
						},
						Action: func(cCtx *cli.Context) error {
							fmt.Println("removed task template: ", cCtx.Args().First())
							return nil
						},
					},
					{
						Name:  "upload",
						Usage: "将本地存储的数据同步至代码仓库。默认同步的仓库是默认仓库",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "name",
								Usage:       "需要同步的代码仓库的名称标识。",
								DefaultText: "repo1",
							},
						},
						Action: func(cCtx *cli.Context) error {
							fmt.Println("removed task template: ", cCtx.Args().First())
							return nil
						},
					},
				},
				Action: func(cCtx *cli.Context) error {
					fmt.Println("removed task template: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:  "ssh",
				Usage: "ssh",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "lang",
						Value:    "english",
						Usage:    "language for the greeting",
						Required: true,
					},
				},
				Action: func(cCtx *cli.Context) error {
					fmt.Println("added task: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("completed task: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "options for task templates",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("new task template: ", cCtx.Args().First())
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("removed task template: ", cCtx.Args().First())
							return nil
						},
					},
				},
			},
		},
	}

	_ = app.Run(os.Args)
}
