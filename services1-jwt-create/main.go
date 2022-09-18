package main

import (
	"fmt"
	"log"
	"net/http"
	"services1-jwt-create/helpper"
)

func main() {
	port := "8080"
	fmt.Printf("Server have started at :%v port !\n", port)
	handleRequests(port)
}

func handleRequests(port string) {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	validToken, err := helpper.GetJwt()
	if err != nil {
		log.Fatalf("Something went wrong: %s\n", err.Error())
	}
	fmt.Println("Valid token -> ", validToken)
	fmt.Fprintf(w, string(validToken))
}
