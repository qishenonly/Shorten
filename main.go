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

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/qishenonly/Shorten/handler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/shorten", handler.ShortenURLHandler).Methods("POST")
	r.HandleFunc("/api/info", handler.GetShortenURLInfoHandler).Methods("GET")
	r.HandleFunc("/api/{shortURL}", handler.RedirectHandler).Methods("GET")

	log.Println("Shorten server started at :3050")
	log.Fatal(http.ListenAndServe(":3050", r))
}
