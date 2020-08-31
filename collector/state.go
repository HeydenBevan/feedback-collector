package main

import "net/http"

type common struct {
	Keys   map[string]string
	Client *http.Client
}
