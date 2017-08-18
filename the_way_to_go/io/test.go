package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

// func main() {
// 	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
// 	// buf, _ := reader.ReadSlice('\n')
// 	// buf, _ := reader.ReadString('\n')
// 	buf, _ := reader.ReadBytes('\n')
// 	fmt.Printf("%s\n", buf)
// 	// n, _ := reader.ReadSlice('\n')
// 	// n, _ := reader.ReadString('\n')
// 	n, _ := reader.ReadBytes('\n')
// 	fmt.Printf("the line:%s\n", buf)
// 	// fmt.Println(string(n))
// 	fmt.Println(string(n))
// }

func main() {
	srcFile, err := os.Open("test.go")
	if err != nil {
		fmt.Println("open file error")
		return
	}
	defer srcFile.Close()

	srcReader := bufio.NewReader(srcFile)
	for {
		con, err := srcReader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Println(con)
	}
}
