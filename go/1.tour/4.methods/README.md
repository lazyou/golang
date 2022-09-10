## 方法和接口

### 方法
* __Go 没有类__。不过你可以为结构体类型定义 __方法__

* 方法就是一类带特殊的 __接收者__ 参数的函数

* 方法接收者在它自己的参数列表内，位于 `func` 关键字和方法名之间 (方法的声明看起来也是怪怪的)

* eg: 结构体的方法可以理解为 一个类型为 `Vertex` 的接收者
    ```go
    type Vertex struct {
        X, Y float64
    }

    // 结构体放在 func 后面
    func (v Vertex) Abs() float64 {
        return math.Sqrt(v.X*v.X + v.Y*v.Y)
    }
    ```

* 方法即函数
    * 记住：__方法只是个带接收者参数的函数__

* 也可以为非结构体类型声明方法
    * 只能为在 _同一包_ 内定义的类型的接收者声明方法，而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法
    * （译注：1.就是接收者的类型定义和方法声明必须在同一包内； 2.不能为内建类型声明方法。）

* eg:
    ```go
    // https://go-zh.org/pkg/builtin/#float64
    // type float64 float64 (float64 是所有IEEE-754 64位浮点数的集合)
    type MyFloat float64

    func (f MyFloat) Abs() float64 {
        if f < 0 {
            return float64(-f)
        }
        return float64(f)
    }
    ```


### 指针接收者
* 可以为指针接收者声明方法

* 这意味着对于某类型 `T`，接收者的类型可以用 `*T` 的文法。（此外，`T` 不能是像 `*int` 这样的指针。）

* 指针接收者的方法可以修改接收者指向的值。由于方法经常需要修改它的接收者，__指针接收者__ 比 _值接收者_ 更常用

* eg:
```go
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
    // 指针接受者, 这里操作 struct 里的值会影响到源 struct, 不像之前的值接受者
	v.X = v.X * f
	v.Y = v.Y * f
}
```


### 指针与函数
* eg:
    ```go
    type Vertex struct {
        X, Y float64
    }

    func Scale(v *Vertex, f float64) {
        v.X = v.X * f
        v.Y = v.Y * f
    }

    func main() {
        v := Vertex{3, 4}
        Scale(&v, 10)
        fmt.Println(v) // 此时 v 结构体里属性的值就被改变了 (使用指针作为函数的参数可以省空间, 假设 v 是一个非常大的结构体)
    }
    ```


### 方法与指针重定向
* 比较前两个程序，你大概会注意到 _带指针参数的函数_ 必须接受一个指针：
    ```go
    var v Vertex
    ScaleFunc(v, 5)  // 编译错误！
    ScaleFunc(&v, 5) // OK
    ```

* 而以 _指针为接收者的方法_ 被调用时，__接收者既能为值又能为指针__：
    ```go
    var v Vertex
    v.Scale(5)  // OK, 此处为方便起见，Go 会将语句 `v.Scale(5)` 解释为 `(&v).Scale(5)`  -- TODO: 这句话显然还是比较难理解的

    p := &v
    p.Scale(10) // OK
    ```

* 同样的事情也发生在相反的方向

* 接受一个值作为参数的函数必须接受一个指定类型的值：
    ```go
    var v Vertex
    fmt.Println(AbsFunc(v))  // OK
    fmt.Println(AbsFunc(&v)) // 编译错误！
    ```

* 而以值为接收者的方法被调用时，接收者既能为值又能为指针：
    ```go
    var v Vertex
    fmt.Println(v.Abs()) // OK, 方法调用 `p.Abs()` 会被解释为 `(*p).Abs()`

    p := &v
    fmt.Println(p.Abs()) // OK
    ```


### 选择值或指针作为接收者
* 使用指针接收者的原因有二：
    1. 首先，方法能够修改其接收者指向的值;

    2. 其次，这样可以避免在每次调用方法时复制该值。 _若值的类型为大型结构体时，这样做会更加高效_.

* 通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混


### 接口
* __接口类型__ 是由一组 _方法_ 签名定义的集合

* 实现接口的方式: _类型_ 通过 __实现一个接口的所有方法__ 来 __实现该接口__

* 即便接口内的具体值为 `nil`，方法仍然会被 `nil` 接收者调用 (也就不会出现空指针异常)
    * nil 接口值既不保存值也不保存具体类型

* 指定了零个方法的接口值被称为 __空接口__： `interface{}`
    * 空接口可保存任何类型的值.(因为每个类型都至少实现了零个方法.)


### 类型断言
* __类型断言__ 提供了访问接口值底层具体值的方式
    * `t := i.(T)` || `t, ok := i.(T)`

    * 该语句断言接口值 `i` 保存了具体类型 `T`，并将其底层类型为 T 的值赋予变量 t


### 错误
* `error` 类型是一个内建接口：
    ```go
    type error interface {
        Error() string
    }
    ```

* 通常函数会返回一个 `error` 值，调用的它的代码应当判断这个 _错误是否等于 `nil`_ 来进行错误处理

### Reader


### 图像
