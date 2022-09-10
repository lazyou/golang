package main

import "fmt"

// 给定一个字符串列表，下面的 nonempty 函数将在原有slice内存空间之上 __返回不包含空字符串的列表__
func main() {
	strings := []string{"aa", "", "cc"}
	fmt.Printf("%q\n", nonempty(strings)) // ["aa" "cc"]
	fmt.Printf("%q\n", strings)           // ["aa" "cc" "cc"]

	strings2 := []string{"aa", "", "cc"}
	fmt.Printf("%q\n", nonempty2(strings2)) // ["aa" "cc"]
	fmt.Printf("%q\n", strings2)            // ["aa" "cc" "cc"]
}

func nonempty(strings []string) []string {
	i := 0

	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]

	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}

	return out
}
