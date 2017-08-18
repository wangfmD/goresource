package main

import (
	"fmt"
)

type myinterface interface {
	print()
}

type myint int32

// print ...
func (m myint) print() {
	fmt.Println(m)
}

func (m myint) String() string {
	return "yes"
}

func main() {
	var mi myinterface

	var age myint
	age = 10

	// age1 myint:= 10
	fmt.Println(age)
	fmt.Println(mi)
	mi = &age
	mi.print()

	var any interface{}
	any = &age
	if v, ok := any.(myinterface); ok {
		fmt.Println(v, " impletation myinterface")
	} else {
		fmt.Println(v, " not impletion myinterface")
	}

	if sv, ok := mi.(Stringer); ok {
		fmt.Printf("v implements String(): %s\n", sv) // note: sv, not v
	} else {
		fmt.Println("222")
	}
}

type Stringer interface {
	String() string
}
