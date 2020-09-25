package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/elsagg/authors/internal/graph/generated"
	"github.com/elsagg/authors/internal/graph/model"
)

func (r *entityResolver) FindAuthorByID(ctx context.Context, id string) (*model.Author, error) {
	/* var author *model.Author

	repository, err := r.AuthorRepository()

	if err != nil {
		return nil, err
	}

	result := repository.FindOne(context.TODO(), bson.D{{"_id", id}}, options.FindOne())

	if err = result.Decode(&author); err != nil {
		return nil, err
	} */

	return nil, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
