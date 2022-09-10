package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lazyou/GolangTips/Packages/Gorm/Models"
)

var db *gorm.DB

// CURD 案例
func main() {
	Models.GromConnection()
	db = Models.DB

	//createTag()
	queryTag()
	//queryToStruct()
	//updateTag()
}

// createTag https://gorm.io/docs/create.html
func createTag() {
	fmt.Println("[增:]")
	defer fmt.Println("")

	tag := Models.Tag{
		Name:          "golang",
		CreatedUserId: 0,
	}

	// 插入一条数据
	// Create Record: 需要下面两部
	fmt.Println(db.NewRecord(tag)) // 返回 bool
	fmt.Println(db.Create(&tag))   // 返回 *gorm.DB
	//fmt.Println(db.NewRecord(tag)) // false: 因为上面已经执行了 db.Create
	//fmt.Println(db.Create(&tag)) // TODO： 这里会报错哦

	// TODO: 如何插入多条
}

func queryToStruct() {
	fmt.Println("[查寻结果放入 Struct:]")
	defer fmt.Println("")

	type NewTag struct {
		Id   int
		Name string
	}

	// 数组结构体
	var tagSlice = new([]NewTag)
	db.Model(&Models.Tag{}).Where("id > ?", 0).Scan(tagSlice)
	fmt.Printf("数组结构体: %v\n", tagSlice)

	// 数组结构体指针
	var tagPonterSlice = new([]*NewTag)
	db.Model(&Models.Tag{}).Where("id > ?", 0).Scan(tagPonterSlice)
	fmt.Printf("数组结构体指针: %v\n", tagPonterSlice)
}

// queryTag https://gorm.io/docs/query.html
func queryTag() {
	fmt.Println("[查:]")
	defer fmt.Println("")

	tag := Models.Tag{
		Id: 2,
	}

	// 查询一条记录.
	db.First(&tag) // 查询成功后会自动将数据赋值在 tag 上, 不需要声明新的变量获取. TODO: 有坑, 若 Tag 的 Id 赋值, 并没有加入到查询条件中. First() 必须使用第二个参数指定主键, 否则就是主键的默认值(而非传入部分).
	//db.First(&tag, 2) // 指定主键
	fmt.Println(tag)
	//return

	// 查询多条(全部)
	var tags = new([]Models.Tag)
	//var tags = new([]interface{}) // 如果改成通用接口显然查不到数据
	db.Where("created_user_id = ?", 1).Find(tags) // 估计是通过反射推导出所查询的模型(表)
	// TODO: 坑 -- 如果查询条件拼写错误(查询不存在的字段), 不会报错 Unknown column 'user_id' in 'where clause'
	fmt.Println(tags)

	// Struct & Map 内写查询条件
	db.Where(&Models.Tag{Name: "golang2"}).Find(&tags)
	fmt.Println(tags)

	db.Select("name").Find(&tags) // select 指定字段
	fmt.Println(tags)

	// More: FirstOrInit, FirstOrCreate ... Group & Having, Joins, Pluck, Scan
}

// updateTag https://gorm.io/docs/update.html
func updateTag() {
	fmt.Println("[改:]")
	defer fmt.Println("")

	// 改: 先查后改
	tag := Models.Tag{Id: 1}
	db.First(&tag)

	tag.Name = "updateTagName"
	db.Save(&tag)

	// 批量更新
	db.Model(Models.Tag{}).
		Where("id IN (?)", []int{1, 2, 3, 4, 5}).
		Updates(Models.Tag{Name: "BatchUpdateTagName"})
}
