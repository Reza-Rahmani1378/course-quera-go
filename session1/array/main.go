package main

import "fmt"

func main() {
	var arr [3]int
	fmt.Println(arr)
	var arr1 = [3]int{10, 20, 30}
	fmt.Println(arr1)

	// we can define array without size if we initialize it
	var arr2 = [...]int{10, 20, 30, 40}

	fmt.Printf("arr2[0]: %v\n", arr2[0])

	var seasons [4]string
	seasons[0] = "spring"
	seasons[1] = "summer"
	seasons[2] = "fall"
	seasons[3] = "winter"

	fmt.Printf("seasons: %v\n", seasons)

	/*
	 Iteration In Array
	*/

	for i := 0; i < len(arr2); i++ {
		fmt.Printf("Element %d At Index %d\n", arr2[i], i)
	}

	/*
		for index, element := range arr {} // Normal usage of range
		for _, element := range arr {} // Omit index with _ and use element
		for index := range arr {} // Use index only
		for range arr {} // Simply loop over the array
	*/

	for index, element := range arr2 {
		fmt.Printf("Element %d At Index %d\n", element, index)
	}

	/*
		Multi-Dimontional Array
	*/

	var tic_tac_toe [3][3]bool = [3][3]bool{
		{true, true, false},
		{true, false, false},
		{false, false, false},
	}

	fmt.Println(tic_tac_toe)
	fmt.Println(tic_tac_toe[0][0])

	for index, element := range tic_tac_toe {
		fmt.Printf("Element %v At Index %d\n", element, index)
	}

}
