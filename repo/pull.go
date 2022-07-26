package repo

import (
	"fmt"
	"io/ioutil"
	"os"
	"password-tool/settings"
	"password-tool/utils"
	"path"
)

func Pull(repoName string, recursive bool) error {

	if recursive {
		files, err := ioutil.ReadDir(settings.Settings.RepoPath)
		if err != nil {
			return err
		}
		for _, f := range files {
			if f.Name() == "default" {
				continue
			}
			err := gitPull(f.Name())
			if err != nil {
				return err
			}
		}
		return nil
	}

	if repoName == "" {
		repoName = settings.Settings.DefaultRepo
	}
	if repoName == "default" {
		return fmt.Errorf("只有远程仓库可以执行 pull 操作")
	}
	repoDir := path.Join(settings.Settings.RepoPath, repoName)
	if _, err := os.Stat(repoDir); err != nil {
		return fmt.Errorf("仓库不存在")
	}
	return gitPull(repoName)
}

func gitPull(repoName string) error {
	_, o, err := utils.RunCommand("cd " + path.Join(settings.Settings.RepoPath, repoName) + " && git pull")
	if err != nil {
		return fmt.Errorf("在仓库 %s 上执行 git pull 失败，请手动检查后，在执行。Error：%s", path.Join(settings.Settings.RepoPath, repoName), o)
	}
	return nil
}
