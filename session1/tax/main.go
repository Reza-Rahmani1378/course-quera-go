package main

import (
	"fmt"
	"math"
)

func main() {

	var income int
	var tax float64
	var tax1 float64 = 0.05
	var tax2 float64 = 0.1
	var tax3 float64 = 0.15
	var tax4 float64 = 0.2

	fmt.Scanf("%d", &income)

	if income >= 10000 {

		tax += math.Floor(tax1*10000*100) / 100
		income -= 10000
		if income > 0 && income <= 40000 {
			tax += math.Floor(tax2*float64(income)*100) / 100

		}

		if income > 40000 {
			tax += math.Floor(tax2*40000*100) / 100

			income -= 40000
			if income > 0 && income <= 50000 {
				tax += math.Floor(tax3*float64(income)*100) / 100

			}
		}

		if income > 50000 {
			tax += math.Floor(tax3*50000*100) / 100
			income -= 50000
			if income > 0 && income <= 100000 {
				tax += math.Floor(tax4*float64(income)*100) / 100
			}
		}

		if income > 100000 {
			tax += math.Floor(tax4*float64(income)*100) / 100
		}

	} else {
		tax += math.Floor(tax1*float64(income)*100) / 100
	}

	fmt.Println(int(tax))

}
