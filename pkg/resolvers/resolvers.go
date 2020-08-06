package resolvers

import "github.com/lukebrobbs/graphql-server/pkg/schema"

type resolvers struct {
}

var people = []schema.Person{
	{
		ID:   1,
		Name: "Luke",
	},
}

func New() schema.Resolvers {
	return &resolvers{}
}

func (r *resolvers) GetPeople() []schema.Person {
	return people
}

func (r *resolvers) GetPersonByID(id int) schema.Person {

	for _, p := range people {
		if p.ID == id {
			return p
		}
	}
	return schema.Person{}
}
