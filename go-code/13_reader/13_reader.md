### Reader
* `io` 包指定了 `io.Reader` 接口，它表示从数据流的末尾进行读取

* `io.Reader` 接口有一个 `Read` 方法：
```go
func (T) Read(b []byte) (n int, err error)
```

