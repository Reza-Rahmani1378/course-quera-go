package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var number int
	fmt.Scanf("%d", &number)

	results := make([][]string, number)
	for i := range results {
		results[i] = make([]string, 2)
	}

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < number; i++ {
		var name string
		fmt.Scanln(&name)
		// fmt.Println("terminal name")
		scores, _ := reader.ReadString('\n')
		// fmt.Println("Terminal Line")
		score := strings.Split(scores, " ")
		for _, el := range score {
			fmt.Println("Score is :", el)
		}
		responeAverage := average(score)

		fmt.Println("Respone Average :", responeAverage)

		results[i][0] = name
		results[i][1] = calculateScore(responeAverage)
	}

	for i := 0; i < number; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(results[i][j], " ")
		}
		fmt.Println()
	}

}

func calculateScore(number float64) string {

	switch {
	case number >= float64(80):
		return "Excellent"
	case number >= float64(60) && number < float64(80):
		return "Very Good"
	case number >= float64(40) && number < float64(60):
		return "Good"
	}

	return "Fair"
}
func average(numbers []string) float64 {

	var sum float64 = 0
	for _, element := range numbers {
		f, _ := strconv.ParseFloat(element, 64)
		sum += f
	}
	fmt.Println("last element is :", numbers[len(numbers)-1])
	lastElement := numbers[len(numbers)-1]
	fmt.Println("Last Element is :", lastElement)
	lastIndex, _ := strconv.ParseFloat(lastElement, 64)
	sum += float64(lastElement)
	fmt.Println("Sum is :", sum)
	fmt.Println("Len is :", len(numbers))
	return (float64(sum)) / (float64(len(numbers)))
}
