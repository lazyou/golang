## 流程控制语句


### for
* Go 只有一种循环结构： `for` 循环, 三部分组成
    * 初始化语句：在第一次迭代前执行
    * 条件表达式：在每次迭代前求值 (条件表达式 __不允许加括号__)
    * 后置语句：在每次迭代的结尾执行

* 初始化语句和后置语句是 _可选_ 的

* 条件表达式 __不允许加括号__: `for i := 0; i < 10; i++`

* 用 `for` 实现 `while`: `for sum < 1000`

* 无限循环: 
    ```go
    for {
        // TODO
    }
    ```


### if
* 表达式外无需小括号 `( )` ，而大括号 `{ }` 则是必须的

* 该语句声明的 __变量作用域__ 仅在 `if` 之内


### if 和 else


### switch
* `switch` 运行第一个值等于条件表达式的 `case` 语句: `switch os := runtime.GOOS; os { case xx: ... default: xxx }`

* Go 的 `switch` __只运行选定的 `case`__，而非之后所有的 `case` 
    * 实际上，Go 自动提供了在这些语言中每个 `case` 后面所需的 `break` 语句

* `switch` 的求值顺序: `switch` 的 `case` 语句从上到下顺次执行，直到 _匹配成功时停止_

* 没有条件的 `switch`
    * 没有条件的 switch 同 `switch true` 一样

    * 这种形式能将一长串 `if-then-else` 写得更加清晰

    * eg:
        ```go
        t := time.Now()
        switch {
            case t.Hour() < 12:
                fmt.Println("Good morning!")
            case t.Hour() < 17:
                fmt.Println("Good afternoon.")
            default:
                fmt.Println("Good evening.")
        }
        ```


### defer
* `defer` 语句会 __将函数推迟到外层函数返回之后执行__
    * 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用

* `defer` 栈: __先进后出__
