package Models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var dbErr error

func GromConnection() {
	DB, dbErr = gorm.Open("mysql", "homestead:secret@tcp(localhost:33060)/curd?charset=utf8mb4&parseTime=True&loc=Local")
	//defer DB.Close()

	if dbErr != nil {
		fmt.Printf("数据库连接失败: %v", dbErr)
	}

	// 表前缀设置
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "" + defaultTableName
		//return "prefix_" + defaultTableName
	}

	DB.LogMode(true)

	fmt.Println("gorm start!")
}
