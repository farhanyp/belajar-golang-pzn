package multipleBinding

type Database struct {
	Name string
}

type DatabaseMongoDB Database
type DatabasePostgreSQL Database

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{})
}

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{})
}

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepository(postgresql *DatabasePostgreSQL, mongodb *DatabaseMongoDB) *DatabaseRepository {

	return &DatabaseRepository{
		DatabasePostgreSQL: postgresql,
		DatabaseMongoDB:    mongodb,
	}

}