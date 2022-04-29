package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := chi.NewRouter()

	r.HandleFunc("/", helloHandler)
	r.Post("/merchants", createMerchantHandler)
	// 本当は /merchants にしたいが、MethodPOSTの分岐を書くのがめんどうだったのであとでリファクタをする
	r.Get("/merchants", listMerchantHandler)
	r.Get("/merchants/{merchantID}", getMerchantHandler)
	r.Get("/health", healthCheckHandler)
	r.Get("/dbinit", dbInitHandler)
	r.Get("/dbseed", dbSeedHandler)
	log.Print(fmt.Sprintf("Listening on port %s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
