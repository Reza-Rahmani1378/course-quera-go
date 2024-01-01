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
	fmt.Scan(&number)

	results := make([][]string, number)
	for i := range results {
		results[i] = make([]string, 2)
	}

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < number; i++ {

		name, _ := reader.ReadString('\n')
		name = strings.TrimSuffix(name, "\n")
		scores, _ := reader.ReadString('\n')
		score := strings.Split(scores, " ")
		responeAverage := average(score)

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
	for i := 0; i < len(numbers)-1; i++ {
		f, _ := strconv.ParseFloat(numbers[i], 64)
		sum += f
	}
	lastElement := strings.TrimSuffix(numbers[len(numbers)-1], "\n")
	lastIndex, _ := strconv.ParseFloat(lastElement, 32)
	sum += float64(lastIndex)
	return (float64(sum)) / (float64(len(numbers)))
}
