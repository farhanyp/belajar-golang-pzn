//go:build wireinject
// +build wireinject

package providerSet

import "github.com/google/wire"

var BarSet = wire.NewSet(NewBarRepository, NewBarService)
var FooSet = wire.NewSet(NewFooRepository, NewFooService)


func InitializedFooBarService() *FooBarService {

	wire.Build(BarSet, FooSet, NewFooBarService)

	return nil
}