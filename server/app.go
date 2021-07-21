package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	initRedis()                                                                  // init redis
	http.HandleFunc("/idempotency-simulator/success", idempotency)               // endpoint success
	http.HandleFunc("/idempotency-simulator/in-progress", idempotencyInProgress) // endpoint in progress

	server := "127.0.0.1:8000"
	fmt.Printf("Starting server for testing HTTP ip: %s\n", server)
	if err := http.ListenAndServe(server, nil); err != nil {
		log.Fatal(err)
	}
}
