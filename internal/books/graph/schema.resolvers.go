package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/elsagg/graphapis/internal/books/graph/generated"
	"github.com/elsagg/graphapis/internal/books/graph/model"
	"github.com/elsagg/graphapis/pkg/events"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
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

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	var books []*model.Book

	repository, err := r.Repository("elsagg", "books")

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

func (r *queryResolver) Book(ctx context.Context, id string) (*model.Book, error) {
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
