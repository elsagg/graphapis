package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/elsagg/books/graph/model"
	generated1 "github.com/elsagg/graphapis/internal/books/graph/generated"
	model1 "github.com/elsagg/graphapis/internal/books/graph/model"
	"github.com/elsagg/graphapis/pkg/events"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *authorResolver) Name(ctx context.Context, obj *model1.Author) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateBook(ctx context.Context, input model1.NewBook) (*model1.Book, error) {
	newBook := &model.Book{
		ID:       input.ID,
		Title:    input.Title,
		Year:     input.Year,
		AuthorID: input.AuthorID,
	}

	err := newBook.CreateEvent(ctx, &events.EventMetadata{
		EventType:        "gg.elsa.authors.CreateAuthor",
		EventSource:      "authors/mutation/CreateAuthor",
		EventKey:         newBook.ID,
		EventDestination: "authors",
		EventTime:        time.Now(),
	})

	if err != nil {
		return nil, err
	}

	err = newBook.Event.SetEventData(newBook)

	if err != nil {
		return nil, err
	}

	err = newBook.Event.Send()

	if err != nil {
		return nil, err
	}

	return newBook, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*model1.Book, error) {
	var books []*model.Book

	repository, err := r.Repository("elsagg", "authors")

	if err != nil {
		return nil, err
	}

	res, err := repository.Find(bson.D{}, options.Find())

	if err != nil {
		return nil, err
	}

	cursor := res.(*mongo.Cursor)

	if err = cursor.All(ctx, &books); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *queryResolver) Book(ctx context.Context, id string) (*model1.Book, error) {
	var book *model.Book

	repository, err := r.Repository("elsagg", "books")

	if err != nil {
		return nil, err
	}

	res, err := repository.FindOne(bson.D{{"_id", id}}, options.FindOne())

	if err != nil {
		return nil, err
	}

	result := res.(*mongo.SingleResult)

	if err = result.Decode(&book); err != nil {
		return nil, err
	}

	return book, nil
}

// Author returns generated1.AuthorResolver implementation.
func (r *Resolver) Author() generated1.AuthorResolver { return &authorResolver{r} }

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type authorResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
