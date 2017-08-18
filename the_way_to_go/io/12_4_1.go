package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "Alice"
	if len(os.Args) > 1 {
		// 	for _, value := range os.Args {
		// 		fmt.Println(value)
		// 	}
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println(who)
}
