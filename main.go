package main

import (
	"net/http"
	"os"

	"github.com/JuanMira/aws_api/controllers"
)


func main() {
	http.HandleFunc("/", controllers.Greet)
	port := os.Getenv("PORT")

	if port == ""{
		port = "5000"

	}
	http.ListenAndServe(":"+port, nil)
}