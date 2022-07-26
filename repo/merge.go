package repo

import (
	"fmt"
	"os"
	"password-tool/settings"
	"password-tool/utils"
	"path"
)

func Merge(repo1, repo2 string) error {
	//if settings.Settings.DefaultRepo == repoName {
	//	return fmt.Errorf("无法删除默认仓库。请先将默认仓库设置为其他仓库后，在执行删除操作")
	//}
	//if repoName == "default" {
	//	return fmt.Errorf("为保证程序正常运行，已禁止删除本地仓库")
	//}

	if _, err := os.Stat(path.Join(settings.Settings.RepoPath, repo1)); err != nil {
		return fmt.Errorf("仓库不存在: %s。无法进行合并", repo1)
	}
	if _, err := os.Stat(path.Join(settings.Settings.RepoPath, repo2)); err != nil {
		return fmt.Errorf("仓库不存在: %s。无法进行合并", repo2)
	}

	if _, err := os.Stat(path.Join(settings.Settings.RepoPath, repo2, "ssh")); err != nil {
		return nil
	} else {
		if _, err := os.Stat(path.Join(settings.Settings.RepoPath, repo1, "ssh")); err != nil {
			_, _, err := utils.RunCommand(fmt.Sprintf("cp -a %s %s/", path.Join(settings.Settings.RepoPath, repo2, "ssh"), path.Join(settings.Settings.RepoPath, repo1)))
			return err
		} else {
			_, _, err := utils.RunCommand(fmt.Sprintf("cp -a %s/* %s/", path.Join(settings.Settings.RepoPath, repo1, "ssh"), path.Join(settings.Settings.RepoPath, repo2, "ssh")))
			return err
		}
	}
}
