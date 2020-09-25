package data

import (
	"context"
)

type mongoViewer struct {
	repository Repository
}

func (m mongoViewer) Repository() Repository {
	return m.repository
}

func NewDataViewerMongo(ctx context.Context, database string, collection string) (DataViewer, error) {
	repo, err := NewRepositoryMongo(ctx, database, collection)

	if err != nil {
		return nil, err
	}

	viewer := mongoViewer{
		repository: repo,
	}

	return viewer, nil
}

/* func (m mongoViewer) GetConnection() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DATABASE_URI")))
	if err != nil {
		m.cancel()
		return err
	}

	m.client = client

	err = m.client.Connect(m.ctx)
	if err != nil {
		m.cancel()
		return err
	}
	m.cancel()
	return nil
}

func (m mongoViewer) SetDatabase(database string) {
	m.database = database
}

func (m mongoViewer) Database() string {
	return m.database
}

func (m mongoViewer) Disconnect() {
	defer m.client.Disconnect(m.ctx)
}

func (m mongoViewer) GetRepository(name string) interface{} {
	return m.client.Database(m.database).Collection(name)
}

func NewDataViewerMongo() DataViewer {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	viewer := mongoViewer{
		ctx:    ctx,
		cancel: cancel,
	}

	viewer.GetConnection()

	return viewer
}

/* type Repository struct {
	client *mongo.Client
	ctx    context.Context
}

func (r *Repository) GetConnection() {

} */

/*
func GetConnection() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DATABASE_URI")))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		cancel()
		return nil, err
	}
	// defer client.Disconnect(ctx)
	cancel()
	return client, nil
}
*/
