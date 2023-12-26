package main

import (
	"fmt"
	_ "math"
)

// Grouping Variables
var (
	name      string
	age       int
	isStudent bool
)

func main() {
	var i int
	var fl float64
	var s string
	var b bool

	var num1, num2 int

	num1 = 2
	num2 = 3

	_ = num1 + num2 // Blank Identifier
	name = "Reza"

	fmt.Printf("name is : %v", name)
	fmt.Printf("age is : %v", age)
	fmt.Printf("student is : %v", isStudent)

	fmt.Printf("Zero Value for : \n"+
		"\tint -> %v\n"+
		"\tfloat64 -> %v\n"+
		"\tstring -> %v\n"+
		"\tbool -> %v\n", i, fl, s, b)

	fmt.Print("Enter input:")
	x := input([]int{}, nil)
	fmt.Println("Input:", x)
}

func input(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scan(&d)
	fmt.Print("n", n)
	if n == 1 {
		x = append(x, d)
	}
	return input(x, err)
}
