// https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/11.1.md
package main

import (
	"fmt"
)

type shaper interface {
	area() float32
}

type square struct {
	side float32
}

func (s *square) area() float32 {
	return s.side * s.side
}

type rectangle struct {
	heigth, weigth float32
}

func (r rectangle) area() float32 {
	return r.heigth * r.weigth
}

func main() {
	// s1 := new(square)
	// s1.side = 5

	// s1 := square{5}

	var s1 square
	s1.side = 5

	r := rectangle{4, 5}
	// var shaper1 shaper
	// shaper1 = s1
	// var shaper1 shaper = &s1
	sha := []shaper{&s1, r}
	for _, value := range sha {
		fmt.Printf("area: %f\n", value.area())
	}
	// fmt.Printf("area: %f", shaper1.area())
}
