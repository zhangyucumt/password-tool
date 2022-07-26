package repo

import (
	"fmt"
	"os"
	"password-tool/settings"
	"path"
)

func SetDefault(repoName string) error {

	repoDir := path.Join(settings.Settings.RepoPath, repoName)
	if _, err := os.Stat(repoDir); err != nil {
		return fmt.Errorf("无法设置不存在的仓库为默认仓库")
	}
	err := settings.Settings.Update("default_repo", repoName)
	return err
}
