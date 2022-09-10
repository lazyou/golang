package main

import (
	"app"
	"controllers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	application := app.New()
	application.Get("site", &controllers.SiteController{})
	application.Get("test", &controllers.TestController{})
	application.Run(":8080")
}
