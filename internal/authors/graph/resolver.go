package graph

import (
	"context"

	"github.com/elsagg/authors/pkg/data"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
}

func (r *Resolver) Repository(database string, collection string) (data.Repository, error) {
	repo, err := data.NewDataViewerMongo(context.TODO(), database, collection)
	if err != nil {
		return nil, err
	}
	return repo.Repository(), nil
	/* conn, err := database.GetConnection()

	if err != nil {
		return nil, err
	}

	return conn.Database("elsagg").Collection("authors"), nil */
}