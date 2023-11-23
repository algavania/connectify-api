package router

import (
	"example/connectify/app/middlewares"
	"example/connectify/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"*"}
	router.Use(cors.New(config))
	router.Static("/public", "./public")

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		auth.POST("/register", init.UserCtrl.AddUserData)
		auth.POST("/login", init.UserCtrl.Login)

		user := api.Group("/user")
		user.Use(middlewares.JwtAuthMiddleware())
		user.GET("/:username", init.UserCtrl.GetUserByUsername)
		user.PUT("/:userID", init.UserCtrl.UpdateUserData)
		user.DELETE("/:userID", init.UserCtrl.DeleteUser)

		userDetail := api.Group("/user-detail")
		userDetail.Use(middlewares.JwtAuthMiddleware())
		userDetail.GET("/all", init.UserDetailCtrl.GetAllUsers)
		userDetail.POST("/:userID", init.UserDetailCtrl.AddUserData)
		userDetail.GET("/:userID", init.UserDetailCtrl.GetUserById)

		post := api.Group("/post")
		post.Use(middlewares.JwtAuthMiddleware())
		post.POST("", init.PostCtrl.AddPostData)
		post.GET("", init.PostCtrl.GetAllPosts)
		post.GET("/:postID", init.PostCtrl.GetPostById)
		post.GET("/user/:userID", init.PostCtrl.GetPostByUserId)
		post.GET("/:postID/reply", init.PostCtrl.GetAllReplies)
		post.PUT("/:postID", init.PostCtrl.UpdatePostData)
		post.DELETE("/:postID", init.PostCtrl.DeletePost)

		chat := api.Group("/chat")
		chat.Use(middlewares.JwtAuthMiddleware())
		chat.POST("", init.ChatCtrl.AddChatData)
		chat.GET("/:chatID", init.ChatCtrl.GetChatById)
		chat.PUT("/:chatID", init.ChatCtrl.UpdateChatData)
		chat.DELETE("/:chatID", init.ChatCtrl.DeleteChat)

		chat.POST("/:chatID/participant", init.ChatCtrl.AddParticipant)
		chat.DELETE("/:chatID/participant", init.ChatCtrl.DeleteParticipant)

		chat.POST("/:chatID/message", init.ChatCtrl.AddMessage)
		chat.DELETE("/:chatID/message/:messageID", init.ChatCtrl.DeleteMessage)

		userFollowing := api.Group("/user-following")
		userFollowing.Use(middlewares.JwtAuthMiddleware())
		userFollowing.POST("", init.UserFollowingCtrl.Follow)
		userFollowing.GET("/:userID", init.UserFollowingCtrl.GetUserFollowing)
		userFollowing.GET("/:userID/followers", init.UserFollowingCtrl.GetUserFollowers)
		userFollowing.GET("/:userID/count", init.UserFollowingCtrl.GetUserFollowingCount)
		userFollowing.GET("/:userID/followers/count", init.UserFollowingCtrl.GetUserFollowersCount)
		userFollowing.GET("/:userID/check", init.UserFollowingCtrl.CheckHasFollowed)
		userFollowing.DELETE("/:userID", init.UserFollowingCtrl.Unfollow)
	}

	return router
}
