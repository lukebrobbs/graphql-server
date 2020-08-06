package schema

import (
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/lukebrobbs/graphql-server/pkg/resolvers"
	"github.com/lukebrobbs/graphql-server/pkg/types"
)

func Queries() *graphql.Object {
	queries := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"people": &graphql.Field{
					Type: graphql.NewList((types.PersonType)),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return resolvers.GetPeople(), nil
					},
				},
				"person": &graphql.Field{
					Type: types.PersonType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Description: "The id of the person",
							Type:        graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						id, err := strconv.Atoi(p.Args["id"].(string))
						if err != nil {
							return nil, err
						}
						return resolvers.GetPersonByID(id), nil
					},
				},
			},
		},
	)
	return queries
}
