package main 
import "fmt"

func printSlice(items []int){
	for _, item := range items{
		fmt.Println(item)
	}
}

func printSlice2(items []string){
	for _, item := range items{
		fmt.Println(item)
	}
}
//LIFO
type Stack[T int | string] struct {
	elements []T
}


func allPrintSlice[T int | string](items []T){ // any == interface{}
	for _, item := range items{
		fmt.Println(item)
	}
}
// multiple
func allPrintSlice2[T comparable , V string](items []T, name V){
	for _, item := range items{
		fmt.Println(item, name)
	}
} // any == interface{}
func main(){
	// nums := []int{1,2,3,4,5}

	names := []string{"Rahul", "Rohit", "Amit", "Saurabh"}	
	printSlice([]int{1,2,3,4,5})
	printSlice2(names)
	// vals := []bool{true,false}
	allPrintSlice([]int{1,2,3,4,5})
	allPrintSlice(names)
	// allPrintSlice(vals) give errror

	stack := Stack[int]{
		elements: []int{1,2,3,4,5},
		// elements: []string{"Rahul", "Rohit", "Amit", "Saurabh"}, // this will give error because we have defined elements as int slice
	}
	stack2 := Stack[string]{
		elements: []string{"Rahul", "Rohit", "Amit", "Saurabh"},
	}

	fmt.Println(stack.elements)
	fmt.Println(stack2.elements)

	allPrintSlice2([]int{1,2,3,4,5}, "Rahul")


}