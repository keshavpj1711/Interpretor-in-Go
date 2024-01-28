package main 

import "fmt"

func main(){

	// Declaring and taking array as input
	var input_array []int

	for {
		var num int

		// Taking input
		fmt.Printf("Enter a number (or 0 to stop): ")
		fmt.Scanf("%d\n", &num)

		// Checking if input was zero
		if num == 0 {
			break
		}

		// Appending the input 
		input_array = append(input_array, num)
	}
}