package main

import "log"

// core function
func doSomething(request Request) Response {
	log.Printf("do something here")
	return Response{
		Data: request.IdempotencyKey,
	}
}
