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

package storage

import (
	"errors"
	"sync"
)

// Memory storage for URL mapping
var (
	urlStore = make(map[string]string)
	mu       sync.RWMutex
)

// SaveURLMapping Save the mapping relationship between
// the short link and the original URL
func SaveURLMapping(shortURL, originalURL string) error {
	mu.Lock()
	defer mu.Unlock()

	// Check if the short URL already exists
	if _, exists := urlStore[shortURL]; exists {
		return errors.New("short URL already exists")
	}
	urlStore[shortURL] = originalURL
	return nil
}

// GetOriginalURL Get the original URL corresponding to the short URL
func GetOriginalURL(shortURL string) (string, error) {
	mu.RLock()
	defer mu.RUnlock()

	// Get the original URL
	originalURL, exists := urlStore[shortURL]
	if !exists {
		return "", errors.New("short URL not found")
	}
	return originalURL, nil
}
