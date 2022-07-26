package repo

import (
	"fmt"
	"os"
	"password-tool/settings"
	"password-tool/utils"
	"path"
)

func Add(repoName string, gitUrl string, isDefault bool) error {
	repoDir := path.Join(settings.Settings.RepoPath, repoName)
	if _, err := os.Stat(repoDir); err == nil {
		return fmt.Errorf("仓库已存在: %s，请先移除之后再执行添加", repoDir)
	}
	_, _, err := utils.RunCommand("cd " + settings.Settings.RepoPath + " && git clone " + gitUrl + " " + repoName)
	if err != nil {
		return err
	}

	if isDefault {
		err := settings.Settings.Update("default_repo", repoName)
		if err != nil {
			return err
		}
	}

	return nil
}
