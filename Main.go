package main

import (
	"fmt"
	"strconv"
)

var i int = 27

func main() {
	i := 42 // := is defined new variable

	// var j int
	// j = 23

	// var k int = 66

	fmt.Printf("%v, %T\n", i, i) // value, type

	var j float32
	j = float32(i)

	fmt.Printf("%v, %T\n", j, j)

	// Conver string
	var k string
	k = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", k, k)
}
