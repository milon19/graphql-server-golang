package repository

import "simple-graphql-server/internal/app/domain"

type Author interface {
	ListByName(name string) ([]domain.Author, error)
	RelatedBooks(ids []string) (map[string][]domain.Book, error)
}
