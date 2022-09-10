## 10 结构（struct）与方法（method）
* Go 通过 _类型别名_（alias types）和 _结构体_ 的形式支持用户自定义类型，或者叫定制类型

* 组成结构体类型的那些数据称为 __字段（fields）__

* 结构体的概念在软件工程上旧的术语叫 ADT（抽象数据类型：Abstract Data Type），在一些老的编程语言中叫 __记录（Record）__


## 10.1 结构体定义
* https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/10.1.md

* 结构体定义的一般方式如下：
    ```go
    type identifier struct {
        field1 type1
        field2 type2
        ...
    }
    ```

* `type T struct {a, b int}` 也是合法的语法，它更适用于简单的结构体

* 使用 new
    *  使用 `new` 函数给一个新的结构体变量分配内存，它 _返回指向已分配内存的指针_：`var t *T = new(T)`
    ```go
    var t *T
    t = new(T) // 这条语句的惯用方法是： `t := new(T)`; 变量 t 是一个指向 T的指针，此时结构体字段的值是它们所属类型的零值
    ```

    * 声明 `var t T` 也会给 `t` 分配内存，并零值化内存，但是这个时候 `t` 是类型T. 在这两种方式中，t 通常被称做类型 T 的一个实例（instance）或对象（object）

* 结构体的内存布局

* 递归结构体

* 结构体转换


## 10.2 使用工厂方法创建结构体实例
* 如果想知道结构体类型T的一个实例占用了多少内存，可以使用：`size := unsafe.Sizeof(T{})`

* 10.2.2 map 和 struct vs new() 和 make()
    * new 和 make 这两个内置函数已经在第 7.2.4 节通过切片的例子说明过一次: https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/07.2.md#724-new-%E5%92%8C-make-%E7%9A%84%E5%8C%BA%E5%88%AB


## 10.3 使用自定义包中的结构体
* structPack.go:
```go
package structPack

type ExpStruct struct {
    Mi1 int
    Mf1 float32
}
```

* main.go:
```go
package main

import (
    "fmt"
    "./struct_pack/structPack"
)

func main() {
    struct1 := new(structPack.ExpStruct)
    struct1.Mi1 = 10
    struct1.Mf1 = 16.

    fmt.Printf("Mi1 = %d\n", struct1.Mi1)
    fmt.Printf("Mf1 = %f\n", struct1.Mf1)
}
```


## 10.4 带标签的结构体
* 结构体中的字段除了有名字和类型外，还可以有一个可选的标签（tag）：它是一个附属于字段的字符串，可以是文档或其他的重要标记

* 标签的内容不可以在一般的编程中使用，只有包 `reflect` 能获取它


## 10.5 匿名字段和内嵌结构体
* 结构体可以包含一个或多个 __匿名（或内嵌）字段__，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型就是字段的名字。匿名字段本身可以是一个结构体类型，即 __结构体可以包含内嵌结构体__

* Go 语言中的 _继承_ 是通过 __内嵌__ 或 __组合__ 来实现的


* 10.5.2 内嵌结构体

* 10.5.3 命名冲突
    * 当两个字段拥有相同的名字（可能是继承来的名字）时该怎么办呢？
        1. 外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式；
        2. 如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）。没有办法来解决这种问题引起的二义性，必须由程序员自己修正。


## 10.6 方法
* 一个类型加上它的方法等价于面向对象中的一个 __类__

* 类型 T（或 *T）上的所有方法的集合叫做类型 T（或 *T）的方法集

* `func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }`
    * __如果 recv 是一个指针，Go 会自动解引用__

* 10.6.2 函数和方法的区别

* 10.6.3 指针或值作为接收者


## 10.7 类型的 String() 方法和格式化描述符 
* 当定义了一个有很多方法的类型时，十之八九你会使用 `String()` 方法来 __定制__ 类型的字符串形式的输出
    * 如果类型定义了 `String()` 方法，它会被用在 `fmt.Printf()` 中生成默认的输出：等同于使用格式化描述符 `%v` 产生的输出。
    
    * 还有 `fmt.Print()` 和 `fmt.Println()` 也会 _自动使用 String() 方法_。

* MARK: 有点类似 PHP 类里面的 `__toString()`


## 10.8 垃圾回收和 SetFinalizer
* 略
