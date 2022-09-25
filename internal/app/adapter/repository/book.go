package repository

import (
	"database/sql"
	"fmt"
	"simple-graphql-server/internal/app/domain"
	"simple-graphql-server/internal/app/domain/repository"
	"strconv"
	"strings"
)

type Book struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) repository.Book {
	return &Book{DB: db}
}

func (b Book) List() ([]domain.Book, error) {
	query := `SELECT * FROM books`
	data, err := b.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var books []domain.Book
	for data.Next() {
		var book domain.Book
		err = data.Scan(&book.ID, &book.Title, &book.Price, &book.IsbnNo)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (b Book) RelatedAuthors(ids []string) (map[string][]domain.Author, error) {
	query := fmt.Sprintf(`SELECT ("books_authors"."book_id") 
    								AS "book_id", "authors"."id", "authors"."name", "authors"."biography"
									FROM "authors" INNER JOIN "books_authors" 
									ON ("authors"."id" = "books_authors"."author_id") 
									WHERE "books_authors"."book_id" IN (%s)`, strings.Join(ids, ","))
	cursor, err := b.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var author domain.Author
	var authorId int
	bookAuthorMap := make(map[string][]domain.Author)
	for cursor.Next() {
		err = cursor.Scan(&authorId, &author.ID, &author.Name, &author.Biography)
		if err != nil {
			return nil, err
		}
		bookAuthorMap[strconv.Itoa(authorId)] = append(bookAuthorMap[strconv.Itoa(authorId)], author)
	}
	return bookAuthorMap, nil
}
