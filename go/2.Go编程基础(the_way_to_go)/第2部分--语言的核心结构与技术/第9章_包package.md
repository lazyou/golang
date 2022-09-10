## 9.1 标准库概述
* 像 fmt、os 等这样具有常用功能的内置包在 Go 语言中有 150 个以上，它们被称为标准库，大部分(一些底层的除外)内置于 Go 本身

* eg:
```go
unsafe: 包含了一些打破 Go 语言“类型安全”的命令，一般的程序中不会被使用，可用在 C/C++ 程序的调用中。
syscall-os-os/exec:
    os: 提供给我们一个平台无关性的操作系统功能接口，采用类Unix设计，隐藏了不同操作系统间差异，让不同的文件系统和操作系统对象表现一致。
    os/exec: 提供我们运行外部操作系统命令和程序的方式。
    syscall: 底层的外部包，提供了操作系统底层调用的基本接口。

archive/tar 和 /zip-compress：压缩(解压缩)文件功能。

fmt-io-bufio-path/filepath-flag:
    fmt: 提供了格式化输入输出功能。
    io: 提供了基本输入输出功能，大多数是围绕系统功能的封装。
    bufio: 缓冲输入输出功能的封装。
    path/filepath: 用来操作在当前系统中的目标文件名路径。
    flag: 对命令行参数的操作。　　

strings-strconv-unicode-regexp-bytes:
    strings: 提供对字符串的操作。
    strconv: 提供将字符串转换为基础类型的功能。
    unicode: 为 unicode 型的字符串提供特殊的功能。
    regexp: 正则表达式功能。
    bytes: 提供对字符型分片的操作。
    index/suffixarray: 子字符串快速查询。

math-math/cmath-math/big-math/rand-sort:
    math: 基本的数学函数。
    math/cmath: 对复数的操作。
    math/rand: 伪随机数生成。
    sort: 为数组排序和自定义集合。
    math/big: 大数的实现和计算。 　　

container-/list-ring-heap: 实现对集合的操作。
    list: 双链表。
    ring: 环形链表。

time-log:
    time: 日期和时间的基本操作。
    log: 记录程序运行时产生的日志,我们将在后面的章节使用它。\

encoding/json-encoding/xml-text/template:
    encoding/json: 读取并解码和写入并编码 JSON 数据。
    encoding/xml:简单的 XML1.0 解析器,有关 JSON 和 XML 的实例请查阅第 12.9/10 章节。
    text/template:生成像 HTML 一样的数据与文本混合的数据驱动模板（参见第 15.7 节）。

net-net/http-html:（参见第 15 章）
    net: 网络数据的基本操作。
    http: 提供了一个可扩展的 HTTP 服务器和基础客户端，解析 HTTP 请求和回复。
    html: HTML5 解析器。

runtime: Go 程序运行时的交互操作，例如垃圾回收和协程创建。

reflect: 实现通过程序运行时反射，让程序操作任意类型的变量。
```


## 9.2 regexp 包


## 9.3 锁和 sync 包
* https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/09.3.md

* >通常通过不同线程执行不同应用来实现程序的并发。当不同线程要使用同一个变量时，经常会出现一个问题：无法预知变量被不同线程修改的顺序！

* _经典的做法_ 是一次只能让一个线程对共享变量进行操作。当变量被一个线程改变时(临界区)，我们为它上锁，直到这个线程执行完成并解锁后，其他线程才能访问它

* 在 Go 语言中这种锁的机制是通过 `sync` 包中 `Mutex` 来实现的。


## 9.4 精密计算和 big 包


### 9.5 自定义包和可见性
* import 的一般格式如下: `import "包的路径或 URL 地址"`

* 自定包
    * TODO: 表示不会

* 导入外部安装包:
    * `go install xxx`

* 包的初始化:
    * 初始化 `main` 包然后调用 `main` 函数

* 可见性:
    * 首字母大写


## 9.6 为自定义包使用 godoc
* 注释 _必须_ 以 // 开始并无空行放在声明（包，类型，函数）前

* godoc 会为每个文件生成一系列的网页
    * 进入目录执行 `godoc -http=:6060 -goroot="."`, 在浏览器打开地址：http://localhost:6060


## 9.7 使用 go install 安装自定义包
* `go install` 是 Go 中自动包安装工具：如需要将包安装到本地它会从远端仓库下载包：检出、编译和安装一气呵成。

* go install 使用了 GOPATH 变量(详见第 2.2 节)。

* eg:
```go
假设我们要安装一个有趣的包 tideland（它包含了许多帮助示例，参见 项目主页）。

因为我们需要创建目录在 Go 安装目录下，所以我们需要使用 root 或者 su 的身份执行命令。

确保 Go 环境变量已经设置在 root 用户下的 ./bashrc 文件中。

使用命令安装：go install tideland-cgl.googlecode.com/hg。

__可执行文件__ hg.a 将被放到 `$GOROOT/pkg/linux_amd64/tideland-cgl.googlecode.com` 目录下，

__源码文件__ 被放置在 `$GOROOT/src/tideland-cgl.googlecode.com/hg` 目录下，同样有个 hg.a 放置在 _obj 的子目录下。

使用: `import cgl "tideland-cgl.googlecode.com/hg"`
```


## 9.8 自定义包的目录结构、go install 和 go test
    * 重要: https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/09.8.md


## 9.9 通过 Git 打包和安装


## 9.10 Go 的外部包和项目
* Go Walker 支持根据包名在海量数据中查询:
    * https://gowalker.org/


## 9.11 在 Go 程序中使用外部库
