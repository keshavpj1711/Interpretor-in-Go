package main 

import "fmt"

// Our task is to Write a program to compare two dates entered by user.
// Make a structure named Date to store the elements day, 
// month and year to store the dates. If the dates are equal, 
// display "Dates are equal" otherwise display "Dates are not equal".

type Date struct{
	day int
	month int
	year int
}

func main(){
	var date1, date2 Date

	// Taking input of two dates
	fmt.Printf("Enter date 1 in the format DD/MM/YYYY")
	fmt.Scanf("%d/%d/%d", &date1.day, &date1.month, &date1.year)

	fmt.Printf("Enter date 2 in the format DD/MM/YYYY")
	fmt.Scanf("%d/%d/%d", &date2.day, &date2.month, &date2.year)

	// Calling the compare Date Function
}