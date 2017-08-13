package main

import (
	"fmt"
)

// main ...
func main() {
	c := make(chan string, 2)
	c <- "buffer1"
	c <- "buffer2"

	fmt.Println(<-c)
	fmt.Println(<-c)

}
