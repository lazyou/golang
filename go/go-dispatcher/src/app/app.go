package app

import (
	"fmt"
	"net/http"
	"reflect"
)

type application struct {
	routes map[string]interface{} // 路由名与控制器映射
}

func New() *application {
	return &application{
		routes: make(map[string]interface{}),
	}
}

func (p *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 根据 url 里面的控制器名和方法名找到的对应方法
	controllerName := r.URL.Query().Get("c")
	actionName := r.URL.Query().Get("a")

	// a 或 c 参数问题
	if controllerName == "" || actionName == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// 找不到对应控制器
	route, ok := p.routes[controllerName]
	if !ok {
		http.Error(w, "Controller Not Found", http.StatusNotFound)
		return
	}

	// 打印各种格式测试
	fmt.Printf("%v \n", route)
	fmt.Printf("%T \n", route)
	fmt.Printf("%#v \n", route)
	fmt.Printf("%+v \n", route)

	// 利用反射机智
	ele := reflect.ValueOf(route).Elem()
	ele.FieldByName("Request").Set(reflect.ValueOf(r))
	ele.FieldByName("Response").Set(reflect.ValueOf(w))
	ele.MethodByName(actionName).Call([]reflect.Value{})
}

// 打印路由信息
func (p *application) printRoutes() {
	for route, controller := range p.routes {
		ele := reflect.ValueOf(controller).Type().String()
		fmt.Printf("%s %s\n", route, ele)
	}
}

// 添加 路由控制器名 与 本地控制器之间的映射
func (p *application) Get(route string, controller interface{}) {
	p.routes[route] = controller
}

func (p *application) Run(addr string) error {
	p.printRoutes()
	fmt.Printf("listen on %s\n", addr)
	return http.ListenAndServe(addr, p)
}
