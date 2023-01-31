// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package sandboxes

import (
	"github.com/google/wire"
)

// Injectors from injector.go:

func InitializeFooBarService() *FooBarService {
	fooRepository := NewFooRepository()
	fooService := NewFooService(fooRepository)
	barRepository := NewBarRepository()
	barService := NewBarService(barRepository)
	fooBarService := NewFooBarService(fooService, barService)
	return fooBarService
}

func InitializeSimpeService(isError bool) (*SimpleService, error) {
	simpleRepository := NewSimpleRepository(isError)
	simpleService, err := NewSimpleService(simpleRepository)
	if err != nil {
		return nil, err
	}
	return simpleService, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
	databasePostgreSQL := NewDatabasePostgreSQL()
	databaseMongoDB := NewDatabaseMongoDB()
	databaseRepository := NewDatabaseRepository(databasePostgreSQL, databaseMongoDB)
	return databaseRepository
}

func InitializeHelloService() *HelloService {
	sayHello := NewSayHello()
	helloService := NewHelloService(sayHello)
	return helloService
}

func InitializeFooBar() *FooBar {
	foo := NewFoo()
	bar := NewBar()
	fooBar := &FooBar{
		Foo: foo,
		Bar: bar,
	}
	return fooBar
}

func InitializeFooBarUsingValue() *FooBar {
	foo := _wireFooValue
	bar := _wireBarValue
	fooBar := &FooBar{
		Foo: foo,
		Bar: bar,
	}
	return fooBar
}

var (
	_wireFooValue = &Foo{}
	_wireBarValue = &Bar{}
)

func InitializeConfiguration() *Configuration {
	application := NewApplication()
	configuration := application.Configuration
	return configuration
}

func InitializeConnection(name string) (*Connection, func()) {
	file, cleanup := NewFile(name)
	connection, cleanup2 := NewConnection(file)
	return connection, func() {
		cleanup2()
		cleanup()
	}
}

// injector.go:

var fooSet = wire.NewSet(NewFooRepository, NewFooService)

var barSet = wire.NewSet(NewBarRepository, NewBarService)

var helloSet = wire.NewSet(NewSayHello, wire.Bind(new(SayHelloContract), new(*SayHello)))

var fooBarSet = wire.NewSet(NewFoo, NewBar)

var fooBarValueSet = wire.NewSet(wire.Value(&Foo{}), wire.Value(&Bar{}))