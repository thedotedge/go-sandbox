package main

import (
	. "strconv"
	"fmt"
)

type MyStringer interface {
	String() string
}

type Temperature int32

func (t Temperature) String() string {
	return Itoa(int(t)) + " °C"
}

type Point struct {
	x, y int
}

func NewPoint(x int, y int) *Point {
	p := new(Point)
	p.x = x
	p.y = y
	return p
	//return &Point{x, y}
}


func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func main()  {
	var x MyStringer

	x = Temperature(24)
	fmt.Println(x.String())
	fmt.Printf("%v %T\n", x, x)

	x = NewPoint(10, 12)
	fmt.Println(x.String())
	//fmt.Println(x)
	fmt.Printf("%v %T\n", x, x)
}
