package main

import "fmt"

func main(){
	firstName, _ := getNames() 
	// This is how we can ignore the multipe returns
	// Like here we did not require lastname so we didn't call for it
	
	// Welcome message
	fmt.Println("Welcome to TEXTIO", firstName)
}

func getNames()(string, string){
	var firstName, lastName string
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	
	return firstName, lastName
}