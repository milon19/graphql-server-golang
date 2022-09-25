package controller

import (
	"context"
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
	"log"
	gql "simple-graphql-server/internal/app/adapter/graphql"
	"simple-graphql-server/internal/app/adapter/schema"
	"simple-graphql-server/internal/app/application/usecase"
)

func InitGraphQL(ctx context.Context, query string, service *Service) *graphql.Result {
	bookUseCase := usecase.Book{Repo: service.NewAuthorRepository()}
	authorUseCase := usecase.Author{Repo: service.NewBookRepository()}
	var loaders = make(map[string]*dataloader.Loader, 1)
	var dataloaderClient = gql.Client{}
	loaders["GetAuthors"] = dataloader.NewBatchedLoader(authorUseCase.GetAuthorsBatchFn)
	loaders["GetBooks"] = dataloader.NewBatchedLoader(bookUseCase.GetBooksBatchFn)
	ctx = context.WithValue(ctx, "loaders", loaders)
	ctx = context.WithValue(ctx, "client", &dataloaderClient)
	_schema := schema.Schema{BookRepo: service.NewBookRepository(), AuthorRepo: service.NewAuthorRepository()}.GetSchema()
	params := graphql.Params{Context: ctx, Schema: _schema, RequestString: query}
	result := graphql.Do(params)
	if result.Errors != nil {
		log.Println(result.Errors)
	}
	return result
}
