package main

import (
	"net/http"
	"example.com/hello-world/function"
)

func main() {
	http.HandleFunc("/", function.HelloWorld)
	http.ListenAndServe(":8080", nil)
} 