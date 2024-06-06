//go:build wireinject
// +build wireinject

package wireWithError

import "github.com/google/wire"


func InitializedService() (*SimpleService, error) {

	wire.Build(NewSimpleRepository, NewSimpleService)

	return nil,nil
}