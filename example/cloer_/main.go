package main

import (
	"fmt"
)

func close(i int) func() int {
	i = i + 1
	return func() int {
		i = i + 10
		return i
	}
}

func main() {
	a := close(10)()
	fmt.Println(a)
}
