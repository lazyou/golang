### 介绍
* Golang 官方编程指南

### 看点
* 注意 .go 文件里的注释  `//`, 每一段注释都是来自教程


### 教程
* basics: https://tour.go-zh.org || https://tour.golang.org

* flowcontrol: https://tour.go-zh.org/flowcontrol/1

* moretypes: https://tour.go-zh.org/moretypes/1

* methods: https://tour.go-zh.org/methods/1

* concurrency: https://tour.go-zh.org/concurrency/1

* https://tour.golang.org/list || https://tour.go-zh.org/list


### api 文档
* https://go-zh.org/pkg || https://golang.org/pkg


### 常用命令
* `go run xx.go`


### go help
```
Go is a tool for managing Go source code.

Usage:

        go command [arguments]

The commands are:

        build       compile packages and dependencies
        clean       remove object files
        doc         show documentation for package or symbol
        env         print Go environment information
        bug         start a bug report
        fix         run go tool fix on packages
        fmt         run gofmt on package sources
        generate    generate Go files by processing source
        get         download and install packages and dependencies
        install     compile and install packages and dependencies
        list        list packages
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         run go tool vet on packages

Use "go help [command]" for more information about a command.

Additional help topics:

        c           calling between Go and C
        buildmode   description of build modes
        filetype    file types
        gopath      GOPATH environment variable
        environment environment variables
        importpath  import path syntax
        packages    description of package lists
        testflag    description of testing flags
        testfunc    description of testing functions

Use "go help [topic]" for more information about that topic.
```


### Golang Feature
* 程序从 `main` 方法开始运行

* `import` 导入包

* 方法名 属性名 __首字母大写__ 表示可以导出（被外部使用）

* 命名返回值: `func split(sum int) (x, y int) { ... }`

* 短变量声明: `:=`

* 强大的 `for`： for循环，while，无限循环

* `defer` : defer 语句会将函数推迟到外层函数返回之后执行.
    * defer 栈: 先进后出。

* 指针： 
    * 指针保存了值的内存地址
    * 类型 `*T` 是指向 `T` 类型值的 __指针__。其零值为 nil
    * `&` 操作符会生成一个指向其操作数的指针
    * `*` 操作符表示指针指向的底层值

* 结构体:
    * 结构体指针: `(*p).X`
    * 结构体的方法： 方法是 __接收者__ 参数的函数 `func (v Vertex) Abs() float64 { ... }`

* 切片：
    * 类型 `[]T` 表示一个元素类型为 `T` 的切片
    * 切片的零值是 `nil`

* 内建函数 `make`

* `range`:
    * `for` 循环的 `range` 形式可遍历切片或映射: `for i, v := range pow { ... }`

* 映射: 
    * `map`

* 函数的闭包:
    * 返回一个函数： `func adder() func(int) int { /// }`

* __接口__:
    * __接口类型__ 是由一组方法签名定义的集合
    * 实现接口: 类型通过 __实现一个接口的所有方法来__ 实现该接口
        * https://tour.go-zh.org/methods/10
    * 空接口： `interface{}`
        * 空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）
        * https://tour.go-zh.org/methods/14

* 类型断言
    * 类似反射

* 错误:
    * Go 程序使用 `error` 值来表示错误状态

* __Go 程__:
    * `go` 程（goroutine）是由 Go 运行时管理的轻量级线程
        * __Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步__。 `sync` 包提供了这种能力，不过在 Go 中并不经常用到，因为还有其它的办法（见下一页）。

* __信道__:
    * `chan` 类型
    * 信道是带有类型的管道，你可以通过它用信道操作符 `<-` 来发送或者接收值:
        * `ch <- v`    // 将 v 发送至信道 ch。
        * `v := <-ch`  // 从 ch 接收值并赋予 v
    * 带缓冲的信道

* `select` 语句:
    * `select` 语句使一个 Go 程可以等待多个通信操作
    * `select` 会阻塞到某个分支可以继续执行为止

* sync.Mutex:
    * https://tour.go-zh.org/concurrency/9
    * 如果我们并不需要通信呢？比如说，若我们只是想保证每次只有一个 Go 程能够访问一个共享的变量，从而避免冲突？
    * 这里涉及的概念叫做 __互斥（mutual_exclusion）__ ，我们通常使用 __互斥锁（Mutex）__ 这一数据结构来提供这种机制。
    * Go 标准库中提供了 sync.Mutex 互斥锁类型及其两个方法： `Lock` `Unlock`
