package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/elsagg/graphapis/internal/authors/graph/generated"
	"github.com/elsagg/graphapis/internal/authors/graph/model"
	"github.com/elsagg/graphapis/pkg/events"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *mutationResolver) CreateAuthor(ctx context.Context, input model.NewAuthor) (*model.Author, error) {
	newAuthor := &model.Author{
		ID:   input.ID,
		Name: input.Name,
	}

	err := newAuthor.CreateEvent(ctx, &events.EventMetadata{
		EventType:        "gg.elsa.authors.CreateAuthor",
		EventSource:      "authors/mutation/CreateAuthor",
		EventKey:         newAuthor.ID,
		EventDestination: "authors",
		EventTime:        time.Now(),
	})

	if err != nil {
		return nil, err
	}

	err = newAuthor.Event.SetEventData(newAuthor)

	if err != nil {
		return nil, err
	}

	err = newAuthor.Event.Send()

	if err != nil {
		return nil, err
	}

	return newAuthor, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	var authors []*model.Author

	repository, err := r.Repository("elsagg", "authors")

	if err != nil {
		return nil, err
	}

	res, err := repository.Find(bson.D{}, options.Find())

	if err != nil {
		return nil, err
	}

	cursor := res.(*mongo.Cursor)

	if err = cursor.All(ctx, &authors); err != nil {
		return nil, err
	}

	return authors, nil
}

func (r *queryResolver) Author(ctx context.Context, id string) (*model.Author, error) {
	var author *model.Author

	repository, err := r.Repository("elsagg", "authors")

	if err != nil {
		return nil, err
	}

	res, err := repository.FindOne(bson.D{{"_id", id}}, options.FindOne())

	if err != nil {
		return nil, err
	}

	result := res.(*mongo.SingleResult)

	if err = result.Decode(&author); err != nil {
		return nil, err
	}

	return author, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
