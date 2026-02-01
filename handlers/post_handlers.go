package handlers

import (
	"blog-platform/models"
	repository "blog-platform/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {

	var userId *int

	uid, ok := c.Get("user_id")
	if ok {
		id := uid.(int)
		userId = &id
	}

	posts, err := repository.GetPostsWithLikes(userId)
	if err != nil {
		fmt.Println("DB ERROR:", err)
		c.JSON(500, gin.H{
			"error": "db error",
		})
		return
	}

	c.JSON(200, posts)
}

func GetPostById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Json",
		})
		return
	}

	post, err := repository.GetPostById(id)

	if err == repository.Notfound {
		c.JSON(404, gin.H{
			"error": "post not found ",
		})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{
			"error": "db error",
		})
		return
	}

	likes, err := repository.GetLikesCount(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "db error",
		})
		return
	}

	LikedByMe := false
	uid, ok := c.Get("user_id")
	if ok {
		userId := uid.(int)
		LikedByMe, _ = repository.DidUserLikePost(id, userId)
	}

	response := models.PostResponse{
		Id:         post.Id,
		Title:      post.Title,
		Content:    post.Content,
		Author_id:  post.Author_id,
		Created_at: post.Created_at,
		Updated_at: post.Updated_at,
		LikesCount: likes,
		LikedByMe:  LikedByMe,
	}

	c.JSON(200, response)

}
func CreatePost(c *gin.Context) {

	var newpost models.CreatePost

	if err := c.ShouldBindJSON(&newpost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
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
	author_id := uid.(int)

	post, err := repository.CreatePost(newpost.Title, newpost.Content, author_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "db error",
		})
		return
	}

	c.JSON(http.StatusCreated, post)

}
func UpdatePost(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	var newPost models.UpdatePost

	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
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
	author_id := uid.(int)

	post, err := repository.UpdatePost(id, newPost.Title, newPost.Content, author_id)
	if err == repository.Notfound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "db error",
		})
		return
	}

	c.JSON(http.StatusOK, post)

}
func PatchPost(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}

	var patP models.PatchPost

	if err := c.ShouldBindJSON(&patP); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}

	if patP.Title == nil && patP.Content == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "nothing to update",
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
	author_id := uid.(int)

	post, err := repository.PatchPost(id, patP.Title, patP.Content, author_id)
	if err == repository.Notfound {
		c.JSON(400, gin.H{
			"error": "Not found",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"error": "db error",
		})
		return
	}

	c.JSON(200, post)
}
func DeletePost(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
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
	author_id := uid.(int)

	err = repository.DeletePost(id, author_id)
	if err == repository.Notfound {
		c.JSON(404, gin.H{
			"error": "post not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "db error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "post deleted",
	})
}
