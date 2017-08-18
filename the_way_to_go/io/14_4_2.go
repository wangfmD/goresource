package main

import (
	"flag"
	// "fmt"
	"os"
)

var newLine = flag.Bool("n", false, "print newline")

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.PrintDefaults()
	flag.Parse()
	// fmt.Println(flag.NArg())
	// fmt.Println(*newLine)
	// s := ""
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {

		if i > 0 {
			s += " "
			if *newLine {
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}
