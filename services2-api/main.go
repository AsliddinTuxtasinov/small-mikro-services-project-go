package main

import (
	"fmt"
	"log"
	"net/http"
	"services2-api/middleware"
)

func main() {
	port := "9001"
	fmt.Printf("Server have started at :%v port !\n", port)
	handleRequests(port)
}

func handleRequests(port string) {
	http.Handle("/", middleware.IsAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Secret Information")
}
