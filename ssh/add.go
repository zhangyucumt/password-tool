package ssh

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"password-tool/settings"
	"path"
)

func Add(name, ip string, port int, user string, password string) error {
	result := db.First(&Model{}, "name = ?", name)
	if result.RowsAffected > 0 {
		return fmt.Errorf("名称为: %s 的配置已存在", name)
	}

	s := Struct{
		Name:     name,
		Ip:       ip,
		Password: password,
		Port:     port,
		User:     user,
	}

	yamlData, err := yaml.Marshal(s)
	if err != nil {
		return err
	}
	err = os.MkdirAll(path.Join(settings.Settings.DefaultRepoPath, "ssh"), 0755)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path.Join(settings.Settings.DefaultRepoPath, "ssh", name+".yaml"), yamlData, 0644)

}
