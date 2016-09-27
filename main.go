package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func homeHanlder(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("home.html")
		t.Execute(w, "")
	}
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	o1, _ := strconv.Atoi(vars["oper1"])
	o2, _ := strconv.Atoi(vars["oper2"])
	result := add(o1, o2)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, result)
}

func subHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	o1, _ := strconv.Atoi(vars["oper1"])
	o2, _ := strconv.Atoi(vars["oper2"])
	result := sub(o1, o2)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, result)
}

func mulHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	o1, _ := strconv.Atoi(vars["oper1"])
	o2, _ := strconv.Atoi(vars["oper2"])
	result := mul(o1, o2)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, result)
}

func divHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	o1, _ := strconv.Atoi(vars["oper1"])
	o2, _ := strconv.Atoi(vars["oper2"])
	result := div(o1, o2)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, result)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeHanlder).Methods("GET")
	router.HandleFunc("/add/{oper1}/{oper2}", addHandler).Methods("GET")
	router.HandleFunc("/sub/{oper1}/{oper2}", subHandler).Methods("GET")
	router.HandleFunc("/mul/{oper1}/{oper2}", mulHandler).Methods("GET")
	router.HandleFunc("/div/{oper1}/{oper2}", divHandler).Methods("GET")

	fmt.Println("Calculator Web Application started on port 9090")
	log.Fatal(http.ListenAndServe(":9090", router))
}
