package main

import (
	"./pack1"
	"fmt"
)

func main()  {
	str := pack1.ReturnStr()
	fmt.Println(str)
	fmt.Println(pack1.Pack1Int)
}
