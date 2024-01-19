package main

import "fmt"

func main(){

	// Info about fmt.Sprintf() 
	// fmt.Sprintf: This is a function from the fmt package that allows you
	// to format a string according to a format specifier and arguments.
	
	s := fmt.Sprintf("I am %v years old", 10)
	fmt.Printf(s)
	// I am 10 years old

	fmt.Printf("\n")

	k := fmt.Sprintf("I am %v years old", "way too many")
	fmt.Printf(k)
	// I am way too many years old

	// Also in GO we can't redeclare a variable i.e. 
	// one variable one declaration

}