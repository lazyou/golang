### Stringer
* `fmt` 包中定义的 `Stringer` 是最普遍的接口之一
```go
type Stringer interface {
    String() string
}
```

* Stringer 是一个可以用字符串描述自己的类型
