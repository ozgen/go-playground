package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample-rest-api/models"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Id could not be parsed", "error": err.Error()})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetAllEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get an event", "error": err.Error()})
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register an event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "registered"})
}
func cancelRegisterForEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Id could not be parsed", "error": err.Error()})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetAllEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get an event", "error": err.Error()})
		return
	}
	err = event.CancelRegister(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel the registered event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "register cancelled"})
}
