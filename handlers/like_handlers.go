package handlers

import (
	repositorty "blog-platform/repository"
	repository "blog-platform/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LikePost(c *gin.Context) {

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}

	uid, exist := c.Get("user_id")
	if !exist {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}
	user_id := uid.(int)

	err = repository.LikePost(postId, user_id)
	if err == repositorty.ErrAlreadyLiked {
		c.JSON(409, gin.H{
			"error": "post already liked",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{"error": "db error"})
		return
	}

	c.JSON(201, gin.H{
		"message": "post liked",
	})

}

func UnlikePost(c *gin.Context) {

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid id ",
		})
		return
	}

	uid, ok := c.Get("user_id")
	if !ok {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}
	userId := uid.(int)

	err = repositorty.UnlikePost(postId, userId)
	if err == repositorty.ErrLikeNotFound {
		c.JSON(404, gin.H{
			"error": "like not found",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"error": "db error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "post unliked",
	})

}

func GetLikesCount(c *gin.Context) {

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	count, err := repositorty.GetLikesCount(postId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "db error",
		})
		return
	}

	c.JSON(200, gin.H{
		"likes": count,
	})

}

func DidUserLikePost(c *gin.Context) {

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Id",
		})
		return
	}

	uid, ok := c.Get("user_id")
	if !ok {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}
	userId := uid.(int)

	liked, err := repositorty.DidUserLikePost(postId, userId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "db error",
		})
		return
	}

	c.JSON(200, gin.H{
		"liked": liked,
	})

}
