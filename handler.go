package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Hello string `json:"hello"`
	}

	resp := response {
		Hello: "world",
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(&resp)
}


func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "text/html; charset=utf8")
	fmt.Fprint(w, "ok")
}

func createMerchantHandler(w http.ResponseWriter, r *http.Request)  {
	// parse
	merchantReq := CreateMerchantRequest{}
	if err := json.NewDecoder(r.Body).Decode(&merchantReq); err != nil {
		fmt.Errorf("failed to decode %w", err)
		return
	}
	// validate

	// save(DB)
	merchant := Merchant{
		ID:       1,
		Email:    merchantReq.Email,
		Name:     merchantReq.Name,
	}

	// response
	merchantResp := CreateMerchantResponse{}
	merchantResp.Merchant = merchant
	merchantResp.Success = true

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(&merchantResp)
}
