package Models

import (
	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Id       int `gorm:"primary_key"`
	Name     string
	Email    string
	RoleId   int
	Role     Role      `gorm:"foreignkey:RoleId"` // use RoleId as foreign key
	Projects []Project `gorm:"many2many:project_users;"`
}

func (user User) TableName() string {
	return "users"
}
