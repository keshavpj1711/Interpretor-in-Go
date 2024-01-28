package main 

import "fmt"

// Creating a function to do our requiored task
func twoSum(nums []int, target int) []int {
	var output []int

	for i := 0; i < len(nums); i++{
		for j := i; j < len(nums); j++{
			if nums[i] + nums[j] == target{
				output = append(output, i, j)
			}
		}
	}

	return output
}

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