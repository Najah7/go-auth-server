package main

import (
	"fmt"
	"net/http"
	_ "github.com/Najah7/go-auth-api/models"
)

func main() {
	http.HandlerFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!")
	})

	http.ListenAndServe(":8080", nil)

}
