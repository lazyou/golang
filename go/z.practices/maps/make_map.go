package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var mm = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	// 映射将键映射到值
	// 映射的零值为 nil 。nil 映射既没有键，也不能添加键

	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	// {40.68433 -74.39967}
	fmt.Println(m["Bell Labs"])

	fmt.Println(mm)
}
