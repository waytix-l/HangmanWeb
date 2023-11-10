package main

import (
	"fmt"
	"net/http"

	"github.com/waytix-l/HangmanWeb/internal/handlers"
)

const port = ":8877"

func main() {
	
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/contact", handlers.Contact)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("(htpp://localhost:8877) - Server started on port ", port)
	http.ListenAndServe(port, nil)
}


