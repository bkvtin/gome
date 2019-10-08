package main

import (
	"fmt"
	"net/http"

	"github.com/bkvtin/gome/simplehttp/handlers"
)

func main() {
	var port string
	port = ":8080"

	http.HandleFunc("/api/v1/health", handlers.CheckHealth)
	http.HandleFunc("/api/v1/getstds", handlers.GetStudents)

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	} else {
		fmt.Println("Server listent at", port)
	}

}
