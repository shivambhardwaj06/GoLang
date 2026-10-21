package main

func main() {

	age := 16

	if age >= 18 {
		println("person is an adult")
	}else {
		println("Person is less than 18 ")
	}

	if age >= 18 {
		println("person is an adult")
	}else if age >= 12{
		println("Person is teenager ")
	}else {
		println("Person is a kid")
	}

	// can use == , != , || , && operator also

	if age := 15; age >= 18 {
		println("person is an adult",age)
	}else if age >= 12 {
		println("person is teenager",age)
	}

	// go does not have ternary operator 

}