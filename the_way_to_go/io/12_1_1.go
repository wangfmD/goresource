package main

import (
	"fmt"
)

func main() {
	var (
		firstName, lastName, s string
		i                      int
		f                      float32
		input                  = "56.12/ 322/ go"
		format                 = "%f / %d / %s"
	)

	fmt.Println("please input!!")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("hi %s %s\n", firstName, lastName)
	fmt.Scanf(input, format, &f, &i, &s)
	fmt.Println("::::", f, i, s)

}
