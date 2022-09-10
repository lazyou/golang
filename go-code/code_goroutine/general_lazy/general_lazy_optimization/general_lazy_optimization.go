package main

import "fmt"

// 生成器优化: 每当取得一个值时，下一个值即被计算 (而不是像之前版本不断做计算)


type Any interface{}
type EvalFunc func(Any) (Any, Any)

// TODO: 表示没看懂
func main() {
	evenFunc := func(state Any) (Any, Any) {
		os := state.(int)
		ns := os + 2
		return os, ns
	}

	even := BuildLazyIntEvaluator(evenFunc, 0)

	fmt.Println(even()) // 断点debug 用

	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%vth even: %v\n", i, even())
	//}
}

func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	retValChan := make(chan Any)

	loopFunc := func() {
		var actState Any = initState
		var retVal Any

		// TODO: 这里什么时候跳出执行的? 难道不是一直死循环吗? 哪里实现了惰性生成器?
		for {
			//fmt.Println("evalFunc")
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}

	retFunc := func() Any {
		return <- retValChan
	}

	go loopFunc()

	return retFunc
}

func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
	ef := BuildLazyEvaluator(evalFunc, initState)

	return func() int {
		return ef().(int)
	}
}