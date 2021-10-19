package main

type CreateMerchantRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type CreateMerchantResponse struct {
	Merchant Merchant `json:"merchant"`
	Success bool `json:"success"`
}
