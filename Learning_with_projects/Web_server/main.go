package main

import(
	"fmt"
	"log"
	"net/http"
)

func main(){
	fileServer := http.FileServer(http.Dir("./static"))  // We want to checkout static folder
	// GO, PHP all these lang by default look for index.html
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	
	fmt.Println("Starting server at port 8080")

	// Error Handling
	// Checking if the server starts or not
	if err := http.ListenAndServe("8080", nil); err != nil{
		log.Fatal(err)
	}

}