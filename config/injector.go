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

var postServiceSet = wire.NewSet(service.PostServiceInit,
	wire.Bind(new(service.PostService), new(*service.PostServiceImpl)),
)

var postRepoSet = wire.NewSet(repository.PostRepositoryInit,
	wire.Bind(new(repository.PostRepository), new(*repository.PostRepositoryImpl)),
)

var postCtrlSet = wire.NewSet(controller.PostControllerInit,
	wire.Bind(new(controller.PostController), new(*controller.PostControllerImpl)),
)

var chatServiceSet = wire.NewSet(service.ChatServiceInit,
	wire.Bind(new(service.ChatService), new(*service.ChatServiceImpl)),
)

var chatRepoSet = wire.NewSet(repository.ChatRepositoryInit,
	wire.Bind(new(repository.ChatRepository), new(*repository.ChatRepositoryImpl)),
)

var chatCtrlSet = wire.NewSet(controller.ChatControllerInit,
	wire.Bind(new(controller.ChatController), new(*controller.ChatControllerImpl)),
)

var userFollowingServiceSet = wire.NewSet(service.UserFollowingServiceInit,
	wire.Bind(new(service.UserFollowingService), new(*service.UserFollowingServiceImpl)),
)

var userFollowingRepoSet = wire.NewSet(repository.UserFollowingRepositoryInit,
	wire.Bind(new(repository.UserFollowingRepository), new(*repository.UserFollowingRepositoryImpl)),
)

var userFollowingCtrlSet = wire.NewSet(controller.UserFollowingControllerInit,
	wire.Bind(new(controller.UserFollowingController), new(*controller.UserFollowingControllerImpl)),
)

func Init() *Initialization {
	wire.Build(
		NewInitialization,
		db,
		userCtrlSet, userServiceSet, userRepoSet,
		userDetailCtrlSet, userDetailServiceSet, userDetailRepoSet,
		postCtrlSet, postServiceSet, postRepoSet,
		chatCtrlSet, chatServiceSet, chatRepoSet,
		userFollowingCtrlSet, userFollowingServiceSet, userFollowingRepoSet,
	)
	return nil
}
