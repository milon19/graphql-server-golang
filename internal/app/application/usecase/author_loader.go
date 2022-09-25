package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"log"
	"simple-graphql-server/internal/app/domain"
	repo "simple-graphql-server/internal/app/domain/repository"
)

type Author struct {
	Repo repo.Book
}

func (a Author) GetAuthorsBatchFn(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	handleError := func(err error) []*dataloader.Result {
		var results []*dataloader.Result
		var result dataloader.Result
		result.Error = err
		log.Println(err)
		results = append(results, &result)
		return results
	}
	var bookIds []string
	var results []*dataloader.Result
	for _, key := range keys {
		bookIds = append(bookIds, key.String())
	}
	bookAuthorMap, err := a.Repo.RelatedAuthors(bookIds)
	if err != nil {
		handleError(err)
	}

	for _, id := range bookIds {
		author, ok := bookAuthorMap[id]
		if !ok {
			err := errors.New(fmt.Sprintf("author not found, "+"author_id: %s", id))
			handleError(err)
			results = append(results, &dataloader.Result{
				Data:  []domain.Author{},
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
