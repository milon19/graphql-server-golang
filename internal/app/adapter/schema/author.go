package schema

import (
	"github.com/graphql-go/graphql"
	gql "simple-graphql-server/internal/app/adapter/graphql"
	repo "simple-graphql-server/internal/app/domain/repository"
)

type Author struct {
	Repo repo.Author
}

func (a Author) AuthorSchema() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(gql.AuthorType()),
		Description: "Get author by name",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var name string
			if val, ok := p.Args["name"]; ok {
				name = val.(string)
			}
			return a.Repo.ListByName(name)
		},
	}
}
