package main

import (
	"encoding/json"
	"log"
	"fmt"
)

// 构体成员Tag是和在编译阶段关联到该成员的元信息字符串：
// Color成员的Tag还带了一个额外的 omitempty 选项，表示当Go语言结构体成员为空或零值时不生成JSON对象（这里false为零值）
type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{
		Title: "Casablanca",
		Year: 1942,
		Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
	},
	{
		Title: "Cool Hand Luke",
		Year: 1967,
		Color: true,
		Actors: []string{"Paul Newman"},
	},
}

// json 编码 解码, 以及 结构体 tag 属性的使用
func main() {
	// 把 struct 进行 json 编码
	jsonstr, err := json.Marshal(movies)

	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", jsonstr)

	// [{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingrid Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,"Actors": ["Paul Newman"]}]
	// 注意:
	// Year 字段被解析成 released,
	// Color 字段被解析成 color, 且当 color 是 false 的时候忽略了该字段 (omitempty 的作用)


	// json 解码
	demovies := []Movie{}
	if err := json.Unmarshal(jsonstr, &demovies); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%v", demovies)
	// [{Title:Casablanca Year:1942 Color:false Actors:[Humphrey Bogart Ingrid Ber gman]} {Title:Cool Hand Luke Year:1967 Color:true Actors:[Paul Newman]}]
}
