package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"models"
	"utils"
	"time"
	"strconv"
)

func GetPayments(c *gin.Context) {

	start := c.DefaultQuery("start", "")
	end := c.DefaultQuery("end", "")

	sd, err := time.Parse(time.RFC3339,start)
	utils.CheckErr(err)

	ed, err := time.Parse(time.RFC3339,end)
	utils.CheckErr(err)

	ps:=models.GetPayments(sd,ed)
	c.JSON(http.StatusOK, ps)
}

func GetDetailPayments(c *gin.Context) {

	start := c.DefaultQuery("start", "")
	end := c.DefaultQuery("end", "")

	sd, err := time.Parse(time.RFC3339,start)
	utils.CheckErr(err)

	ed, err := time.Parse(time.RFC3339,end)
	utils.CheckErr(err)

	worker := c.DefaultQuery("worker", "")
	workerId, err := strconv.ParseInt(worker, 10, 64)
	utils.CheckErr(err)

	ps:=models.GetDetailPayments(workerId,sd,ed)
	c.JSON(http.StatusOK, ps)
}

func UpdatePaidTaskWorker(c *gin.Context) {

	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)
	existItem := models.GetTaskWorkerById(id)

	if existItem == nil {
		c.JSON(http.StatusUnprocessableEntity, ApiStatus{id,utils.EntityWithIdNotExistMessage(taskTitle)})
	}else {
		if models.UpdatePaymentTaskWorker(*existItem,!existItem.Paid) {
			c.JSON(http.StatusOK, ApiStatus{id,utils.SuccessMessage})
		} else {
			c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotUpdatedMessage(taskTitle)})
		}
	}

}

func UpdateAllPaidTaskWorker(c *gin.Context) {

	start := c.DefaultQuery("start", "")
	end := c.DefaultQuery("end", "")

	sd, err := time.Parse(time.RFC3339,start)
	utils.CheckErr(err)

	ed, err := time.Parse(time.RFC3339,end)
	utils.CheckErr(err)

	worker := c.DefaultQuery("worker", "")
	workerId, err := strconv.ParseInt(worker, 10, 64)
	utils.CheckErr(err)

	b := c.DefaultQuery("flag", "")
	flag, err := strconv.ParseBool(b)

	existItems:=models.GetDetailPayments(workerId,sd,ed)

	if len(existItems) == 0 {
		c.JSON(http.StatusUnprocessableEntity, ApiStatus{workerId,utils.EntityWithIdNotExistMessage(taskTitle)})
	}else {
		errs:=0
		for _,item:=range existItems{
			if models.UpdatePaymentTaskWorker(item,flag) {
			} else {
				errs=errs+1
			}
		}

		if errs==0 {
			c.JSON(http.StatusOK, ApiStatus{workerId,utils.SuccessMessage})
		} else {
			c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotUpdatedMessage(taskTitle)})
		}
	}

}






