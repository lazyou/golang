package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Method:   GET
	// Resource: http://localhost:9000
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:9000/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:9000/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	// 参数获取: 如果 id 传入非数字, 会提示 Not Found
	app.Get("/get/{id:uint64}", func(ctx iris.Context) {
		id := ctx.Params().GetUint64Default("id", 0)
		ctx.JSON(iris.Map{"message": id})
	})

	app.Get("/user/{name:string}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef("Hello %s", name)
	})

	// Dependency Injection

	// 1. Path Parameters - Built-in Dependencies
	helloHandler := hero.Handler(hello)
	app.Get("/di//{to:string}", helloHandler)

	// 2. Services - Static Dependencies
	hero.Register(
		&myTestService{
			prefix: "Service: Hello",
		},
	)

	helloServiceHandler := hero.Handler(helloService)
	app.Get("/di/service/{to:string}", helloServiceHandler)

	// 3. Per-Request - Dynamic Dependencies
	hero.Register(func(ctx iris.Context) (form LoginForm) {
		// it binds the "form" with a
		// x-www-form-urlencoded form data and returns it.
		ctx.ReadForm(&form)
		return
	})

	loginHandler := hero.Handler(login)
	app.Post("/di/login", loginHandler)

	// http://localhost:9000
	// http://localhost:9000/ping
	// http://localhost:9000/hello
	app.Run(iris.Addr(":9000"), iris.WithoutServerError(iris.ErrServerClosed))
}

// for Dependency Injection 1
func hello(s string) string {
	return "hello " + s
}

// for Dependency Injection 2
type Service interface {
	SayHello(to string) string
}

type myTestService struct {
	prefix string
}

func (s *myTestService) SayHello(to string) string {
	return s.prefix + " " + to
}

func helloService(to string, service Service) string {
	return service.SayHello(to)
}

// for Dependency Injection 3
type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func login(form LoginForm) string {
	return "Hello " + form.Username
}
