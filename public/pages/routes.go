package routes

import (
	"blog-platform/handlers"
	"blog-platform/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.Static("/static", "./public")

	r.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	r.GET("/login", func(c *gin.Context) {
		c.File("./public/login.html")
	})

	r.GET("/register", func(c *gin.Context) {
		c.File("./public/register.html")
	})

	posts := r.Group("/posts")
	{
		posts.GET("", handlers.GetPosts)
		posts.GET("/:id", handlers.GetPostById)
		//Comments
		posts.GET("/:id/comments", handlers.GetCommentByPostId)

		//Likes
		posts.GET("/:id/likes", handlers.GetLikesCount)
	}

	authPosts := posts.Group("")
	authPosts.Use(middleware.AuthMiddleware())
	{
		// Posts actions protected
		authPosts.POST("", handlers.CreatePost)
		authPosts.PUT("/:id", handlers.UpdatePost)
		authPosts.PATCH("/:id", handlers.PatchPost)
		authPosts.DELETE("/:id", handlers.DeletePost)

		// Comment create
		authPosts.POST("/:id/comments", handlers.CreateComment)

		//Likes create/delete
		authPosts.POST("/:id/like", handlers.LikePost)
		authPosts.DELETE("/:id/like", handlers.UnlikePost)
		posts.GET("/:id/liked", handlers.DidUserLikePost)
	}

	authComments := r.Group("/comments")
	authComments.Use(middleware.AuthMiddleware())
	{
		authComments.DELETE("/:id", handlers.DeleteComment)
	}

	// AUTH
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

}
