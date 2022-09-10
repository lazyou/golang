package Models

import (
	"github.com/jinzhu/gorm"
)

// Role 角色模型
type Role struct {
	gorm.Model
	Id    int `gorm:"primary_key"`
	Name  string
	Type  int
	Users []User `gorm:"foreignkey:RoleId"`
}

func (role Role) TableName() string {
	return "roles"
}
