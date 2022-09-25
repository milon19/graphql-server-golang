package schema

import (
	"github.com/graphql-go/graphql"
	"log"
	repo "simple-graphql-server/internal/app/domain/repository"
)

type Schema struct {
	BookRepo   repo.Book
	AuthorRepo repo.Author
}

func (s Schema) GetSchema() graphql.Schema {
	fields := graphql.Fields{
		"books":   Book{Repo: s.BookRepo}.BookSchema(),
		"authors": Author{Repo: s.AuthorRepo}.AuthorSchema(),
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return schema
}
