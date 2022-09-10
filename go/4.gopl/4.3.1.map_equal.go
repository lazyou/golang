package main

import "fmt"

// map[string]int 比较方法实现
func main()  {
	ages1 := map[string]int{
		"lin":1,
		"guan":2,
		"zhang":3,
	}

	ages2 := map[string]int{
		"lin":1,
		"guan":2,
		"zhang":3,
	}

	ages3 := map[string]int{
		"diff":1,
		"guan":2,
		"zhang":3,
	}

	fmt.Println(mapEqual(ages1, ages2))		// true
	fmt.Println(mapEqual(ages1, ages3))		// false
}

func mapEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for xKey, xValue := range x {
		yValue, ok := y[xKey]

		if !ok || yValue != xValue {
			return false
		}
	}

	return true
}