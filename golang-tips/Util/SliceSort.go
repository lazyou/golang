package Util

import "fmt"

type IntSlice []int

func (p IntSlice) Len() int {
	return len(p)
}

func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	slice := IntSlice{1, 2, 10, 20, 3, 4, 30, 40}

	for i := 0; i < slice.Len(); i++ {
		for j := 0; j < slice.Len(); j++ {
			if slice.Less(i, j) {
				slice.Swap(i, j)
			}
		}
	}

	fmt.Println(slice)
}
