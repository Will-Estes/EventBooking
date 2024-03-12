package routes

import (
	"eventbooking/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getRegistrationsForEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing eventId"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting event"})
		return
	}

	registers, err := event.GetRegistrations()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting registrations for event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"registrations": registers})
}

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing eventId"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error registering for event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registration complete"})
}

func cancelRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing eventId"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting event"})
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error cancelling registration for event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registration cancelled"})
}
