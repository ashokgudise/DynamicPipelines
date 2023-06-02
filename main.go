package main

import (
	"dyna-pod-pipeline/rest"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/process", rest.ProcessRequest)

	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	
}
