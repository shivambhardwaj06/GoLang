package main

import "fmt"
import "maps"
// maps -> hash , object , dictionary

func main() {
	// creating map
	m := make(map[string]string)

	// setting an elemetn
	m["name"] = "golang"
	m["area"] = "backend"
	// get an element
	fmt.Println("m['name']:", m["name"], "m['area']:", m["area"]) // golang
	// IMP : if key does not exist, it will return zero value of the type
	fmt.Println(m["phone"])
	
	n := make(map[string]int)
	n["age"] = 10
	fmt.Println("n['age']:", n["age"]) // 10
	fmt.Println(n["phone"]) // 0
	fmt.Println(len(n))

	delete(m, "area")
	fmt.Println("m['area']:", m["area"]) // ""
	fmt.Print(m)
	
	fmt.Println()

	// clear
	clear(n)
	fmt.Println(n==nil)

	// agar element pehle se pta hai toh
	s := map[string]int{"price":40, "quantity": 10}
	fmt.Printf("s: %v\n", s)

	V, ok := s["price"]
	fmt.Println("V:", V, "ok:", ok)
	if ok {
		fmt.Println("all ok")

	} else {
		fmt.Println("not ok")
	}


	//compare two maps 
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 1, "b": 2}
	m3 := map[string]int{"a": 1, "b": 3}

	// invalid way to compare two maps
	// fmt.Println("m1 == m2:", m1 == m2) // true
	// fmt.Println("m1 == m3:", m1 == m3) // false
	
	// valid way to compare two maps
	fmt.Println(maps.Equal(m1, m2)) // true
	fmt.Println(maps.Equal(m1, m3)) // false

}