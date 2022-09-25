package graphql

import (
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
	"simple-graphql-server/internal/app/domain"
	"sync"
)

func BookType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Books",
			Fields: graphql.Fields{
				"id":      &graphql.Field{Type: graphql.Int},
				"title":   &graphql.Field{Type: graphql.String},
				"price":   &graphql.Field{Type: graphql.Float},
				"isbn_no": &graphql.Field{Type: graphql.Int},
				"authors": &graphql.Field{
					Type: graphql.NewList(graphql.NewObject(
						graphql.ObjectConfig{
							Name: "BookAuthor",
							Fields: graphql.Fields{
								"id":        &graphql.Field{Type: graphql.Int},
								"name":      &graphql.Field{Type: graphql.String},
								"biography": &graphql.Field{Type: graphql.String},
							},
						})),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						var (
							book        = p.Source.(domain.Book)
							v           = p.Context.Value
							loaders     = v("loaders").(map[string]*dataloader.Loader)
							c           = v("client").(*Client)
							thunks      []dataloader.Thunk
							wg          sync.WaitGroup
							handleError = func(error error) error {
								return fmt.Errorf(error.Error())
							}
						)
						id := book.ID
						key := NewResolverKey(fmt.Sprintf("%d", id), c)
						thunk := loaders["GetAuthors"].Load(p.Context, key)
						thunks = append(thunks, thunk)
						type result struct {
							authors []domain.Author
							err     error
						}
						ch := make(chan *result, 1)
						go func() {
							var errs []error
							for _, thunk := range thunks {
								wg.Add(1)
								go func(t dataloader.Thunk) {
									defer wg.Done()
									r, err := t()
									if err != nil {
										errs = append(errs, err)
										return
									}
									c := r.([]domain.Author)
									ch <- &result{authors: c, err: err}
								}(thunk)
							}
							wg.Wait()
						}()
						return func() (interface{}, error) {
							r := <-ch
							if r.err != nil {
								return nil, handleError(r.err)
							}
							return r.authors, nil
						}, nil
					},
				},
			},
		})
}

func AuthorType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Author",
			Fields: graphql.Fields{
				"id":        &graphql.Field{Type: graphql.Int},
				"name":      &graphql.Field{Type: graphql.String},
				"biography": &graphql.Field{Type: graphql.String},
				"books": &graphql.Field{
					Type: graphql.NewList(graphql.NewObject(
						graphql.ObjectConfig{
							Name: "AuthorBook",
							Fields: graphql.Fields{
								"id":      &graphql.Field{Type: graphql.Int},
								"title":   &graphql.Field{Type: graphql.String},
								"price":   &graphql.Field{Type: graphql.Float},
								"isbn_no": &graphql.Field{Type: graphql.Int},
							},
						})),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						var (
							author      = p.Source.(domain.Author)
							v           = p.Context.Value
							loaders     = v("loaders").(map[string]*dataloader.Loader)
							c           = v("client").(*Client)
							thunks      []dataloader.Thunk
							wg          sync.WaitGroup
							handleError = func(error error) error {
								return fmt.Errorf(error.Error())
							}
						)
						id := author.ID
						key := NewResolverKey(fmt.Sprintf("%d", id), c)
						thunk := loaders["GetBooks"].Load(p.Context, key)
						thunks = append(thunks, thunk)
						type result struct {
							books []domain.Book
							err   error
						}
						ch := make(chan *result, 1)
						go func() {
							var errs []error
							for _, thunk := range thunks {
								wg.Add(1)
								go func(t dataloader.Thunk) {
									defer wg.Done()
									r, err := t()
									if err != nil {
										errs = append(errs, err)
										return
									}
									c := r.([]domain.Book)
									ch <- &result{books: c, err: err}
								}(thunk)
							}
							wg.Wait()
						}()
						return func() (interface{}, error) {
							r := <-ch
							if r.err != nil {
								return nil, handleError(r.err)
							}
							return r.books, nil
						}, nil
					},
				},
			},
		})
}
