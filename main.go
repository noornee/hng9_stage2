package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	mux.HandleFunc("/", arithmeticHandler)

	fmt.Printf("Server running on port %s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}

}
