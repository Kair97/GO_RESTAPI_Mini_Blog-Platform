package handlers

import (
	"blog-platform/models"
	repositorty "blog-platform/repository"
	"blog-platform/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var input models.RegisterUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid input",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed while processing password",
		})
		return
	}

	err = repositorty.CreateUser(
		input.Username,
		input.Email,
		hashedPassword,
	)

	if err == repositorty.ErrEmailAlreadyExists {
		c.JSON(409, gin.H{
			"error": "email already exists",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"error": "could not create user",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "user registered succesfully",
	})

}

func Login(c *gin.Context) {

	var input models.LoginUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid input",
		})
		return
	}

	user_id, passwordHash, err := repositorty.GetUserByEmail(input.Email)
	if err == repositorty.ErrInvalidCredentials {
		c.JSON(401, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	err = utils.CheckPassword(input.Password, passwordHash)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	token, err := utils.GenerateToken(user_id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.JSON(200, gin.H{
		"token":   token,
		"user_id": user_id,
	})

}
