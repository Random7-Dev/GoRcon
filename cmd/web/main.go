package main

import (
	"fmt"
	"net/http"

	"github.com/Random-7/GoCon/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println("Starting Web")
	_ = http.ListenAndServe(":8080", nil)
}
