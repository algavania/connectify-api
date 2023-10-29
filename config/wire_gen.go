// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

import (
	"example/connectify/app/controller"
	"example/connectify/app/repository"
	"example/connectify/app/service"
	"github.com/google/wire"
)

// Injectors from injector.go:

func Init() *Initialization {
	gormDB := ConnectToDB()
	userRepositoryImpl := repository.UserRepositoryInit(gormDB)
	userServiceImpl := service.UserServiceInit(userRepositoryImpl)
	userControllerImpl := controller.UserControllerInit(userServiceImpl)
	userDetailRepositoryImpl := repository.UserDetailRepositoryInit(gormDB)
	userDetailServiceImpl := service.UserDetailServiceInit(userDetailRepositoryImpl)
	userDetailControllerImpl := controller.UserDetailControllerInit(userDetailServiceImpl)
	postRepositoryImpl := repository.PostRepositoryInit(gormDB)
	postServiceImpl := service.PostServiceInit(postRepositoryImpl)
	postControllerImpl := controller.PostControllerInit(postServiceImpl)
	initialization := NewInitialization(userRepositoryImpl, userServiceImpl, userControllerImpl, userDetailRepositoryImpl, userDetailServiceImpl, userDetailControllerImpl, postRepositoryImpl, postServiceImpl, postControllerImpl)
	return initialization
}

// injector.go:

var db = wire.NewSet(ConnectToDB)

var userServiceSet = wire.NewSet(service.UserServiceInit, wire.Bind(new(service.UserService), new(*service.UserServiceImpl)))

var userRepoSet = wire.NewSet(repository.UserRepositoryInit, wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)))

var userCtrlSet = wire.NewSet(controller.UserControllerInit, wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)))

var userDetailServiceSet = wire.NewSet(service.UserDetailServiceInit, wire.Bind(new(service.UserDetailService), new(*service.UserDetailServiceImpl)))

var userDetailRepoSet = wire.NewSet(repository.UserDetailRepositoryInit, wire.Bind(new(repository.UserDetailRepository), new(*repository.UserDetailRepositoryImpl)))

var userDetailCtrlSet = wire.NewSet(controller.UserDetailControllerInit, wire.Bind(new(controller.UserDetailController), new(*controller.UserDetailControllerImpl)))

var postServiceSet = wire.NewSet(service.PostServiceInit, wire.Bind(new(service.PostService), new(*service.PostServiceImpl)))

var postRepoSet = wire.NewSet(repository.PostRepositoryInit, wire.Bind(new(repository.PostRepository), new(*repository.PostRepositoryImpl)))

var postCtrlSet = wire.NewSet(controller.PostControllerInit, wire.Bind(new(controller.PostController), new(*controller.PostControllerImpl)))
