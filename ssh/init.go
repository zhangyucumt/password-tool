package ssh

import (
	"gopkg.in/yaml.v2"
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm/logger"
	"io/ioutil"
	"password-tool/settings"
	"path"
	"strings"

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

var db *gorm.DB

type Struct struct {
	Name     string `yaml:"name" json:"name"`
	Ip       string `yaml:"ip" json:"ip"`
	Password string `yaml:"password" json:"password"`
	Port     int    `yaml:"port" json:"port"`
	User     string `yaml:"user" json:"user"`
}

func init() {
	// github.com/mattn/go-sqlite3
	db, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true, Logger: logger.Default.LogMode(logger.Silent)})
	err := db.AutoMigrate(&Model{})
	//err := db.Migrator().CreateTable(&Model{})
	if err != nil {
		panic(err)
	}

	repoPath := path.Join(settings.Settings.RepoPath, settings.Settings.DefaultRepo, "ssh")

	files, err := ioutil.ReadDir(repoPath)
	if err != nil {
		return
	}
	for _, f := range files {
		var s *Struct
		if !strings.HasSuffix(f.Name(), ".yaml") {
			continue
		}

		yamlFile, err := ioutil.ReadFile(path.Join(repoPath, f.Name()))
		if err != nil {
			continue
		}
		err = yaml.Unmarshal(yamlFile, &s)
		if err != nil {
			continue
		}
		db.Create(&Model{
			Name:     s.Name,
			Ip:       s.Ip,
			Password: s.Password,
			Port:     s.Port,
			User:     s.User,
		})
	}

}
