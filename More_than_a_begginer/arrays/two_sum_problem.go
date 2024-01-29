package main 

import "fmt"

// Creating a function to do our requiored task
func twoSum(nums []int, target int) []int {
	var output []int

	for i := 0; i < len(nums); i++{
		for j := i + 1; j < len(nums); j++{
			if nums[i] + nums[j] == target{
				output = append(output, i, j)
				return output
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

	fmt.Printf("Input array: %v\n",input_array) // Printing the input array 

	// Taking input for the target sum
	var targetSum int
	fmt.Printf("Enter the target sum: ")
	fmt.Scanf("%d\n", &targetSum)


	// Calling function to do the main computation 
	output := twoSum(input_array, targetSum)

	// Printing the output
	fmt.Printf("Output index required: %v", output)
}