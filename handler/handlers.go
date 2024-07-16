// Copyright 2024 qishenonly. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// MIT License: https://opensource.org/licenses/MIT

package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/qishenonly/Shorten/storage"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randSequnce generates a random string of length n
// from the letters slice
func randSequnce(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// ShortenURLHandler handles the shorten URL request
func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "bad request : "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Unmarshal the request body
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "bad request : "+err.Error(), http.StatusBadRequest)
		return
	}

	// Generate a random short URL
	shortURL := randSequnce(6)

	// Save the short URL and original URL mapping
	err = storage.SaveURLMapping(shortURL, req.URL)
	if err != nil {
		http.Error(w, "failed to save URL mapping : "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal the response
	resp := ShortenResponse{ShortURL: shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func GetShortenURLInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is GetShortenURLInfoHandler")
}

// RedirectHandler handles the redirect request
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Get the short URL from the request path
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	// Get the original URL
	originalURL, err := storage.GetOriginalURL(shortURL)
	if err != nil {
		http.Error(w, "short URL not found : "+err.Error(), http.StatusNotFound)
		return
	}

	// Redirect to the original URL
	http.Redirect(w, r, originalURL, http.StatusFound)
}
