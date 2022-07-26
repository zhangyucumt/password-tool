package settings

import (
	"gopkg.in/yaml.v2"
	"os"
	"password-tool/utils"
	"path"
)

func initConfigFile(appPath string) error {
	config := Config{
		DefaultRepo: "default",
	}
	yamlData, _ := yaml.Marshal(config)

	f, err := os.Create(path.Join(appPath, "config.yaml"))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(yamlData)

	if err != nil {
		return err
	}

	err = os.MkdirAll(path.Join(appPath, "repos/default"), 0755)
	if err != nil {
		return err
	}

	_, _, err = utils.RunCommand("cd " + path.Join(appPath, "repos/default") + " && git init")
	if err != nil {
		return err
	}
	return nil
}
