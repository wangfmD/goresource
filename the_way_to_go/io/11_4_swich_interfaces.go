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

	switch t := sh.(type) {
	case *zfx:
		fmt.Printf("type:%T, value:%v\n", *t, *t)
	case *circle:
		fmt.Printf("type:%T, value:%v\n", *t, *t)
	}
}
