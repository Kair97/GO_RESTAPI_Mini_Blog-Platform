package handlers

import (
	"blog-platform/models"
	repositorty "blog-platform/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}

	var input models.CreateComment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid json",
		})
		return
	}

	uid, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{
			"error": "not authorized",
		})
		return
	}
	userId := uid.(int)

	comment, err := repositorty.CreateComment(postId, userId, input.Content)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "db error",
		})
		return
	}

	c.JSON(201, comment)
}

func GetCommentByPostId(c *gin.Context) {

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}

	comments, err := repositorty.GetCommentByPostId(postId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "db error",
		})
		return
	}

	c.JSON(200, comments)

}

func DeleteComment(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}

	uid, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}
	userId := uid.(int)

	err = repositorty.DeleteComment(id, userId)
	if err == repositorty.ErrNotFound {
		c.JSON(404, gin.H{
			"error": "Comment not found",
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
		"message": "comment deleted",
	})

}
