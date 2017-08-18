package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func test() {
	// inputFile, _ := os.Open("input")
	inputFile, ferro := os.Open("input")
	if ferro != nil {
		fmt.Println("open file error")
		return
	}

	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError != nil {
			return
		}
	}
}

func test1() {
	inputFile, inputError := os.Open("input")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, _, readerError := inputReader.ReadLine()
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}
func main() {
	test1()
	// test()
}
