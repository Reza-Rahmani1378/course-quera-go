package main

import "fmt"

func main() {

	/* Slicing Array with This Way
	s := arr[start:end]
	*/
	var colors = [5]string{
		"Red",
		"Blue",
		"Green",
		"Pink",
		"Purple",
	}
	s1 := colors[1:3] // Selected Indices : 1,2

	s2 := colors[:3] // Selected Indices : 0,1,2

	s3 := colors[2:] // Selected Indices : 2,3,4

	s4 := colors[:] // Selected Indices : 0,1,2,3,4

	fmt.Println("Array:", colors)
	fmt.Println("Slice 1:", s1)
	fmt.Println("Slice 2:", s2)
	fmt.Println("Slice 3:", s3)
	fmt.Println("Slice 4:", s4)

	/*
		By means of this function, memory is allocated to the slice and initialization of the slice is done.
		make([]T,len) T --> Type of Data
		len --> Length Of Slice,
	*/

	s5 := make([]string, 0)
	fmt.Println(s5)

	s6 := make([]int, 3)
	fmt.Println(s6)

	// Initialize Slice
	var lessons = []string{"Math", "Physics"}

	fmt.Println(lessons)

	// Iteration In Slice

	letters := []string{"a", "b", "c", "d"}

	for _, letter := range letters {
		fmt.Printf("%s\t", letter)
	}
	fmt.Printf("\n")

	for i := 0; i < len(letters); i++ {
		fmt.Printf("%s\t", letters[i])
	}

	fmt.Printf("\n")

	// Function append(...)

	letters = append(letters, "e")

	fmt.Println("Slice:", letters)

	// Getting Length and Capacity of a Slice with methods len() & cap()
	/*
			length indicates the number of elements in the slice, and capacity indicates
			the number of elements in the array that the slice points to. (Starting from
		    the index of the array that the slice points to.) When the length of the slice
			is equal to its capacity and an element In addition, another place of memory with
		    twice the current capacity is allocated and the elements are copied there.
	*/

	// test := []string{}
	metals := make([]string, 0)
	/* metals := []string{"gold", "silver"} */
	fmt.Printf("len : %d , cap : %d\n", len(metals), cap(metals))

	/* Output :
	   len : 2 , cap : 2
	*/

	metals = append(metals, "iron", "test", "reza", "rahmani", "rezzzzz")
	fmt.Printf("len : %d , cap : %d\n", len(metals), cap(metals))
	/* Output :
	   len : 3 , cap : 4
	*/

	// appending two slices with this way
	even := []int{2, 4, 6}
	odd := []int{1, 3, 5}

	numbers := append(odd, even...)

	fmt.Println(numbers)

}
