package schema

import (
	"github.com/graphql-go/graphql"
	gql "simple-graphql-server/internal/app/adapter/graphql"
	repo "simple-graphql-server/internal/app/domain/repository"
)

type Book struct {
	Repo repo.Book
}

func (b Book) BookSchema() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(gql.BookType()),
		Description: "Get book list",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return b.Repo.List()
		},
	}
}
