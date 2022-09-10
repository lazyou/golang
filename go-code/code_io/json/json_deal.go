package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	add1 := &Address{"private", "Fujian", "China"}
	add2 := &Address{"word", "Boom", "Belgium"}
	vCard := VCard{"Jan", "Kersschot", []*Address{add1, add2}, "none"}

	jsonStr, _ := json.Marshal(vCard)
	fmt.Printf("JSON format: %s\n", jsonStr)

	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()

	encoder := json.NewEncoder(file)
	// 将数据对象 v 的json编码写入 io.Writer w 中
	err := encoder.Encode(vCard)
	if err != nil {
		log.Println("Error in encoding json")
	}
}
