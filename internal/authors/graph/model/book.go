package model

type Book struct {
	ID       string `json:"id"`
	AuthorID string `json:"authorId"`
}

/* func (b *Book) Author() (*Author, error) {
	var author *Author

	conn, err := database.GetConnection()

	if err != nil {
		return nil, err
	}

	repository := conn.Database("elsagg").Collection("authors")

	result := repository.FindOne(context.TODO(), bson.D{{"_id", b.AuthorID}}, options.FindOne())

	if err = result.Decode(&author); err != nil {
		panic(err)
	}

	return author, nil
} */

func (Book) IsEntity() {}
