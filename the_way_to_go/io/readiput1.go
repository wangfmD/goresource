package main

import (
	"fmt"
)

func main() {
	var firstName, lastName string
	fmt.Scanln(&firstName, &lastName)
	fmt.Println("hi:", firstName, lastName)
}
