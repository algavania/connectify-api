package config

import (
	controller "example/connectify/app/controller"
	repository "example/connectify/app/repository"
	service "example/connectify/app/service"
)

type Initialization struct {
	userRepo       repository.UserRepository
	userSvc        service.UserService
	UserCtrl       controller.UserController
	userDetailRepo repository.UserDetailRepository
	userDetailSvc  service.UserDetailService
	UserDetailCtrl controller.UserDetailController
	postRepo       repository.PostRepository
	postSvc        service.PostService
	PostCtrl       controller.PostController
}

func NewInitialization(
	userRepo repository.UserRepository,
	userService service.UserService,
	userCtrl controller.UserController,
	userDetailRepo repository.UserDetailRepository,
	userDetailService service.UserDetailService,
	UserDetailCtrl controller.UserDetailController,
	postRepo repository.PostRepository,
	postService service.PostService,
	postCtrl controller.PostController,
) *Initialization {
	return &Initialization{
		userRepo:       userRepo,
		userSvc:        userService,
		UserCtrl:       userCtrl,
		userDetailRepo: userDetailRepo,
		userDetailSvc:  userDetailService,
		UserDetailCtrl: UserDetailCtrl,
		postRepo:       postRepo,
		postSvc:        postService,
		PostCtrl:       postCtrl,
	}
}
