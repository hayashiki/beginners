package main

type CreateMerchantRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	PhotoURL string `json:"photo_url"`
	Password string `json:"password"`
}

type CreateMerchantResponse struct {
	Merchant Merchant `json:"merchant"`
	Success bool `json:"success"`
}
