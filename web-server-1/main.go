package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func formHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/form" {
		http.Error(res, "Not Found", http.StatusNotFound)
		return
	}
	if req.Method != "POST" {
		http.Error(res, "Method not supported", http.StatusNotFound)
		return
	}
	err := req.ParseForm()
	if err != nil {
		fmt.Fprintf(res, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(res, "Success")
	name := req.FormValue("name")
	address := req.FormValue("address")
	fmt.Fprintf(res, "Name = %s\n", name)
	fmt.Fprintf(res, "Address = %s\n", address)
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(res, "Not Found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(res, "Method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(res, "Hello")
}
