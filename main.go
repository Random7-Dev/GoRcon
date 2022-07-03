package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println("Starting Web")
	_ = http.ListenAndServe(":8080", nil)
}
