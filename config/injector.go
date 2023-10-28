// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"example/connectify/app/controller"
	"example/connectify/app/repository"
	"example/connectify/app/service"

	wire "github.com/google/wire"
)

var db = wire.NewSet(ConnectToDB)

var userServiceSet = wire.NewSet(service.UserServiceInit,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

var userCtrlSet = wire.NewSet(controller.UserControllerInit,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

var userDetailServiceSet = wire.NewSet(service.UserDetailServiceInit,
	wire.Bind(new(service.UserDetailService), new(*service.UserDetailServiceImpl)),
)

var userDetailRepoSet = wire.NewSet(repository.UserDetailRepositoryInit,
	wire.Bind(new(repository.UserDetailRepository), new(*repository.UserDetailRepositoryImpl)),
)

var userDetailCtrlSet = wire.NewSet(controller.UserDetailControllerInit,
	wire.Bind(new(controller.UserDetailController), new(*controller.UserDetailControllerImpl)),
)

func Init() *Initialization {
	wire.Build(
		NewInitialization,
		db,
		userCtrlSet, userServiceSet, userRepoSet,
		userDetailCtrlSet, userDetailServiceSet, userDetailRepoSet,
	)
	return nil
}
