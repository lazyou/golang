package Models

import (
	"github.com/jinzhu/gorm"
)

// 标签模型

// Tag default table name is `tags`
type Tag struct {
	gorm.Model
	Id            int `gorm:"primary_key"`
	Name          string
	CreatedUserId int
	// 以下三个字段已在 gorm.Model 嵌入
	//CreatedAt     time.Time
	//UpdatedAt     time.Time
	//DeletedAt     *time.Time
}

// TableName 设置模型表名
func (tag Tag) TableName() string {
	return "tags"
}
