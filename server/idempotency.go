package main

import (
	"context"
	"errors"

	redigo "github.com/go-redis/redis"
)

func ValidateIdempotency(ctx context.Context, request Request) (Response, bool, error) {
	resp, progress, err := ValidateIdempotencyCore(ctx, request)
	if err != nil && err != redigo.Nil {
		return resp, progress, err
	}

	if progress {
		return resp, progress, err
	}

	return resp, false, nil
}

// ValidateIdempotencyCore : if this function return true, it means the data is exist or the progress still running
func ValidateIdempotencyCore(ctx context.Context, request Request) (Response, bool, error) {
	response := Response{}

	//TODO: validate still progress or not
	nilErr := redigo.Nil
	progress, err := GetProductCampaignStockProgress(ctx, request)
	if err != nil {
		if err.Error() != nilErr.Error() {
			return response, false, err
		}
	}

	// it means th validation still in progress
	if progress {
		return response, true, errors.New("still in progress")
	}

	//TODO: get data redis exist or not
	response, err = GetProductCampaignStock(ctx, request)
	if err != nil {
		return response, false, err
	}

	return response, true, nil
}
