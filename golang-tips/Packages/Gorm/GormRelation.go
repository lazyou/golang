package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lazyou/GolangTips/Packages/Gorm/Models"
)

var gormDB *gorm.DB

// 模型关联
func main() {
	Models.GromConnection()
	gormDB = Models.DB

	//hasOne()
	//hasMany()
	manyToMany()
}

func hasOne() {
	// 先查询主模型
	user := Models.User{Id: 1}
	gormDB.First(&user)

	// 再查询关联模型
	role := Models.Role{}
	gormDB.Model(&user).Related(&role, "Role")

	fmt.Printf("%+v\n", user)
	fmt.Printf("%+v\n", role)
}

func hasMany() {
	// 先查询主模型
	role := &Models.Role{Id: 1}
	gormDB.First(role)

	users := new([]Models.User)
	gormDB.Model(role).Related(users)

	fmt.Printf("%+v\n", role)
	fmt.Printf("%+v\n", users)
}

func manyToMany() {
	// 先查询主模型
	user := &Models.User{Id: 1}
	gormDB.First(user)

	// 再查询关联模型
	projects := new([]Models.Project)
	// TODO: 这里必须指定 Related 的模型, 因为 User 已经关联太多模型了
	gormDB.Model(user).Related(projects, "Projects")

	fmt.Printf("%+v\n", user)
	fmt.Printf("%+v\n", projects)
}
