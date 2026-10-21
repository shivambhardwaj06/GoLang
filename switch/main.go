package main

import (
	"fmt"
	"time"
)


func main(){
	// simple switch
	// switch is used to select one of many code blocks to be executed
	// break statement is not required in go switch case

	i := 5
	switch i {
	case 1 :
		println("one")	
	case 2 :
		println("two")
	case 3 :
		println("three")
	}

	// multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("It's the weekend!")
	default:
		println("It's a weekday.")	}
	

	// double quotes are used for string literals and back quotes are used for raw string literals
	//type switch
	// whoAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case int:
	// 		println("It's an int")
	// 	case string:
	// 		println("It's a string")
	// 	default:
	// 		println("It's something else", t)
	// 		fmt.Println("It's something else ", t)
	// 	}
	// }
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			println("It's an int")
		case string:
			println("It's a string")
		default:
			println("It's something else", i)
			fmt.Println("It's something else ", i)
		}
	}
	whoAmI(1)
	whoAmI("hello")
	whoAmI(true)

}