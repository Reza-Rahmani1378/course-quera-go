package main

import (
	"bufio"
	"fmt"
	"os"
)

// Grouping Variables
var (
	name      string
	age       int
	isStudent bool
)

func main() {

	fmt.Println("Enter text: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	// input, _ := reader.ReadString('\n')

	// remove the delimeter from the string
	/* var name string
	_, _ = fmt.Scanf("%s", &name)
	fmt.Println("Termiate...")
	fmt.Println(name)
	test := strings.TrimSuffix(input, "\n")
	fmt.Println(test) */

	var name string
	fmt.Scanln(&name)

	fmt.Println("Terminate")
	fmt.Println(name)

	var result string
	result, _ = reader.ReadString('\n')
	fmt.Println("Terminate Line")
	fmt.Println(result)

	/* var test string
	fmt.Scanln(&test)
	fmt.Print("Terminal With Enter Or Space") */
	/* var i int
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
	fmt.Println("Input:", x) */
}

/* func input(x []int, err error) []int {
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
} */
