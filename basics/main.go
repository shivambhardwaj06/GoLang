package main 

import "fmt"

func main(){
	fmt.Println("Hello World")

	// simple values

	fmt.Println(+41)

	// string ;
	fmt.Println(("hello golang"))

	// bool 
	fmt.Println(true)
	fmt.Println(false)

	// floats
	fmt.Println(10.5)

	fmt.Println(3.0/2.0)

	// let's try Print
	fmt.Print(454)
	fmt.Print("Hello World")
	fmt.Println()
	fmt.Print("hello golang\n")


	// Variable in Go - must be used in program
	var name string = "golang"
	fmt.Println(name);
	var name2 = "shivam"
	fmt.Println(name2)

	// shorthand syntax

	name3 := "golang"
	fmt.Println(name3);

	// value baad mai assign karni hai 
	var name5 string
	name5 = "golang"
	println(name5)

	var price float32 = 50.55
	price2 := 50.55
	var price3 = 50.55
	println(price)
	println(price2)
	println(price3)

	// constants 



}