package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))

	http.Handle("/", http.FileServer(http.Dir("./public")))

	port := "3031"

	// http.HandleFunc("/insert", insert)
	// http.HandleFunc("/find", find)
	// http.HandleFunc("/remove", remove)
	fmt.Println("Serving on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
