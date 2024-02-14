package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var n int
	fmt.Scan(&n)

	results := make([][]string, n)
	for i := range results {
		results[i] = make([]string, 2)
	}

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < n; i++ {

		spaceCraft, _ := reader.ReadString('\n')
		spaceCraft = strings.TrimSuffix(spaceCraft, "\n")
		spaceCrafts := strings.Split(spaceCraft, " ")
		// responeAverage := average(score)

		results[i][0] = spaceCrafts[0]
		results[i][1] = "1"
	}

}

func numberOfSequences(spaceCrafts []string, n int) int {

	if n >= 1 && n < 3 {
		return 0
	}

	var sum int = 0

	return 0
}
