## Go 基础
* https://tour.golang.org/

### 包
* 每个 Go 程序都是由包构成的。

* 程序从包的 `main` 方法开始运行。

* `package` 关键字 __声明包__ (TODO: tour 暂时没有这个)

* `import` 关键字 __导入包__
    ```go
    import "fmt"
    import "math"

    # 或者: 分组形式导入
    import (
        "fmt"
        "math"
    )
    ```

* 导出名
    * 如果一个名字(方法名)以 __大写字母开头，那么它就是已导出的__ 


### 函数
* 注意: __数据类型__ 在 __变量名__ 的 __后面__

* 当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略
    * eg: `func add(x, y int)` // x, y 都是 int

* 多值返回: 函数可以返回任意数量的返回值
    ```go
    // 返回值是 y 和 x, string 类型
    func swap(x, y string) (string, string) {
        return y, x
    }
    ```

* 命名返回值: Go 的返回值可被命名，它们会 _被视作定义在函数顶部的变量_
    * eg: `func split(sum int) (x, y int)` // 函数结尾直接 `return`, 返回的就是 `x` 和 `y`


### 变量
* `var` 语句用于声明一个变量列表，跟函数的参数列表一样，__类型在最后__

* 变量的初始化:
    * 如果 _初始化值已存在，则可以省略类型_;

    * eg: `var c, python, java = true, false, "no!"`

* 短变量声明: `:=`
    * _在函数中_(因此 `:=` 结构不能在函数外使用)，简洁赋值语句 `:=` 可在类型明确的地方代替 `var` 声明


### 基本类型
```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
     // 表示一个 Unicode 码点

float32 float64

complex64 complex128
```

* `int`, `uint` 和 `uintptr` 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽

* 零值:
    * 没有明确初始值的变量声明会被赋予它们的 __零值__

    * 零值是：
        * 数值类型为 `0`;
        * 布尔类型为 `false`;
        * 字符串为 `""`（空字符串）;

* 类型转换:
    * 表达式 `T(v)` 将值 `v` 转换为类型 `T`

    * eg: `var i int = 42; var f float64 = float64(i);`

* 类型推导
    * 在声明一个变量而不指定其类型时（即使用不带类型的 `:=` 语法或 `var =` 表达式语法），变量的 __类型由右值推导得出__


### 常量
* `const` 关键字声明

* 常量可以是字符、字符串、布尔值或数值 (只能是基础类型的意思?)

* 常量不能用 `:=` 语法声明

* _数值常量_ 是高精度的 __值__
