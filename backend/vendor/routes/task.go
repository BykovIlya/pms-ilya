package routes

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"models"
	"net/http"
	"roles"
	"strconv"
	"time"
	"utils"
)

var taskTitle = "Task"

func CreateTask(c *gin.Context) {
	var w models.Task
	if err := c.ShouldBindJSON(&w); err == nil {
		//Создаем задачу
		uid := models.CreateTask(w)
		//Привязываем грузчиков к задаче
		for _, tw := range w.TaskWorkers {
			tw.TaskId = uid
			models.CreateTaskWorker(tw)
		}
		c.JSON(http.StatusOK, ApiStatus{uid, utils.SuccessMessage})
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func ReadTaskById(c *gin.Context) {
	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)
	w := models.GetTaskById(id)
	w.TaskWorkers = models.GetTaskWorkerByTaskId(w.Id)

	if w != nil {
		c.JSON(http.StatusOK, w)
	} else {
		c.JSON(http.StatusNotFound, ApiMessage{utils.EntityNotExistMessage(taskTitle)})
	}
}

func UpdateTask(c *gin.Context) {
	var w models.Task
	if err := c.ShouldBindJSON(&w); err == nil {
		existItem := models.GetTaskById(w.Id)
		if existItem == nil {
			c.JSON(http.StatusUnprocessableEntity, ApiStatus{existItem.Id, utils.EntityWithIdNotExistMessage(taskTitle)})
		} else {
			if models.UpdateTask(w) {
				//Привязываем новых грузчиков к задаче
				models.DeleteTaskWorker(w.Id)

				for _, tw := range w.TaskWorkers {
					tw.TaskId = existItem.Id
					models.CreateTaskWorker(tw)
				}
				c.JSON(http.StatusOK, ApiStatus{w.Id, utils.SuccessMessage})
			} else {
				c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotUpdatedMessage(taskTitle)})
			}
		}
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func DeleteTask(c *gin.Context) {
	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)

	existItem := models.GetTaskById(id)
	if existItem == nil {
		c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityNotExistMessage(taskTitle)})
	} else {
		if models.DeleteTask(*existItem) {
			c.JSON(http.StatusOK, ApiMessage{utils.SuccessMessage})
		} else {
			c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotDeletedMessage(taskTitle)})
		}
	}
}

func GetTasks(c *gin.Context) {
	townParam := c.Param("town")
	town, err := strconv.ParseInt(townParam, 10, 64)
	utils.CheckErr(err)

	claims := jwt.ExtractClaims(c)
	username := claims["id"].(string)
	user := models.GetUserByUsername(username)

	if user.Role == roles.ManagerRole {
		tasks := models.GetTasksByTownAndExecutor(town, user.Id)
		c.JSON(http.StatusOK, tasks)
	} else {
		tasks := models.GetTasksByTown(town)
		c.JSON(http.StatusOK, tasks)
	}

}

func GetTasksStatisticsForManagers(c *gin.Context) {
	id := utils.StringToInt(c.DefaultQuery("id", ""))
	start := c.DefaultQuery("start", "")
	end := c.DefaultQuery("end", "")

	sd, err := time.Parse(time.RFC3339, start)
	utils.CheckErr(err)

	ed, err := time.Parse(time.RFC3339, end)
	utils.CheckErr(err)

	/*claims := jwt.ExtractClaims(c)
	username := claims["id"].(string)
	user := models.GetUserByUsername(username)*/

	user := models.GetUserById(id)
	/*	if user.Role == roles.AdminRole {
		if user.Id == id {
			tasks := models.GetTasksByExecutorAndDates(user.Id, sd, ed)
			c.JSON(http.StatusOK, tasks)
		}
	}*/
	if user.Role == roles.ManagerRole {
		if user.Id == id {
			tasks := models.GetTasksByExecutorAndDates(user.Id, sd, ed)
			c.JSON(http.StatusOK, tasks)
		}
	} else {
		tasks := models.GetTasksByDates(sd, ed)
		c.JSON(http.StatusOK, tasks)
	}

}

func GetTasksStatistics(c *gin.Context) {
	start := c.DefaultQuery("start", "")
	end := c.DefaultQuery("end", "")

	sd, err := time.Parse(time.RFC3339, start)
	utils.CheckErr(err)

	ed, err := time.Parse(time.RFC3339, end)
	utils.CheckErr(err)

	claims := jwt.ExtractClaims(c)
	username := claims["id"].(string)
	user := models.GetUserByUsername(username)
	if user.Role == roles.ManagerRole {
		tasks := models.GetTasksByExecutorAndDates(user.Id, sd, ed)
		c.JSON(http.StatusOK, tasks)
	} else {
		tasks := models.GetTasksByDates(sd, ed)
		c.JSON(http.StatusOK, tasks)
	}

}
