package main

import (
	"context"
	"encoding/json"

	redigo "github.com/gomodule/redigo/redis"
)

/* Section Core Data */
func GetProductCampaignStock(ctx context.Context, request Request) (Response, error) {
	response := Response{}

	key := getDataKey(request.IdempotencyKey)
	res, err := client.Get(key).Result()
	if err != nil {
		return response, err
	}

	if res == "" {
		return response, redigo.ErrNil
	}

	err = json.Unmarshal([]byte(res), &response)
	return response, err
}

func SetProductCampaignStock(ctx context.Context, request Request, data Response) error {
	key := getDataKey(request.IdempotencyKey)
	strData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = client.SetNX(key, strData, ttlCache).Err()
	return err
}

/* Section Progress */
func GetProductCampaignStockProgress(ctx context.Context, request Request) (bool, error) {
	key := getProgressKey(request.IdempotencyKey)
	res, err := client.Get(key).Result()
	if err != nil {
		return false, err
	}

	if res == "" {
		return false, redigo.ErrNil
	}

	if res == "1" {
		return true, nil
	}

	return false, nil
}

func SetProductCampaignStockProgress(ctx context.Context, request Request) error {
	key := getProgressKey(request.IdempotencyKey)
	raw := "1"

	err := client.SetNX(key, raw, ttlProgress).Err()
	return err
}

func DelProductCampaignStockProgress(ctx context.Context, request Request) error {
	key := getProgressKey(request.IdempotencyKey)

	err := client.Del(key).Err()
	return err
}
