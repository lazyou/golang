### 教程
* https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md

* http://wiki.jikexueyuan.com/project/the-way-to-go/


### 6.2.1 按值传递（call by value） 按引用传递（call by reference）
* 在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的 __引用类型__ 都是默认使用引用传递（即使没有显式的指出指针）。


### 6.2.2 命名的返回值（named return variables）
* eg: `func getX2AndX3_2(input int) (x2 int, x3 int) { ... }`


### 6.2.3 空白符（blank identifier）
* 空白符 `_` 用来匹配一些不需要的值，然后丢弃掉


### 6.3 传递变长参数
* 函数的最后一个参数是采用 `...type` 的形式，那么这个函数就可以处理一个变长的参数

* 如果参数被存储在一个 slice 类型的变量 slice 中，则可以通过 `slice...` 的形式来传递参数，调用变参函数。


### 6.4 defer 和追踪
* 关键字 defer 允许我们进行一些函数执行完成后的收尾工作，例如：
```go
// 1.关闭文件流 （详见 第 12.2 节）
// open a file  
defer file.Close()


// 2. 解锁一个加锁的资源 （详见 第 9.3 节）
mu.Lock()  
defer mu.Unlock() 


// 3.打印最终报告
printHeader()  
defer printFooter()


// 4.关闭数据库链接
// open a database connection  
defer disconnectFromDB()
```


### 6.5 内置函数
* `close`	    用于管道通信

* `len、cap`	len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；cap 是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）

* `new、make`	new 和 make 均是用于分配内存：
    * `new` 用于值类型和用户定义的类型，如自定义结构，`make` 用于内置引用类型（切片、map 和管道）
    
    * 它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。
    
    * new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针（详见第 10.1 节）。它也可以被用于基本类型：v := new(int)。
    
    * make(T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作（详见第 7.2.3/4 节、第 8.1.1 节和第 14.2.1 节）new() 是一个函数，不要忘记它的括号

* `copy、append`	用于复制和连接切片

* `panic、recover`	两者均用于错误处理机制

* `print、println`	底层打印函数（详见第 4.2 节），在部署环境中建议使用 fmt 包

* `complex、real imag`	用于创建和操作复数（详见第 4.5.2.2 节）


### 接口
* 接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。

* 类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口。

* 接口仅仅是定义方法，好像没啥用？


### 13.1 错误处理
* Go 有一个预先定义的 error 接口类型
```go
type error interface {
	Error() string
}
```

* // 当发生错误必须中止程序时，`panic` 可以用于错误处理模式：
```go
if err != nil {
	panic("ERROR occurred:" + err.Error())
}
```

* 在多层嵌套的函数调用中调用 `panic`，可以马上中止当前函数的执行，所有的 `defer` 语句都会保证执行并把控制权交还给接收到 `panic` 的函数调用者。这样向上冒泡直到最顶层，并执行（每层的） `defer`，在栈顶处程序崩溃，并在命令行中用传给 `panic` 的值报告错误情况：这个终止过程就是 `panicking`。


### 13.3 从 panic 中恢复（Recover）
* `recover`内建函数被用于从 `panic` 或 错误场景中恢复：让程序可以从 panicking 重新获得控制权，停止终止过程进而恢复正常执行。

* `recover` 只能在 `defer` 修饰的函数（参见 6.4 节）中使用：用于取得 `panic` 调用中传递过来的错误值，如果是正常执行，调用 `recover` 会返回 `nil`，且没有其它效果。

* 自定义包中的错误处理和 panicking (【没实践】): https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/13.4.md


### 14.1 并发、并行和协程
* 在 Go 中，应用程序并发处理的部分被称作 `goroutines`（协程），它可以进行更有效的并发运算。在协程和操作系统线程之间并无一对一的关系：协程是根据一个或多个线程的可用性，映射（多路复用，执行于）在他们之上的；协程调度器在 Go 运行时很好的完成了这个工作。

* 协程工作在相同的地址空间中，所以共享内存的方式一定是同步的；这个可以使用 `sync` 包来实现（参见第 9.3 节），不过我们很不鼓励这样做：Go 使用 `channels` 来同步协程（可以参见第 14.2 节等章节）

* 任何 Go 程序都必须有的 `main()` 函数也可以看做是一个协程，尽管它并没有通过 `go` 来启动。协程可以在程序初始化的过程中运行（在 `init()` 函数中）。


### 14.2 协程间的信道
* Go 有一种特殊的类型，通道（channel），就像一个可以用于发送类型化数据的管道，由其负责协程之间的通信，从而避开所有由共享内存导致的陷阱；这种通过通道进行通信的方式保证了同步性。

* 通信操作符 `<-`

* 默认情况下，通信是同步且无缓冲的：在有接受者接收数据之前，发送不会结束。


### 14.4 使用 select 切换协程
* 从不同的并发执行的协程中获取值可以通过关键字 `select` 来完成，它和 `switch` 控制语句非常相似（章节5.3）也被称作通信开关；它的行为像是“你准备好了吗”的轮询机制；`select` 监听进入通道的数据，也可以是用通道发送值的时候。
