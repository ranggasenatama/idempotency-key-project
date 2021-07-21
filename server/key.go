package main

import "fmt"

func getDataKey(idempotencyKey string) string {
	return fmt.Sprintf("data:idempotencyKey:%s", idempotencyKey)
}

func getProgressKey(idempotencyKey string) string {
	return fmt.Sprintf("progress:idempotencyKey:%s", idempotencyKey)
}
