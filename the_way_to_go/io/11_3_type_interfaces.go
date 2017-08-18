package main

import (
	"fmt"
)

type zfx struct {
	side float32
}

type circle struct {
	radius float32
}

type shaper interface {
	area() float32
}

// area ...
func (z *zfx) area() float32 {
	return z.side * z.side
}

func (c *circle) area() float32 {
	return c.radius * c.radius * 3.14
}

func main() {
	var sh shaper
	z := new(zfx)
	z.side = 4
	sh = z

	if t, ok := sh.(*zfx); ok {
		fmt.Printf("type is %T\n", t)
		fmt.Println(*t)
	}
	if u, ok := sh.(*circle); ok {
		fmt.Printf("type is %T\n", u)
		fmt.Println(*u)
	} else {
		fmt.Println("nukown type")
	}
	a1 := sh.area()
	y := circle{2}

	fmt.Println(a1)
	sh = &y
	fmt.Println(sh.area())
}
