package repo

import (
	"fmt"
	"io/ioutil"
	"os"
	"password-tool/settings"
	"password-tool/utils"
	"path"
)

func Push(repoName string) error {

	if repoName == "" {
		files, err := ioutil.ReadDir(settings.Settings.RepoPath)
		if err != nil {
			return err
		}
		for _, f := range files {
			if f.Name() == "default" {
				continue
			}
			err := gitPush(f.Name())
			if err != nil {
				return err
			}
		}
	}

	repoDir := path.Join(settings.Settings.RepoPath, repoName)
	if _, err := os.Stat(repoDir); err != nil {
		return fmt.Errorf("仓库不存在")
	}
	return gitPush(repoName)
}

func gitPush(repoName string) error {
	_, _, err := utils.RunCommand("cd " + path.Join(settings.Settings.RepoPath, repoName) + " && git add . && git commit -m 'push by command' && git push")
	if err != nil {
		return fmt.Errorf("在仓库 %s 上执行 git push 失败，请手动检查后，在执行", path.Join(settings.Settings.RepoPath, repoName))
	}
	return nil
}
