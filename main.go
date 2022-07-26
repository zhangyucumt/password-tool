package main

import (
	"fmt"
	"os"
	"password-tool/repo"

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
								Name:  "url",
								Usage: "代码仓库的地址。默认应具有代码仓库的读写权限",
								//Required:    true,
								DefaultText: "git@github.com:zhangyucumt/password-tool.git",
							},
							&cli.StringFlag{
								Name:  "name",
								Usage: "代码仓库的名称标识。",
								//Required:    true,
								DefaultText: "repo1",
							},
						},
						Action: func(cCtx *cli.Context) error {
							name := cCtx.String("name")
							url := cCtx.String("url")
							if name == "" {
								name = cCtx.Args().First()
							}
							if url == "" {
								url = cCtx.Args().Get(1)
							}
							if name == "" || url == "" {
								fmt.Println("请指定名称和地址")
								return fmt.Errorf("缺少参数")
							}
							defaultRepo := cCtx.Bool("default")
							err := repo.Add(name, url, defaultRepo)
							if err != nil {
								fmt.Println("新增代码仓库失败:", err)
							} else {
								fmt.Println("已成功代码仓库成功")
							}
							return err
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
							name := cCtx.String("name")
							if name == "" {
								name = cCtx.Args().First()
							}
							if name == "" {
								fmt.Println("请指定移除仓库名称")
								return fmt.Errorf("缺少参数")
							}
							return repo.Delete(name)
						},
					},
					{
						Name:  "list",
						Usage: "数据存储的代码仓库列表",
						Action: func(cCtx *cli.Context) error {
							return repo.List()
						},
					},
					{
						Name:  "set-default",
						Usage: "设置某个代码仓库为默认的数据存储仓库",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "name",
								Usage: "代码仓库的名称标识。",
								//Required:    true,
								DefaultText: "repo1",
							},
						},
						Action: func(cCtx *cli.Context) error {
							name := cCtx.String("name")
							if name == "" {
								name = cCtx.Args().First()
							}
							if name == "" {
								fmt.Println("请指定默认仓库名称")
								return fmt.Errorf("缺少参数")
							}
							return repo.SetDefault(name)
						},
					},
					{
						Name:  "pull",
						Usage: "将代码仓库数据同步至本地存储。",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "name",
								Usage:       "需要同步的代码仓库的名称标识。",
								DefaultText: "repo1",
							},
						},
						Action: func(cCtx *cli.Context) error {
							//fmt.Println("removed task template: ", cCtx.Args().First())
							return repo.Pull(cCtx.String("name"))
						},
					},
					{
						Name:  "push",
						Usage: "将本地存储的数据同步至代码仓库。默认同步的仓库是默认仓库",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "name",
								Usage:       "需要同步的代码仓库的名称标识。",
								DefaultText: "repo1",
							},
						},
						Action: func(cCtx *cli.Context) error {
							//fmt.Println("removed task template: ", cCtx.Args().First())
							return repo.Push(cCtx.String("name"))
						},
					},
				},
				//Action: func(cCtx *cli.Context) error {
				//	fmt.Println("removed task template: ", cCtx.Args().First())
				//	return nil
				//},
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
