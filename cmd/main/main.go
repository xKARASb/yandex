package main

import (
	"calculate/internal/transport"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/calculate", transport.CalculateHandler)

	http.ListenAndServe(":8080", nil)
}
