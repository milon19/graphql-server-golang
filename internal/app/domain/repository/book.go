package repository

import "simple-graphql-server/internal/app/domain"

type Book interface {
	List() ([]domain.Book, error)
	RelatedAuthors(ids []string) (map[string][]domain.Author, error)
}
