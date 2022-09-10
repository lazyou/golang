### 错误
* Go 程序使用 `error` 值来表示错误状态

* error 类型是一个内建接口：
```go
type error interface {
    Error() string
}
```
