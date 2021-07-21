package main

type Request struct {
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}

type Response struct {
	Data         string `json:"data,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}
