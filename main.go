package main

import (
	"fmt"
	"os"
	"password-tool/repo"
	"password-tool/ssh"

	"github.com/urfave/cli/v2"
)

func main() {

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "print-version",
		Aliases: []string{"V", "version", "v"},
		Usage:   "版本信息",
	}

	app := &cli.App{
		Version:              "v3.1",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "repo",
				Aliases: []string{"r"},
				Usage:   "数据存储的代码仓库",
				Subcommands: []*cli.Command{
					{
						Name:    "add",
						Aliases: []string{"a"},
						Usage:   "新增数据存储的代码仓库",
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
								//DefaultText: "git@github.com:zhangyucumt/password-tool.git",
							},
							&cli.StringFlag{
								Name:  "name",
								Usage: "代码仓库的名称标识。",
								//Required:    true,
								//DefaultText: "repo1",
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
								return cli.Exit("新增代码仓库失败: "+err.Error(), 1)
							} else {
								fmt.Println("已添加代码仓库成功")
								return nil
							}
						},
					},
					{
						Name:    "remove",
						Aliases: []string{"rm", "del", "delete"},
						Usage:   "移除数据存储的代码仓库",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "name",
								Usage: "代码仓库的名称标识。",
								//DefaultText: "repo1",
							},
						},
						Action: func(cCtx *cli.Context) error {
							name := cCtx.String("name")
							if name == "" {
								name = cCtx.Args().First()
							}
							if name == "" {
								fmt.Println("请指定移除仓库名称")
								return cli.Exit("缺少参数", 2)
							}
							err := repo.Delete(name)
							if err != nil {
								return cli.Exit("移除代码仓库失败: "+err.Error(), 1)
							} else {
								fmt.Println("移除代码仓库成功")
								return nil
							}
						},
					},
					{
						Name:    "list",
						Usage:   "数据存储的代码仓库列表",
						Aliases: []string{"ls", "l"},
						Action: func(cCtx *cli.Context) error {
							err := repo.List()
							if err != nil {
								return cli.Exit(err.Error(), 2)
							} else {
								return nil
							}
						},
					},
					{
						Name:    "set-default",
						Aliases: []string{"sd", "default", "setD"},
						Usage:   "设置某个代码仓库为默认的数据存储仓库",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "name",
								Usage: "代码仓库的名称标识。",
								//Required:    true,
								//DefaultText: "repo1",
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
							err := repo.SetDefault(name)
							if err != nil {
								return cli.Exit(err.Error(), 2)
							} else {
								fmt.Println("设置默认仓库成功")
								return nil
							}
						},
					},
					{
						Name:  "pull",
						Usage: "将代码仓库数据同步至本地存储。默认同步的仓库是默认仓库",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "name",
								Usage: "需要同步的代码仓库的名称标识。",
								//DefaultText: "repo1",
							},
							&cli.BoolFlag{
								Name:    "recursive",
								Aliases: []string{"r"},
								Usage:   "是否递归同步所有仓库。",
								Value:   false,
							},
						},
						Action: func(cCtx *cli.Context) error {
							//fmt.Println("removed task template: ", cCtx.Args().First())
							err := repo.Pull(cCtx.String("name"), cCtx.Bool("recursive"))
							if err != nil {
								return cli.Exit(err.Error(), 2)
							} else {
								fmt.Println("将代码仓库数据同步至本地存储成功")
								return nil
							}
						},
					},
					{
						Name:  "push",
						Usage: "将本地存储的数据同步至代码仓库。默认同步的仓库是默认仓库",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "name",
								Usage: "需要同步的代码仓库的名称标识。",
								//DefaultText: "repo1",
							},
							&cli.BoolFlag{
								Name:    "recursive",
								Aliases: []string{"r"},
								Usage:   "是否递归同步所有仓库。",
								Value:   false,
							},
						},
						Action: func(cCtx *cli.Context) error {
							//fmt.Println("removed task template: ", cCtx.Args().First())
							err := repo.Push(cCtx.String("name"), cCtx.Bool("recursive"))
							if err != nil {
								return cli.Exit("代码运行异常: "+err.Error(), 2)
							} else {
								fmt.Println("将本地存储的数据同步至代码仓库成功")
								return nil
							}
						},
					},
					{
						Name: "merge",
						//Aliases: []string{"rm", "del", "delete"},
						Usage: "将其他的代码仓库合并至一个代码仓库中。",
						Action: func(cCtx *cli.Context) error {
							repo1 := cCtx.Args().First()
							repo2 := cCtx.Args().Get(1)
							err := repo.Merge(repo1, repo2)
							if err != nil {
								return cli.Exit("合并代码仓库失败: "+err.Error(), 1)
							} else {
								fmt.Println("合并代码仓库成功")
								return nil
							}
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
				Subcommands: []*cli.Command{
					{
						Name:    "add",
						Aliases: []string{"a"},
						Usage:   "添加一个ssh配置",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "name",
								Usage:    "配置名称",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "host",
								Usage:    "ssh主机",
								Required: true,
							},
							&cli.IntFlag{
								Name:  "port",
								Usage: "ssh端口",
								Value: 22,
							},
							&cli.StringFlag{
								Name:  "user",
								Usage: "ssh用户",
								Value: "root",
							},
							&cli.StringFlag{
								Name:  "password",
								Usage: "ssh密码",
								Value: "",
							},
						},
						Action: func(cCtx *cli.Context) error {
							err := ssh.Add(cCtx.String("name"), cCtx.String("host"), cCtx.Int("port"), cCtx.String("user"), cCtx.String("password"))
							if err != nil {
								return cli.Exit(err.Error(), 2)
							} else {
								fmt.Println("添加ssh配置成功")
								return nil
							}
						},
					},
					{
						Name:    "list",
						Aliases: []string{"ls", "l"},
						Usage:   "列出ssh配置",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "search",
								Usage: "搜索ssh配置关键字",
							},
						},
						Action: func(cCtx *cli.Context) error {
							keyword := cCtx.String("search")
							if keyword == "" {
								keyword = cCtx.Args().First()
							}
							err := ssh.List(keyword)
							if err != nil {
								return cli.Exit(err.Error(), 2)
							} else {
								return nil
							}
						},
					},
					{
						Name:    "update",
						Aliases: []string{"edit", "e", "u"},
						Usage:   "更新ssh配置",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "name",
								Usage: "配置名称",
								Value: "",
							},
							&cli.StringFlag{
								Name:  "newName",
								Usage: "新的配置名称",
								Value: "",
							},
							&cli.StringFlag{
								Name:  "host",
								Usage: "ssh主机",
								Value: "",
							},
							&cli.IntFlag{
								Name:  "port",
								Usage: "ssh端口",
								Value: 0,
							},
							&cli.StringFlag{
								Name:  "user",
								Usage: "ssh用户",
								Value: "",
							},
							&cli.StringFlag{
								Name:  "password",
								Usage: "ssh密码",
								Value: "",
							},
						},
						Action: func(cCtx *cli.Context) error {
							name := cCtx.String("name")
							if name == "" {
								name = cCtx.Args().First()
							}
							err := ssh.Update(name, cCtx.String("host"), cCtx.Int("port"), cCtx.String("user"), cCtx.String("password"), cCtx.String("newName"))
							if err != nil {
								return cli.Exit(err.Error(), 2)
							} else {
								fmt.Println("更新ssh配置成功")
								return nil
							}
						},
					},
					{
						Name:    "remove",
						Aliases: []string{"rm", "delete", "d", "del"},
						Usage:   "删除ssh配置",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "name",
								Usage: "配置名称",
								Value: "",
							},
						},
						Action: func(cCtx *cli.Context) error {
							name := cCtx.String("name")
							if name == "" {
								name = cCtx.Args().First()
							}
							err := ssh.Delete(name)
							if err != nil {
								return cli.Exit(err.Error(), 2)
							} else {
								fmt.Println("删除ssh配置成功")
								return nil
							}
						},
					},
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "添加一个ssh配置",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "配置名称",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "host",
						Usage:    "ssh主机",
						Required: true,
					},
					&cli.IntFlag{
						Name:  "port",
						Usage: "ssh端口",
						Value: 22,
					},
					&cli.StringFlag{
						Name:  "user",
						Usage: "ssh用户",
						Value: "root",
					},
					&cli.StringFlag{
						Name:  "password",
						Usage: "ssh密码",
						//Required: true,
						Value: "",
					},
				},
				Action: func(cCtx *cli.Context) error {
					err := ssh.Add(cCtx.String("name"), cCtx.String("host"), cCtx.Int("port"), cCtx.String("user"), cCtx.String("password"))
					if err != nil {
						return cli.Exit(err.Error(), 2)
					} else {
						fmt.Println("添加ssh配置成功")
						return nil
					}
				},
			},
			{
				Name:    "list",
				Aliases: []string{"ls", "l"},
				Usage:   "列出ssh配置",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "search",
						Usage: "搜索ssh配置关键字",
					},
				},
				Action: func(cCtx *cli.Context) error {
					keyword := cCtx.String("search")
					if keyword == "" {
						keyword = cCtx.Args().First()
					}
					err := ssh.List(keyword)
					if err != nil {
						return cli.Exit(err.Error(), 2)
					} else {
						return nil
					}
				},
			},
			{
				Name:    "update",
				Aliases: []string{"edit", "e", "u"},
				Usage:   "更新ssh配置",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "name",
						Usage: "配置名称",
						Value: "",
					},
					&cli.StringFlag{
						Name:  "newName",
						Usage: "新的配置名称",
						Value: "",
					},
					&cli.StringFlag{
						Name:  "host",
						Usage: "ssh主机",
						Value: "",
					},
					&cli.IntFlag{
						Name:  "port",
						Usage: "ssh端口",
						Value: 0,
					},
					&cli.StringFlag{
						Name:  "user",
						Usage: "ssh用户",
						Value: "",
					},
					&cli.StringFlag{
						Name:  "password",
						Usage: "ssh密码",
						Value: "",
					},
				},
				Action: func(cCtx *cli.Context) error {
					name := cCtx.String("name")
					if name == "" {
						name = cCtx.Args().First()
					}
					err := ssh.Update(name, cCtx.String("host"), cCtx.Int("port"), cCtx.String("user"), cCtx.String("password"), cCtx.String("newName"))
					if err != nil {
						return cli.Exit(err.Error(), 2)
					} else {
						fmt.Println("更新ssh配置成功")
						return nil
					}
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"rm", "delete", "d", "del"},
				Usage:   "删除ssh配置",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "name",
						Usage: "配置名称",
						Value: "",
					},
				},
				Action: func(cCtx *cli.Context) error {
					name := cCtx.String("name")
					if name == "" {
						name = cCtx.Args().First()
					}
					err := ssh.Delete(name)
					if err != nil {
						return cli.Exit(err.Error(), 2)
					} else {
						fmt.Println("删除ssh配置成功")
						return nil
					}
				},
			},
		},
	}

	_ = app.Run(os.Args)
}
