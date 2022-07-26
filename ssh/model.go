package ssh

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	Ip       string `gorm:"type:varchar(255);index"`
	Name     string `gorm:"type:varchar(255);unique_index"`
	Password string `gorm:"type:varchar(255);index"`
	Port     int    `gorm:"type:int;index"`
	User     string `gorm:"type:varchar(255);index"`
}
