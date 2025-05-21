package main

import (
	"fmt"
	"net/http"
	"log"
)

func demo_response(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world - go server")
}

func main() {
	http.HandleFunc("/",demo_response)
	log.Fatal(http.ListenAndServe(":8081",nil))
}