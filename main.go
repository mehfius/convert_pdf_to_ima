package main

import (
	"net/http"
	"github.com/mehfius/convert_pdf_to_ima/function"
)

func main() {
	http.HandleFunc("/", function.HelloWorld)
	http.ListenAndServe(":8080", nil)
} 