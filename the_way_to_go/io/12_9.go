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
	country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "nanjing", "China"}
	wa := &Address{"word", "shanghai", "China"}
	vc := VCard{"Jan", "ddd", []*Address{pa, wa}, "none"}
	fmt.Printf("%v: \n", vc)
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error Encode")
	}

}
