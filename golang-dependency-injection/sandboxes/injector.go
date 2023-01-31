//go:build wireinject
// +build wireinject

package sandboxes

import (
	"github.com/google/wire"
)

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializeFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

func InitializeSimpeService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
	wire.Build(NewDatabaseMongoDB, NewDatabasePostgreSQL, NewDatabaseRepository)
	return nil
}

var helloSet = wire.NewSet(NewSayHello, wire.Bind(new(SayHelloContract), new(*SayHello)))

func InitializeHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

var fooBarSet = wire.NewSet(NewFoo, NewBar)

func InitializeFooBar() *FooBar {
	wire.Build(fooBarSet, wire.Struct(new(FooBar), "Foo", "Bar")) // "*" for all fiels
	return nil
}

var fooBarValueSet = wire.NewSet(
	wire.Value(&Foo{}),
	wire.Value(&Bar{}))

func InitializeFooBarUsingValue() *FooBar {
	wire.Build(fooBarValueSet, wire.Struct(new(FooBar), "*")) // use "*" for inject all fields
	return nil
}

func InitializeConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}

func InitializeConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
