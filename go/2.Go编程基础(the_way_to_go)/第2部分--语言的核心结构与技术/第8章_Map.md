## 8.0 Map
* `map` 是一种特殊的数据结构：一种元素对（pair）的无序集合，pair 的一个元素是 key，对应的另一个元素是 value，所以这个结构也称为 _关联数组_ 或 _字典_


## 8.1 声明、初始化和 make
* map 是引用类型，可以使用如下声明：
    ```go
    var map1 map[keytype]valuetype
    var map1 map[string]int
    ```

* 声明的时候不需要知道 map 的长度，map 是可以动态增长的

* 未初始化的 map 的值是 `nil`

* `key` 可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float

* `value` 可以是任意类型的；__通过使用空接口类型__（详见第 11.9 节），我们可以存储任意值，但是使用这种类型作为值时需要先做一次类型断言（详见第 11.3 节）。

* map 是 __引用类型__ 的： 内存用 `make` 方法来分配
    * eg: `var map1 = make(map[keytype]valuetype)` || `map1 := make(map[keytype]valuetype)`

    * 不要使用 new，__永远用 make 来构造 map__

* 8.1.2 map 容量: 动态的伸缩

* 8.1.3 用切片作为 map 的值
    ```go
    mp1 := make(map[int][]int)
    mp2 := make(map[int]*[]int)
    ```


## 8.2 测试键值对是否存在及删除元素
* `val1, isPresent = map1[key1]`

* 如果你只是想判断某个 key 是否存在而不关心它对应的值到底是多少，你可以这么做：
    * `_, ok := map1[key1]` // 如果key1存在则ok == true，否则ok为false

* 删除: `delete(map1, key1)`


## 8.3 for-range 的配套用法
* 使用 for 循环构造 map：
    ```go
    for key, value := range map1 {
        ...
    }
    ```

* 如果只想获取 key，你可以这么使用：
    ```go
    for key := range map1 {
        fmt.Printf("key is: %d\n", key)
    }
    ```


## 8.4 map 类型的切片
* 假设我们想获取一个 `map` 类型的切片，我们必须使用两次 `make()` 函数，_第一次分配切片，第二次分配 切片中每个 map 元素_

* eg:
    ```go
    // Version A:
    items := make([]map[int]int, 5)

    for i:= range items {
        items[i] = make(map[int]int, 1)
        items[i][1] = 2
    }

    fmt.Printf("Version A: Value of items: %v\n", items)
    ```


## 8.5 map 的排序
* __map 默认是无序的__，不管是按照 key 还是按照 value 默认都不排序

* _如果你想为 map 排序_，需要将 key（或者 value）拷贝到一个 _切片_，再对切片排序（使用 `sort` 包，详见第 7.6.6 节）

* 但是如果你想要一个排序的列表你最好使用结构体切片，这样会更有效：
    ```go
    type name struct {
        key string
        value int
    }
    ```


## 8.6 将 map 的键值对调
