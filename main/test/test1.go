package main

import (
	"fmt"
	"errors"
)

type mstruct struct {
	s string
}

func main() {
	t1 := &mstruct{"wfm"}
	fmt.Println(t1.s)
	err := errors.New("my error")
	fmt.Printf("%T", err)
}
