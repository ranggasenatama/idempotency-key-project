package main

import (
	"fmt"
	"net/http"
)

func idempotency(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/idempotency-simulator/success" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		idempotencyKey := r.FormValue("idempotency_key")
		request := Request{
			IdempotencyKey: idempotencyKey,
		}

		resp, err := core(request)
		if err != nil {
			resp.ErrorMessage = err.Error()
			resp.Write(w, r, http.StatusInternalServerError)
			return
		}

		resp.Write(w, r, http.StatusOK)
	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}

func idempotencyInProgress(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/idempotency-simulator/in-progress" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		idempotencyKey := r.FormValue("idempotency_key")
		request := Request{
			IdempotencyKey: idempotencyKey,
		}

		resp, err := coreInProgress(request)
		if err != nil {
			resp.ErrorMessage = err.Error()
			resp.Write(w, r, http.StatusInternalServerError)
			return
		}

		resp.Write(w, r, http.StatusOK)
	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}
