package data

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repositoryMongo struct {
	name       string
	ctx        context.Context
	cancel     context.CancelFunc
	client     *mongo.Client
	collection *mongo.Collection
	database   string
}

func (r repositoryMongo) SetRepository(name string) {
	r.name = name
}

func (r repositoryMongo) SetDatabase(database string) {
	r.database = database
}

func (r repositoryMongo) Database() string {
	return r.database
}

func (r repositoryMongo) Disconnect() {
	defer r.client.Disconnect(r.ctx)
}

func (r repositoryMongo) Find(args ...interface{}) (interface{}, error) {
	return r.client.Database(r.database).Collection(r.name).Find(r.ctx, args[0], args[1].(*options.FindOptions))
}

func (r repositoryMongo) FindOne(args ...interface{}) (interface{}, error) {
	return r.client.Database(r.database).Collection(r.name).FindOne(r.ctx, args[0], args[1].(*options.FindOneOptions)), nil
}

func NewRepositoryMongo(ctx context.Context, database string, collection string) (Repository, error) {
	client, err := GetMongoConnection(ctx)
	repo := repositoryMongo{
		ctx:      ctx,
		database: database,
		name:     collection,
		client:   client,
	}

	if err != nil {
		return nil, err
	}

	return repo, nil
}

func GetMongoConnection(ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}
