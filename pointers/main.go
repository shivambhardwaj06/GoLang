package main

import "fmt"

func changeNum(num int) {
	num = 5
	fmt.Println("Inside changeNum:", num)	
}

func changeNumPointer(num *int) {
	*num = 5
	fmt.Println("Inside changeNumPointer:", *num)
}
func main(){
	num := 10
	fmt.Println("Before changeNum:", num)
	changeNum(num)
	fmt.Println("After changeNum:", num)

	println("Before changeNumPointer:", num)
	changeNumPointer(&num)
	fmt.Println("After changeNumPointer:", num)
}