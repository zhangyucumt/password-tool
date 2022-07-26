package ssh

import (
	"password-tool/settings"
	"password-tool/utils"
	"path"
)

func Delete(name string) error {
	_, _, err := utils.RunCommand("rm -rf " + path.Join(settings.Settings.DefaultRepoPath, "ssh", name+".yaml"))
	return err
}
