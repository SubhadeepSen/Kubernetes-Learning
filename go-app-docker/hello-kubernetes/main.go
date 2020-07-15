package main

import (
    "fmt"
    "net/http"
)

func sayHello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Hello, Welcome to Kubernetes\n")
}

func main() {
    http.HandleFunc("/", sayHello)
	fmt.Println("Server listening at localhost:8080")
    http.ListenAndServe(":8080", nil)
}