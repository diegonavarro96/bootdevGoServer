package main

import (
	"fmt"
	"net/http"
)

func main() {
	const port = "8080"
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.Handle("/", http.FileServer((http.Dir("."))))
	mux.Handle("/assets", http.FileServer((http.Dir("./logo.png"))))

	fmt.Println("Hello world")
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Server connection was not made Panic!")
	} else {
		fmt.Println("Server on port : ", port)
	}

}
