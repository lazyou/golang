package main

import (
	"fmt"
	"sync"
)

type Info struct {
	mu   sync.Mutex
	age  int64
	name string
}

func main() {
	var info = Info{}
	Update(&info)
	fmt.Println(info)
}

func Update(info *Info) {
	// 有点像事务
	info.mu.Lock()
	info.age = 20
	info.mu.Unlock()
}
