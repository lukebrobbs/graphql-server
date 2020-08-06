package resolvers

type Person struct {
	ID   int
	Name string
}

var people = []Person{
	{
		ID:   1,
		Name: "Luke",
	},
}

func GetPeople() []Person {
	return people
}

func GetPersonByID(id int) Person {

	for _, p := range people {
		if p.ID == id {
			return p
		}
	}
	return Person{}
}
