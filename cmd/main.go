package main

import (
	"SMSApp/internal/smsapp"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	r := smsapp.GetRouter()

	fmt.Println("starting server...")
	server := &http.Server{
		Handler:     r,
		Addr:        "127.0.0.1:8080",
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
