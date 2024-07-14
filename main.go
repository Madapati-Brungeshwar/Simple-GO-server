package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./HTML")) // this will create a function that will serves file from the directory.
	http.Handle("/", fileServer)                      //http.Handle for the custom handler .Basically we can write custom logic in the custom handler. In the custom handler it should implement the common ServerHTTP method to satisfy handler interface.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	// If the url path is "/" then the Handle(HTTP handler mux) directs the request to appropriate handler function
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err) //this is best method for error handling, this will first logs and then exits the program withs exit status 1.

	}

	fmt.Println("lss")
}
