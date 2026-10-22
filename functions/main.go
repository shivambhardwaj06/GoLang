package main
import "fmt"

func add(a int, b int) int {
	return a + b
}

// func add(a , b int) int {	
// 	return a + b
// }


func getLanguage() (string,string,string){
	return "Go", "Python", "Java"
}

func processIt(fn func(a int) int) int{
	return fn(1)
}

func funcasreturn() func(a int) int {
	return func(a int) int {
		return a + 1
	}
}

func main(){
	fn := func(a int) int {
		return a + 1
	}
	processIt(fn)
	println(funcasreturn()(1))
	println(processIt(fn))	
	lang1, lang2, _ := getLanguage()
	fmt.Println(lang1, lang2)
	fmt.Println(add(5, 3))
}

