package ssh

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"password-tool/settings"
	"path"
)

func Update(name, ip string, port int, user string, password string) error {
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
