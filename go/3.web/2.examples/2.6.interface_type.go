package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // 匿名字段Human
	school string
	loan   float32
}

type Employee struct {
	Human   // 匿名字段Human
	company string
	money   float32
}

// Human 对象实现 Sayhi 方法
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human 对象实现 Sing 方法
func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

//Human 对象实现 Guzzle 方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee 重载 Human 的 Sayhi 方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行
}

// Student 实现 BorrowMoney 方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

// Employee 实现 SpendSalary 方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}

// 定义 interface
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

// TODO: 上面的 Men interface 被 Human、Student 和 Employee 实现

/**
通过上面的代码我们可以知道，interface可以被任意的对象实现。

我们看到上面的 Men interface 被 Human、Student 和 Employee 实现。
	* (TODO: 然而通过肉眼我根本看不出来是这样)

同理，一个对象可以实现任意多个 interface，例如上面的 Student 实现了 Men 和 YoungChap 两个 interface。

任意的类型都实现了 空interface (我们这样定义：interface{})，也就是包含 0 个 method 的 interface。
*/
