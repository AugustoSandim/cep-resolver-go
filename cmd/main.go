package main

import (
	"fmt"
	"net/http"

	"cep-resolver-go/internal/handlers"
)

func main() {
	http.HandleFunc("/address", handlers.GetAddress)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
