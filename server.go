package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/merchants", createMerchantHandler)
	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/dbinit", dbInitHandler)
	http.HandleFunc("/dbseed", dbSeedHandler)
	log.Print(fmt.Sprintf("Listening on port %s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
