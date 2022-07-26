package repo

import (
	"fmt"
	"io/ioutil"
	"os"
	"password-tool/settings"
	"password-tool/utils"
	"path"
)

func Push(repoName string, recursive bool) error {

	if recursive {
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
		return nil
	}

	if repoName == "" {
		repoName = settings.Settings.DefaultRepo
	}
	if repoName == "default" {
		return fmt.Errorf("只有远程仓库可以执行 push 操作")
	}

	repoDir := path.Join(settings.Settings.RepoPath, repoName)
	if _, err := os.Stat(repoDir); err != nil {
		return fmt.Errorf("仓库不存在")
	}
	return gitPush(repoName)
}

func gitPush(repoName string) error {
	o, _, err := utils.RunCommand("cd " + path.Join(settings.Settings.RepoPath, repoName) + " && git add . && git commit -m 'push by command' && git push")
	if err != nil {
		return fmt.Errorf("在仓库 %s 上执行 git push 失败，请手动检查后，在执行。Error：%s", path.Join(settings.Settings.RepoPath, repoName), o)
	}
	return nil
}
