package main

import "fmt"

type Person struct {
    Name string
    Age  int
    City string
}

// constructor hack in go
func newPerson (Name string, Age int , City string) *Person {
	    p := Person{
        Name: Name,
        Age:  Age,
        City: City,
    }
	return &p
}


func (p * Person) changeAge(newAge int) {	
	p.Age = newAge // humne yaha * nhi diya kyuki struct as default pointer pass hota hai
}

	// struct embedding
	type Employee struct {
		Person
		EmployeeID string
		salary     float64
	}

func main() {
    // Create a struct variable
    p1 := newPerson("Rahul", 22, "Dehradun")
	p1.changeAge(40)

    // Access struct fields
    fmt.Println(p1.Name)
    fmt.Println(p1.Age)
    fmt.Println(p1.City)

	// inline struct
	language := struct {
		name string
		isGood bool
	} {
		name: "Go",
		isGood: true,
	}
	fmt.Println(language)

	employee := Employee{
		Person: Person{
			Name: "Alice",
			Age:  30,
			City: "New York",
		},
		EmployeeID: "E123",
		salary:     50000.0,
	}
	employee.Person.Name="Shivam"
	fmt.Println(employee)

}

// if you don't set any field , default value is zero value

