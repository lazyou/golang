# From
* 配合文档： https://books.studygolang.com/gopl-zh/

* 文档代码： https://github.com/adonovan/gopl.io/


### 2.4 赋值
```go
// 交换两个变量的值
x, y := 0, 1
x, y = y, x+y


v = m[key]      // map 查找，失败时返回零值
v = x.(T)       // type 断言，失败时panic异常
v = <-ch        // 管道接收，失败时返回零值（阻塞不算是失败）


_, err = io.Copy(dst, src) // 丢弃字节数
_, ok = x.(T)              // 只检测类型，忽略具体值
```

### 5.9 Panic 异常
* 一般而言，当 panic 异常发生时，程序会中断运行，并立即执行在该 goroutine（可以先理解成线程，在第8章会详细介绍）中被延迟的函数（defer 机制）。

### 5.10. Recover捕获异常
* 通常来说，不应该对 panic 异常做任何处理，但有时，也许我们可以从异常中恢复，至少我们可以在程序崩溃前，做一些操作。


### 7.3. 实现接口的条件
* 一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口。


### 8.1. Goroutines
* 在Go语言中，每一个 __并发__ 的执行单元叫作一个 `goroutine`

* 当一个程序启动时，其主函数即在一个单独的 goroutine 中运行，我们叫它 __main goroutine__


### 8.4. Channels
* 如果说 `goroutine` 是Go语言程序的 __并发体__ 的话，那么 `channels` 则是它们之间的 __通信机制__

```go
ch <- x  // a send statement
x = <-ch // a receive expression in an assignment statement
<-ch     // a receive statement; result is discarded
```

* 无缓存 Channels 有时候也被称为同步 Channels。


# The Go Programming Language

This repository provides the downloadable example programs
for the book, "The Go Programming Language"; see http://www.gopl.io.

These example programs are licensed under a <a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/">Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License</a>.<br/>
<a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by-nc-sa/4.0/88x31.png"/></a>

You can download, build, and run the programs with the following commands:

	$ export GOPATH=$HOME/gobook            # choose workspace directory
	$ go get gopl.io/ch1/helloworld         # fetch, build, install
	$ $GOPATH/bin/helloworld                # run
	Hello, 世界

Many of the programs contain comments of the form `//!+` and `//!-`.
These comments bracket the parts of the programs that are excerpted in the
book; you can safely ignore them.  In a few cases, programs
have been reformatted in an unnatural way so that they can be presented
in stages in the book.

