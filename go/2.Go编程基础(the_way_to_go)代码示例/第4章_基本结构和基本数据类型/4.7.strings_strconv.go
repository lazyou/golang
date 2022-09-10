package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"

	fmt.Printf("是否有前缀 %t\n", strings.HasPrefix(str, "Th"))

	fmt.Printf("%d\n", strings.Index(str, "an"))

	fmt.Printf("%d\n", strings.Count(str, "a"))

	strRepeat := strings.Repeat("Hi there! ", 3)
	fmt.Printf("%s", strRepeat)

	// more ...
}
