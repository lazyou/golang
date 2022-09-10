package main

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

var errNotFound error = errors.New("Not found error")

func Test(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math - square root of negative number")
	} else {
		return f, nil
	}
}

func main() {
	fmt.Printf("error: %v", errNotFound)
	fmt.Println()

	fmt.Printf("error: %v", errors.New("New error 2"))
	fmt.Println()

	if f, err := Test(-1); err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Println()
	} else {
		fmt.Println(f)
	}
}
