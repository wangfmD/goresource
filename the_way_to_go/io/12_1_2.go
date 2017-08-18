package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input!!")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Println("input:", input)
	}
}
