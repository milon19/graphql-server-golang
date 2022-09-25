package repository

import (
	"database/sql"
	"fmt"
	"simple-graphql-server/internal/app/domain"
	"simple-graphql-server/internal/app/domain/repository"
	"strconv"
	"strings"
)

type Author struct {
	DB *sql.DB
}

func NewAuthorRepository(db *sql.DB) repository.Author {
	return &Author{DB: db}
}

func (a Author) ListByName(name string) ([]domain.Author, error) {
	query := fmt.Sprintf(`SELECT * FROM authors where lower(name) LIKE '%s%%'`, strings.ToLower(name))
	data, err := a.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var authors []domain.Author
	for data.Next() {
		var author domain.Author
		err = data.Scan(&author.ID, &author.Name, &author.Biography)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (a Author) RelatedBooks(ids []string) (map[string][]domain.Book, error) {
	query := fmt.Sprintf(`SELECT ("books_authors"."author_id") 
    							AS "author_id", "books"."id", "books"."title", "books"."price", "books"."isbn_no"
								FROM "books" INNER JOIN "books_authors" 
								ON ("books"."id" = "books_authors"."book_id") 
								WHERE "books_authors"."author_id" IN (%s)`, strings.Join(ids, ","))
	cursor, err := a.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var book domain.Book
	var bookId int
	authorBookMap := make(map[string][]domain.Book)
	for cursor.Next() {
		err = cursor.Scan(&bookId, &book.ID, &book.Title, &book.Price, &book.IsbnNo)
		if err != nil {
			return nil, err
		}
		authorBookMap[strconv.Itoa(bookId)] = append(authorBookMap[strconv.Itoa(bookId)], book)
	}
	return authorBookMap, nil
}
