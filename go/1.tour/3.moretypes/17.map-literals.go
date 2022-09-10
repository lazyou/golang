package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map 的文法: map 的文法跟结构体文法相似，不过必须有键名
var m = map[string]Vertex{
	"key1": Vertex {
		40.68433, -74.39967,
	},
	"key2": Vertex {
		37.42202, -122.08408,
	},
}

func main() {
	// map[key1:{40.68433 -74.39967} key2:{37.42202 -122.08408}]
	fmt.Println(m)
}