package ssh

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"password-tool/settings"
	"password-tool/utils"
	"path"
)

func Update(name, host string, port int, user string, password string, newName string) error {
	if _, err := os.Stat(path.Join(settings.Settings.DefaultRepoPath, "ssh", name+".yaml")); err != nil {
		return fmt.Errorf("配置不存在: %s", name)
	}

	var s Struct
	yamlData, err := ioutil.ReadFile(path.Join(settings.Settings.DefaultRepoPath, "ssh", name+".yaml"))
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlData, &s)
	if err != nil {
		return err
	}
	if host != "" {
		s.Host = host
	}
	if port != 0 {
		s.Port = port
	}
	if user != "" {
		s.User = user
	}
	if password != "" {
		s.Password = password
	}
	if newName != "" {
		s.Name = newName
	}

	yamlData, err = yaml.Marshal(s)
	if err != nil {
		return err
	}

	if s.Name != name {
		_, _, err := utils.RunCommand("rm -rf " + path.Join(settings.Settings.DefaultRepoPath, "ssh", name+".yaml"))
		if err != nil {
			return err
		}
	}

	return ioutil.WriteFile(path.Join(settings.Settings.DefaultRepoPath, "ssh", s.Name+".yaml"), yamlData, 0644)

}
