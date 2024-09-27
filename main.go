// We need to import this always in GO
package main

// These are packages

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	//r.parseForm() This method is called on the http.Request object (r), and it parses the form data in the request. This includes both query parameters (e.g., in a URL like ?name=value) and the body of POST requests that contain form data.

	//After calling this method, you can access the parsed form data through r.Form, which is a map that holds the parsed values.

	// The statement if err := r.ParseForm(); err != nil checks if there was an error during the parsing process.

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful \n")
	// This line retrieves the value associated with the key "name" from the parsed form data. If the key doesn’t exist, it returns an empty string.
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address =%s\n", address)

}

//Passing a pointer instead of a value is generally more efficient, especially for large structs. When you pass a struct by value, Go makes a copy of it. For a struct like http.Request, which can contain a lot of data, copying it could be costly in terms of memory and performance. Using a pointer avoids this overhead.

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	// Golang provides another similar function Printf. The only difference between Frprintf and Printf is that Fprintf writes to the io.Writer instance passed to it while Printf function writes to the standard output

	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	//To register a specific function as the handler for requests to a given URL path
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting Server at Port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
