package routes

import (
	"blog-platform/handlers"
	"blog-platform/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	// Static frontend
	r.Static("/static", "./public")

	r.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	r.GET("/login", func(c *gin.Context) {
		c.File("./public/pages/login.html")
	})

	r.GET("/register", func(c *gin.Context) {
		c.File("./public/pages/register.html")
	})

	r.GET("/post", func(c *gin.Context) {
		c.File("./public/pages/post.html")
	})

	r.GET("/create-post", func(c *gin.Context) {
		c.File("./public/pages/create-post.html")
	})

	// AUTH (public)
	auth := r.Group("/auth")

	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	// POSTS (public)
	posts := r.Group("/posts")
	posts.Use(middleware.OptionalMiddlware())
	{
		posts.GET("", handlers.GetPosts)
		posts.GET("/:id", handlers.GetPostById)

		// comments (read)
		posts.GET("/:id/comments", handlers.GetCommentByPostId)

		// likes (read)
		posts.GET("/:id/likes", handlers.GetLikesCount)
	}

	// POSTS (protected)
	authPosts := posts.Group("")
	authPosts.Use(middleware.AuthMiddleware())
	{
		// posts CRUD
		authPosts.POST("", handlers.CreatePost)
		authPosts.PUT("/:id", handlers.UpdatePost)
		authPosts.PATCH("/:id", handlers.PatchPost)
		authPosts.DELETE("/:id", handlers.DeletePost)

		// comments create
		authPosts.POST("/:id/comments", handlers.CreateComment)

		// likes
		authPosts.POST("/:id/like", handlers.LikePost)
		authPosts.DELETE("/:id/like", handlers.UnlikePost)

		// did I like this post?
		authPosts.GET("/:id/liked", handlers.DidUserLikePost)
	}

	// COMMENTS (protected)
	comments := r.Group("/comments")
	comments.Use(middleware.AuthMiddleware())
	{
		comments.DELETE("/:id", handlers.DeleteComment)
	}
}
