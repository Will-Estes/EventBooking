package routes

import (
	"eventbooking/models"
	"eventbooking/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing data"})
		return
	}

	err = user.Save()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving data"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not authenticate"})
		return
	}

	err = user.ValidateUser()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not authenticate"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not authenticate"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successful login", "token": token})
}
