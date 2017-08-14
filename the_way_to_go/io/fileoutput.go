package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile, outputError := os.OpenFile("output", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Println("an error occurred!")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello word!\n"

	for i := 0; i < 11; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
