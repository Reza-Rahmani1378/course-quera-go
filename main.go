package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Hello World.")

	x := 12.3353
	var y float32 = float32(math.Floor(x*100) / 100)

	fmt.Println(y)

}
