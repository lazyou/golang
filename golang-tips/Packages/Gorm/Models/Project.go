package Models

import (
	"github.com/jinzhu/gorm"
)

// Project 项目模型
type Project struct {
	gorm.Model
	Id    int `gorm:"primary_key"`
	Name  string
	Users []User `gorm:"many2many:project_users;"`
}

func (role Project) TableName() string {
	return "projects"
}
