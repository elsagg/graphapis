package model

import (
	"context"

	"github.com/elsagg/graphapis/pkg/data"
	"github.com/elsagg/graphapis/pkg/events"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Author struct {
	events.EventData
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (a *Author) Books() ([]*Book, error) {
	var books []*Book

	repo, err := data.NewDataViewerMongo(context.TODO(), "elsagg", "books")
	if err != nil {
		return nil, err
	}

	repository := repo.Repository()

	res, err := repository.Find(bson.D{{"authorId", a.ID}}, options.Find().SetSort(bson.D{{"year", 1}}))

	if err != nil {
		return nil, err
	}

	cursor := res.(*mongo.Cursor)

	if err = cursor.All(context.TODO(), &books); err != nil {
		panic(err)
	}

	return books, nil
}

func (Author) IsEntity() {}

/* func (a *Author) Books() ([]*Book, error) {
	var books []*Book

	conn, err := database.GetConnection()

	if err != nil {
		return nil, err
	}

	repository := conn.Database("elsagg").Collection("books")

	cursor, err := repository.Find(context.TODO(), bson.D{{"authorId", a.ID}}, options.Find().SetSort(bson.D{{"year", 1}}))

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &books); err != nil {
		panic(err)
	}

	return books, nil
} */

type NewAuthor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
