package handlers

import (
	"net/http"

	"github.com/rs/cors"
)

func Cors(next http.Handler) http.Handler {

	c := cors.New(cors.Options{

		AllowedOrigins: []string{"https://rzhbadhon.github.io"},

		AllowedMethods: []string{"POST", "GET", "OPTIONS"},

		AllowedHeaders: []string{"Content-Type"},
	})

	return c.Handler(next)
}
