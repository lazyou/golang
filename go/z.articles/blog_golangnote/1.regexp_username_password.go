package main

import (
	"fmt"
	"regexp"
)

func CheckPasswordAndUsername(inputStr string) bool {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", inputStr); !ok {
		return false
	}

	return true
}

func main() {
	fmt.Println(CheckPasswordAndUsername("aa"))
	fmt.Println(CheckPasswordAndUsername("admin"))
}
