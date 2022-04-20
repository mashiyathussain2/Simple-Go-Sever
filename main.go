package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func formHanlder(w http.ResponseWriter, r *http.Request) {
	// For all requests, ParseForm parses the raw query from the URL and updates r. Form. For POST, PUT, and PATCH requests, it also parses the request body as a form and puts the results into both r. PostForm and r. Form.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not allowed", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func forms(w http.ResponseWriter, r *http.Request) {
	// generating HTML output using template package
	mytemp, err := template.ParseFiles("./static/form.html")
	if err != nil {
		return
	}
	err = mytemp.Execute(w, nil)
	if err != nil {
		return
	}
}

func main() {
	// here we tell the go to checkout the static folder
	fileserver := http.FileServer(http.Dir("./static"))

	//hanlde function inside the http package
	// handling the / route
	http.Handle("/", fileserver) // send that to fileserver it will server index.html file
	http.HandleFunc("/formss", forms)

	http.HandleFunc("/form", formHanlder) // handle the form route
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil { // creating the server
		log.Fatal(err)
	}

}
