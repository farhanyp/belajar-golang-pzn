//go:build wireinject
// +build wireinject

package bindingInterface

import "github.com/google/wire"


var helloSet = wire.NewSet(NewSayHelloImpl, wire.Bind(new(SayHello), new(*SayHelloImpl)))

func InitializedHelloService() *HelloService {

	wire.Build(helloSet, NewHelloService)

	return nil
}