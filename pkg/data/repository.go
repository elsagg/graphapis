package data

type Repository interface {
	SetDatabase(database string)
	SetRepository(name string)
	Find(args ...interface{}) (interface{}, error)
	FindOne(args ...interface{}) (interface{}, error)
}
