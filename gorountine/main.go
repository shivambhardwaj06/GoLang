package main

import (
	"fmt"
	"time"
)

func task(id int){
	fmt.Println("Doing task", id)
}

func main(){
	for i := 0; i <= 10; i++{
		// go task(i)
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 2)
}

// explain this behavorThe provided code demonstrates the use of goroutines in Go, which allows for concurrent execution of functions. In this example, the `task` function takes an integer `id` as an argument and prints a message indicating that it is performing a task with that ID.