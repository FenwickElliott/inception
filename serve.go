package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	port := flag.String("port", "3032", "port")
	fmt.Println("Serving on port:", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
