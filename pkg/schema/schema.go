package schema

import (
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/lukebrobbs/graphql-server/pkg/types"
)

type schema struct {
	r Resolvers
}

type Schema interface {
	Queries() *graphql.Object
}

type Resolvers interface {
	GetPeople() []Person
	GetPersonByID(id int) Person
}
type Person struct {
	ID   int
	Name string
}

func New(r Resolvers) (graphql.Schema, error) {
	s := &schema{
		r: r,
	}
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: s.Queries(),
	})
}

func (s *schema) Queries() *graphql.Object {
	queries := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"people": &graphql.Field{
					Type: graphql.NewList((types.PersonType)),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return s.r.GetPeople(), nil
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
						return s.r.GetPersonByID(id), nil
					},
				},
			},
		},
	)
	return queries
}
