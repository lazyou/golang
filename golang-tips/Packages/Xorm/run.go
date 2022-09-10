package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/core"
	"xorm.io/xorm"
)

var DB *xorm.Engine

func main() {
	var err error
	var dbLink = "homestead:secret@tcp(localhost:33060)/curd?charset=utf8mb4"

	DB, err = xorm.NewEngine("mysql", dbLink)
	if err != nil {
		fmt.Println(err)
	}

	DB.ShowSQL(true)
	DB.Logger().SetLevel(core.LOG_DEBUG)

	DB.SetMapper(core.SnakeMapper{})
	//DB.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, ""))
}
