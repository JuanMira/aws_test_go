package main

import (
	"net/http"
	"os"
)


func main() {
	http.HandleFunc("/", Greet)
	port := os.Getenv("PORT")

	if port == ""{
		port = "5000"
		
	}
	http.ListenAndServe(":"+port, nil)
}