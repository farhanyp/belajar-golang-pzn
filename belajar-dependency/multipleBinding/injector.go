//go:build wireinject
// +build wireinject

package multipleBinding

import "github.com/google/wire"


func InitializedDatabaseRepository() *DatabaseRepository {

	wire.Build(NewDatabaseMongoDB, NewDatabasePostgreSQL, NewDatabaseRepository)

	return nil
}