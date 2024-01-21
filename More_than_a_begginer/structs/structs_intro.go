package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message string
}

type cars struct{
	modelName string
	modelNo string
	height float64  // In meters
	width float64  // In meters
}

// This function is to print
func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("====================================")
}

func main() {
	test(messageToSend{
		phoneNumber: 148255510981,
		message:     "Thanks for signing up",
	})
	test(messageToSend{
		phoneNumber: 148255510982,
		message:     "Love to have you aboard!",
	})
	test(messageToSend{
		phoneNumber: 148255510983,
		message:     "We're so excited to have you",
	})

	// Storing data of car variable
	var mycar cars
	mycar.modelName = "NissanGT"
	mycar.modelNo = "129762"
	mycar.height = 1.72
	mycar.width = 1.5

	// Printing the specs of my car
	fmt.Printf("My %s(%s) has a height of %f and a width of %f", mycar.modelName, mycar.modelNo, mycar.height, mycar.width)
}
