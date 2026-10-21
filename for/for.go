package main

// for -> only construct in go for looping
func main(){
	// while loop
	i := 1
	for i <= 3 {
		println(i)
		i=i+1
	}
	// infinte loop

	// for {
	// 	println("1")
	// }

	// classic for loop
	for i:=0;i<3;i++{

		if i == 2 {
			continue
		}

		println(i)

		//break 
		// continue
	}
	for i := range 11{
		println(i)
	}
	

}