package main

import (
	"fmt"
	"log"
	"net/http"
	"rest"
)

func main() {

	fmt.Println("starting server on port 8080...")
	router := rest.InitHandlers()
	log.Fatal(http.ListenAndServe(":8080", router))

}
