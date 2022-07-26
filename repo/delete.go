package repo

import (
	"fmt"
	"password-tool/settings"
	"password-tool/utils"
)

func Delete(repoName string) error {
	if settings.Settings.DefaultRepo == repoName {
		return fmt.Errorf("无法删除默认仓库。请先将默认仓库设置为其他仓库后，在执行删除操作")
	}
	if repoName == "default" {
		return fmt.Errorf("为保证程序正常运行，已禁止删除本地仓库")
	}

	//if _, err := os.Stat(path.Join(settings.Settings.RepoPath, repoName)); err != nil {
	//	return fmt.Errorf("仓库不存在: %s。无法删除", repoName)
	//}

	_, _, err := utils.RunCommand("cd " + settings.Settings.RepoPath + " && rm -rf " + repoName)
	return err
}
