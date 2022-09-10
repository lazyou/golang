package main

import "fmt"

func main() {
	fmt.Println("Starting the program 【输出】")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program 【没有输出】")
}
