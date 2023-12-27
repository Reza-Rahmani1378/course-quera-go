package main

import "fmt"

func main() {
	var number int
	fmt.Scanf("%d", &number)
	names := make([]string, number)
	// var averages[n] float32
	// averages := make([]float32 , number)

	fmt.Print(names)
	for i := 0; i < len(names); i++ {
		/* if names[i] == "" {
			fmt.Println("Yes")
		} */

	}

	for _, element := range names {

		fmt.Print(element)

	}

}

func input(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		x = append(x, d)
	}
	return input(x, err)
}

func average(numbers []float32) float32 {

	return 0
}
