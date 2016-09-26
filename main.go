package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func add1(a, b int) int {
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

type result struct {
	Operand1 string
	Operand2 string
	Result   string
}

func homeHanlder(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("home_sample_1.html")
		t.Execute(w, "")
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	log.Println("Responsing to /add request")
	log.Println(r.UserAgent())

	vars := mux.Vars(r)
	oper1 := vars["oper1"]
	oper2 := vars["oper2"]
	fmt.Println(oper1, oper2)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello:")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/add/{oper1}/{oper2}", add).Methods("GET")
	router.HandleFunc("/", homeHanlder).Methods("GET")

	fmt.Println("Calculator Web Application started on port 9090")
	log.Fatal(http.ListenAndServe(":9090", router))
}
