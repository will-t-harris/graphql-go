package main

import (
	"net/http"

	"github.com/graphql-go/handler"
)

func main() {
	h := handler.New(&handler.Config{
		Schema:   &BeastSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)

	http.ListenAndServe(":8080", nil)
}
