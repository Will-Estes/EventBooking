package routes

import (
	"eventbooking/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEventById(context *gin.Context) {
	id := context.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing eventId"})
		return
	}

	event, err := models.GetEventById(i)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting data"})
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing data"})
		return
	}

	event.UserId = context.GetInt64("userId")
	err = event.Save()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving data"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func updateEvent(context *gin.Context) {
	id := context.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing eventId"})
		return
	}
	event, err := models.GetEventById(i)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting event"})
		return
	}

	if event.UserId != context.GetInt64("userId") {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized for event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error parsing event"})
		return
	}

	updatedEvent.ID = i
	err = updatedEvent.Update()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating event"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "Event updated", "event": event})
}

func deleteEvent(context *gin.Context) {
	id := context.Param("id")
	i, _ := strconv.ParseInt(id, 10, 64)
	event, err := models.GetEventById(i)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting event"})
		return
	}

	if event.UserId != context.GetInt64("userId") {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized for event"})
		return
	}

	err = event.Delete()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
