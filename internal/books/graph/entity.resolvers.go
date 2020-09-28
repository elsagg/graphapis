package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/elsagg/graphapis/internal/books/graph/generated"
	"github.com/elsagg/graphapis/internal/books/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *entityResolver) FindBookByID(ctx context.Context, id string) (*model.Book, error) {
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

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
