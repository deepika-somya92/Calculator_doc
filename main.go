package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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

type result struct {
	Operand1 string
	Operand2 string
	Result   string
}

func homeHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello calculator!")
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	res := result{Operand1: "", Operand2: "", Result: ""}
	if r.Method == "GET" {
		t, _ := template.ParseFiles("home.html")
		t.Execute(w, res)
	} else {
		r.ParseForm()
		// Convert string to integer
		oper1, _ := strconv.Atoi(r.Form["operand1"][0])
		oper2, _ := strconv.Atoi(r.Form["operand2"][0])
		sum := add(oper1, oper2)
		// Convert integer to string
		res.Result = strconv.Itoa(sum)
		res.Operand1 = r.Form["operand1"][0]
		res.Operand2 = r.Form["operand2"][0]
		t, _ := template.ParseFiles("result.html")
		t.Execute(w, res)
	}
}

func subHandler(w http.ResponseWriter, r *http.Request) {
	type result struct {
		subOper1 string
		subOper2 string
		subRes   string
	}
	res := result{subOper1: "", subOper2: "", subRes: ""}
	if r.Method == "GET" {
		t, _ := template.ParseFiles("home.html")
		t.Execute(w, res)
	} else {
		r.ParseForm()
		// Convert string to integer
		oper1, _ := strconv.Atoi(r.Form["subOper1"][0])
		oper2, _ := strconv.Atoi(r.Form["subOper2"][0])
		sum := sub(oper1, oper2)
		// Convert integer to string
		res.subRes = strconv.Itoa(sum)
		res.subOper1 = r.Form["subOper1"][0]
		res.subOper2 = r.Form["subOper2"][0]
		t, _ := template.ParseFiles("result.html")
		t.Execute(w, res)
	}
}

func main() {
	http.HandleFunc("/", homeHanlder)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/sub", subHandler)

	fmt.Println("Calculator Web Application started on port 9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("ERROR: ListenAndServe:", err)
	}

}
