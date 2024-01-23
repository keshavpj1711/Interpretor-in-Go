package main

import(
	"fmt"
	"log"
	"net/http"
)

//For any API call we have a response and request and this is similar to what we have
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello"{
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	
	fmt.Fprintf(w, "Post Request Successful")
	
	name := r.FormValue("name")
	address := r.FormValue("address")
	
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

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