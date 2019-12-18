package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome to the Go web server run by a Docker image")
}

func main(){
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
