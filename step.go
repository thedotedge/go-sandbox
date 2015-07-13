package main

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/go-martini/martini"
)

func main() {
	fmt.Println("My favorite number is ", rand.Intn(10))
	fmt.Println("The time is ", time.Now())
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	_, y, _ := split(17)
	fmt.Println(y)
}


func swap(x, y string) (string, string) {
	return y, x
}


func split(sum int) (x, y, z int) {
	x = sum * 4 / 9
	y = sum - x + 2
	z = sum + x
	return
}