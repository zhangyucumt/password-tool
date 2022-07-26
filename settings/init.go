package settings

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

type Config struct {
	DefaultRepo string `yaml:"default_repo" json:"default_repo"`
}

type Struct struct {
	HomePath        string
	AppPath         string
	RepoPath        string
	ConfigFilePath  string
	DefaultRepoPath string
	Config
}

var Settings *Struct

func init() {
	homePath := os.Getenv("HOME")
	appPath := path.Join(homePath, ".password-tool")
	configFilePath := path.Join(homePath, ".password-tool/config.yaml")

	if _, err := os.Stat(configFilePath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(appPath, 0755)
		if err != nil {
			panic(fmt.Errorf("无法创建应用文件夹: " + appPath))
		}
		err = initConfigFile(appPath)
		if err != nil {
			panic(err)
		}
	}

	yamlFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	Settings = &Struct{
		HomePath:        homePath,
		AppPath:         appPath,
		RepoPath:        path.Join(appPath, "repos"),
		ConfigFilePath:  configFilePath,
		DefaultRepoPath: path.Join(appPath, "repos", config.DefaultRepo),
		Config:          config,
	}
}
