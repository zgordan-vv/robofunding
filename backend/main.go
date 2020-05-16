package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/graph-gophers/graphql-go"

	//"github.com/graph-gophers/graphql-go/relay"
	"github.com/zgordan-vv/robofunding/backend/auth"
	"github.com/zgordan-vv/robofunding/backend/resolvers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	schema, err := parseSchema("./schema.graphql")
	must(err)
	//graphQLHandler := &relay.Handler{Schema: schema}
	graphQLHandler := &handler{Schema: schema}
	router := chi.NewRouter()
	router.Use(auth.Middleware())
	router.Handle("/graphql", graphQLHandler)
	must(http.ListenAndServe(":"+port, router))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func parseSchema(path string) (*graphql.Schema, error) {
	schemaBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	schema, err := graphql.ParseSchema(string(schemaBytes), resolvers.NewResolver())
	return schema, err
}
