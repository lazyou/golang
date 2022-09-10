### 新旧模型对比：任务和worker 
* 假设我们需要处理很多任务；一个worker处理一项任务。

### 旧模式：使用共享内存进行同步
* 由各个任务组成的任务池共享内存；为了同步各个worker以及避免资源竞争，我们需要对任务池进行 __加锁__ 保护：

* `sync.Mutex` 是互斥锁：它用来在代码中保护临界区资源：同一时间只有一个go协程（goroutine）可以进入该临界区
```go
type Pool struct {
    Mu      sync.Mutex
    Tasks   []*Task
}
```


* 一个 worker 先将 pool锁定，从 pool 获取第一项任务，再解锁和处理任务。加锁保证了同一时间只有一个 go 协程可以进入到pool中：一项任务有且只能被赋予一个worker。
```go
func Worker(pool *Pool) {
    for {
        pool.Mu.Lock()

        // begin critical section:
        task := pool.Tasks[0]        // take the first task
        pool.Tasks = pool.Tasks[1:]  // update the pool of tasks
        // end critical section

        pool.Mu.Unlock()

        process(task)
    }
}
```

* 旧模式缺点: 
    * 但是当工作协程数量很大，任务量也很多时，处理效率将会因为频繁的加锁/解锁开销而降低。
    * 当工作协程数增加到一个阈值时，程序效率会急剧下降，这就成为了瓶颈。
    * TODO: 难道新模式使用 channel 的本质不也是加锁吗? 能解决旧模式的缺点?


### 新模式：使用通道 -- Master-Worker 模式
* 使用通道进行同步：使用一个通道接受需要处理的任务，一个通道接受处理完成的任务（及其结果）。
    * worker在协程中启动，其数量N应该根据任务数量进行调整。

* 主线程扮演着Master节点角色，可能写成如下形式：
```go
func main() {
    pending, done := make(chan *Task), make(chan *Task)
    go sendWork(pending)       // put tasks with work on the channel

    for i := 0; i < N; i++ {   // start N goroutines to do work
        go Worker(pending, done)
    }

    consumeWork(done)          // continue with the processed tasks
}
```

* worker的逻辑比较简单：从pending通道拿任务，处理后将其放到done通道中：
```go
func Worker(in, out chan *Task) {
    for {
        t := <-in
        process(t)
        out <- t
    }
}
```

* 这里 __并不使用锁__：从通道得到新任务的过程没有任何竞争。
    * 随着任务数量增加，worker数量也应该相应增加，同时性能并不会像第一种方式那样下降明显。

* 在pending通道中存在一份任务的拷贝，第一个worker从pending通道中获得第一个任务并进行处理，这里并不存在竞争
    * 对一个通道读数据和写数据的整个过程是 __原子性__ 的


### 怎么选择是该使用锁还是通道？
* 使用锁的情景：
    * 访问共享数据结构中的缓存信息;
    * 保存应用程序上下文和状态信息数据;

* 使用通道的情景：
    * 与异步操作的结果进行交互;
    * 分发任务;
    * 传递数据所有权;
