package controller

import (
	"database/sql"
	"simple-graphql-server/internal/app/adapter/db"
	repo "simple-graphql-server/internal/app/adapter/repository"
	"simple-graphql-server/internal/app/domain/repository"
)

type Service struct {
	DB *sql.DB
}

func InitializeService() *Service {
	_db := db.Connect()
	return &Service{_db}
}

func (c *Service) NewAuthorRepository() repository.Author {
	return repo.NewAuthorRepository(c.DB)
}

func (c *Service) NewBookRepository() repository.Book {
	return repo.NewBookRepository(c.DB)
}
