package main

import (
	"io/ioutil"
	"os"

	"github.com/go-chi/chi"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	schema, err := parseSchema("./schema.graphql")
	if err != nil {
		panic(err)
	}
	graphQLHandler := &relay.Handler{Schema: schema}
	router := chi.NewRouter()
	router.Handle("/graphql", graphQLHandler)
}

func parseSchema(path string) (*graphql.Schema, error) {
	schemaBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	schema, err := graphql.ParseSchema(string(schemaBytes), &Resolver{})
	return schema, err
}

type Resolver struct {
}
