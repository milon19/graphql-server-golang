package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"log"
	"simple-graphql-server/internal/app/domain"
	"simple-graphql-server/internal/app/domain/repository"
)

type Book struct {
	Repo repository.Author
}

func (b Book) GetBooksBatchFn(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	handleError := func(err error) []*dataloader.Result {
		var results []*dataloader.Result
		var result dataloader.Result
		result.Error = err
		log.Println(err)
		results = append(results, &result)
		return results
	}
	var authorIds []string
	var results []*dataloader.Result
	for _, key := range keys {
		authorIds = append(authorIds, key.String())
	}
	authorBookMap, err := b.Repo.RelatedBooks(authorIds)
	if err != nil {
		handleError(err)
	}
	for _, id := range authorIds {
		author, ok := authorBookMap[id]
		if !ok {
			err := errors.New(fmt.Sprintf("book not found, "+"book_id: %s", id))
			handleError(err)
			results = append(results, &dataloader.Result{
				Data:  []domain.Book{},
				Error: nil,
			})
			continue
		}
		result := dataloader.Result{
			Data:  author,
			Error: nil,
		}
		results = append(results, &result)
	}
	return results
}
