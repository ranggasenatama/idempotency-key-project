package main

import (
	"context"
	"log"
)

func core(request Request) (Response, error) {
	resp, progress, err := ValidateIdempotency(context.Background(), request)
	if err != nil {
		log.Printf("[core] error validate idemptency key - key: %+v - err: %+v", request.IdempotencyKey, err)
		return resp, err
	}

	if !progress {
		log.Printf("[core] New idemptency key start - key: %+v", request.IdempotencyKey)
		// asd
		err = SetProductCampaignStockProgress(context.Background(), request)
		if err != nil {
			return resp, err
		}

		defer DelProductCampaignStockProgress(context.Background(), request)

		resp = doSomething(request) // your core logic put here

		err = SetProductCampaignStock(context.Background(), request, resp)
		if err != nil {
			return resp, err
		}
		log.Printf("[core] New idemptency key - value: %+v", resp)
	} else {
		log.Printf("[core] Exist idemptency key - key: %+v", request.IdempotencyKey)
		log.Printf("[core] Exist idemptency key - value: %+v", resp)
	}
	return resp, err
}
