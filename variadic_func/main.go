package main

import "fmt"

// closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
func counter() func() int {
	var count int = 0
	return func() int {
		count++
		return count
	}
}

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
// int -> interface{} but then cannot use + operator

func main() {
	nums := []int{1 ,3 ,5}
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(nums...))

	c := counter()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
}