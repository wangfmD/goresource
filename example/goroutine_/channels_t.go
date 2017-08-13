package main

import (
	"fmt"
)

func main() {
	message := make(chan string)
	go func() { message <- "ip ping" }()
	v := <-message
	fmt.Println(v)

}
