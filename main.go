package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/lukebrobbs/graphql-server/pkg/schema"
)

func main() {

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: schema.Queries(),
	})
	if err != nil {
		// panic if there is an error in schema
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	// static file server to serve Graphiql in-browser editor
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// graphql api  server
	http.Handle("/graphql", h)

	fmt.Println("server is started at: http://localhost:8080/")
	fmt.Println("graphql api server is started at: http://localhost:8080/graphql")
	http.ListenAndServe(":8080", nil)
}
