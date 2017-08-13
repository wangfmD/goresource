package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePirce float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePirce * s.count
}

type car struct {
	make  string
	model string
	price float32
}

// getValue ...
func (c car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

// showValue ...
func showValue(v valuable) {
	fmt.Printf("Value as the asset is %f\n", v.getValue())
}

func main() {
	s1 := stockPosition{"ti11", 1000.00, 10}
	c1 := car{"dazong", "st", 2000}
	showValue(s1)
	showValue(c1)
}
