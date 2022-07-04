package main

import (
	"fmt"
	"net/http"

	"github.com/Random-7/GoRcon/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println("Starting Web...")
	_ = http.ListenAndServe(portNumber, nil)
}
