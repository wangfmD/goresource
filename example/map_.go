package main

import (
	// "bufio"
	"fmt"
)

// m1
func main() {
	// var m1 map[string]string = map[string]string{"name": "www", "age": "20"}
	// fmt.Printf("\"name\": %s\n", m1["name"])
	// mapF1()
	// mapF2()
	mapF3()
}

// mapF1 ...
func mapF1() {
	// init
	var m1 = make(map[string]int)
	m1["a"] = 1
	fmt.Println(m1)

	m2 := make(map[string]int)
	m2["b"] = 2
	fmt.Println(m2)

	m3 := map[string]string{"a": "b", "c": "d"}
	fmt.Println(m3)

	// m4 := make(map[string]string){"a": "b", "c": "d"} error

	// map的值可以为任意类型,可以为func
	m4 := map[string]func() int{
		"a": func() int { return 10 },
		"b": func() int { return 20 },
	}
	fmt.Println(m4["a"]())
	fmt.Println(m4["b"]())
}

// mapF2 ...
func mapF2() {
	m1 := map[string]string{"a0": "bb", "a1": "aa"}
	fmt.Println(m1)
	if val, ok := m1["a0"]; ok {
		fmt.Println("a0 exist  m1 value:", val)
		fmt.Println(ok)
	}
	delete(m1, "a0")
	if val, ok := m1["a0"]; ok {
		fmt.Println("a0 exist! value:", val)
	} else {
		fmt.Println("a0 not exist!")
	}
}

type istring struct {
	s string
}

func new(text string) error {
	return &istring{text}
}

func mapF3() {
	e1 := new("iserr")
	fmt.Println(el)
}
