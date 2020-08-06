package types

import "github.com/graphql-go/graphql"

var PersonType = graphql.NewObject(graphql.ObjectConfig{

	Name:        "Person",
	Description: "A person",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Id of the person",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "The name of the person",
		},
	},
})
