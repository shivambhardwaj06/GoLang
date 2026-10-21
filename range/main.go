package main 
import "fmt"
// iterating over data structures 
func main(){
	nums := []int {6,7,8}
	for i:=0; i<len(nums); i++{	
		fmt.Println(nums[i])
	}

	for _, v := range nums{
		fmt.Println(v)
	}
	sum := 0
	for i, num := range nums{
		fmt.Println("i:", i, "v:", num)
		sum = sum + num
	}
	println("sum:", sum)

	m := map[string]string{"fname":"John", "lname":"Doe"}
	for k, v := range m{
		fmt.Println(k, ":", v)
	}

	// unicode code point rune
	// i=starting byte of rune
	// 255 -> 1 byte // greater than 255 char -> 2 bytes
	for i, c := range "gΩlang"{
		fmt.Println(i, c)// prints the unicode code point of the character
		fmt.Println(i, string(c))// prints the character
	}

}