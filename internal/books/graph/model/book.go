package model

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/elsagg/graphapis/pkg/data"
	"github.com/elsagg/graphapis/pkg/events"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	events.EventData
	ID       string `json:"id"`
	Title    string `json:"title"`
	Year     string `json:"year"`
	AuthorID string `json:"authorId"`
}

func (b *Book) Author() (*Author, error) {
	var author *Author

	repo, err := data.NewDataViewerMongo(context.TODO(), "elsagg", "authors")

	if err != nil {
		return nil, err
	}

	repository := repo.Repository()

	result, err := repository.FindOne(bson.D{{"_id", b.AuthorID}}, options.FindOne())

	if err != nil {
		return nil, err
	}

	if err = result.(*mongo.SingleResult).Decode(&author); err != nil {
		log.Error().Err(err).Msg(err.Error())
	}

	return author, nil
}

func (Book) IsEntity() {}

type NewBook struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Year     string `json:"year"`
	AuthorID string `json:"authorId"`
}
