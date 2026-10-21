package main

// numbered sequence of specific length

import "fmt"

func main() {
	var nums [5]int
	// fmt.Println("nums:", nums) [0,0,0,0,0]
	// fmt.Println("length of nums:", len(nums))
	nums[0] = 1
	fmt.Println("nums[0]:", nums[0]) // [1,0,0,0,0]

	var vals [4]bool
	vals[2]	= true
	fmt.Println("vals:", vals) // [false,false,false,false]

	var names [3]string
	names[0] = "Alice"
	fmt.Println("names:", names) // [Alice, , ]

	// To declare and initialize an array in a single line, you can use the following syntax:
	arr := [3]int{1, 2, 3}
	fmt.Println("arr:", arr) // [1,2,3]
	
	// You can also use the ellipsis (...) to let the compiler determine the length of the array based on the number of elements provided:

	// arr2 := [...]int{4, 5, 6, 7}
	 
	// 2d array
	nums2D := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("nums2D:", nums2D)

	// in slice we can use append to add elements but in array we cannot use append to add elements because array has fixed size
	//slice is a dynamically-sized, flexible view into the elements of an array. In contrast, an array has a fixed size and cannot be resized. Slices are more commonly used in Go because they provide more flexibility and functionality than arrays.
	// array provide constant time access to elements, while slices provide linear time access to elements. This means that accessing an element in an array is faster than accessing an element in a slice, especially for large arrays or slices.
}