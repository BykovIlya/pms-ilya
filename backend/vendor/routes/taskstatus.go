package routes

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"

	"models"
	"utils"
)


var TaskStatusTitle = "Task Status"

func CreateTaskStatus(c *gin.Context) {
	var w models.TaskStatus
	if err := c.ShouldBindJSON(&w); err == nil {
		existItem := models.GetTaskStatusByName(w.Name)
		if existItem != nil {
			c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityExistMessage(TaskStatusTitle)})
		} else {
			uid := models.CreateTaskStatus(w)
			c.JSON(http.StatusOK, ApiStatus{uid,utils.SuccessMessage})
		}
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func ReadTaskStatusById(c *gin.Context) {
	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)
	w := models.GetTaskStatusById(id)
	if w != nil {
		c.JSON(http.StatusOK, w)
	} else {
		c.JSON(http.StatusNotFound, ApiMessage{utils.EntityNotExistMessage(TaskStatusTitle)})
	}
}

func UpdateTaskStatus(c *gin.Context) {
	var w models.TaskStatus
	if err := c.ShouldBindJSON(&w); err == nil {
		existItem := models.GetTaskStatusById(w.Id)
		if existItem == nil {
			c.JSON(http.StatusUnprocessableEntity, ApiStatus{existItem.Id,utils.EntityWithIdNotExistMessage(TaskStatusTitle)})
		}else {
			if models.UpdateTaskStatus(w) {
				c.JSON(http.StatusOK, ApiStatus{w.Id,utils.SuccessMessage})
			} else {
				c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotUpdatedMessage(TaskStatusTitle)})
			}
		}
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func DeleteTaskStatus(c *gin.Context) {
	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)

	existItem := models.GetTaskStatusById(id)
	if existItem == nil {
		c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityNotExistMessage(TaskStatusTitle)})
	} else {
		if models.DeleteTaskStatus(*existItem) {
			c.JSON(http.StatusOK, ApiMessage{utils.SuccessMessage})
		} else {
			c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotDeletedMessage(TaskStatusTitle)})
		}
	}
}

func GetTaskStatuses(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetTaskStatuses())
}
