package repo

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"os"
	"password-tool/settings"
)

func List() error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"名称", "是否默认"})

	files, err := ioutil.ReadDir(settings.Settings.RepoPath)
	if err != nil {
		return err
	}
	for _, f := range files {
		table.Append([]string{f.Name(), fmt.Sprintf("%t", f.Name() == settings.Settings.DefaultRepo)})
	}
	table.Render() // Send output
	return nil
}
