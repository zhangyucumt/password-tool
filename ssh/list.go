package ssh

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func List(keyword string) error {
	var configs []Model
	if keyword != "" {
		db.Where("name like ? OR host = ?", "%"+keyword+"%", "%"+keyword+"%").Find(&configs)
	} else {
		db.Find(&configs)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(200)
	table.SetHeader([]string{"名称", "主机", "端口", "用户", "密码", "命令"})

	for _, c := range configs {
		command := fmt.Sprintf("ssh %s@%s -p %d", c.User, c.Host, c.Port)
		if c.Password != "" {
			command = fmt.Sprintf("sshpass -p %s ssh %s@%s -p %d", c.Password, c.User, c.Host, c.Port)
		}
		table.Append([]string{
			c.Name,
			c.Host,
			fmt.Sprintf("%d", c.Port),
			c.User,
			c.Password,
			command,
		})
	}
	table.Render() // Send output
	return nil

}
