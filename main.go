package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func mod(a, b int) int {
	return a % b
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // write data to response
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add handler")

	fmt.Println(r)
	fmt.Println("Inside homeHandler...")
	t, _ := template.ParseFiles("home.html")
	t.Execute(w, "")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add", addHandler)

	fmt.Println("Calculator Web Application started on port 9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("ERROR: ListenAndServe:", err)
	}

}
