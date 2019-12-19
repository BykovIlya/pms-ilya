package routes

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"

	"models"
	"utils"
	"fmt"
	"time"

	"log"
	"github.com/tealeg/xlsx"
	"sort"
)


var workerTitle = "Worker"

func CreateWorker(c *gin.Context) {
	var w models.Worker
	if err := c.ShouldBindJSON(&w); err == nil {
		existItem := models.GetWorkerByPhone(w.Phone)
		if existItem != nil {
			c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityExistMessage(workerTitle)})
		} else {
			uid := models.CreateWorker(w)
			c.JSON(http.StatusOK, ApiStatus{uid,utils.SuccessMessage})
		}
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func ReadWorkerById(c *gin.Context) {
	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)
	w := models.GetWorkerById(id)
	if w != nil {
		c.JSON(http.StatusOK, w)
	} else {
		c.JSON(http.StatusNotFound, ApiMessage{utils.EntityNotExistMessage(workerTitle)})
	}
}

func UpdateWorker(c *gin.Context) {
	var w models.Worker
	if err := c.ShouldBindJSON(&w); err == nil {
		existItem := models.GetWorkerById(w.Id)
		if existItem == nil {
			c.JSON(http.StatusUnprocessableEntity, ApiStatus{existItem.Id,utils.EntityWithIdNotExistMessage(workerTitle)})
		}else {
			if models.UpdateWorker(w) {
				c.JSON(http.StatusOK, ApiStatus{w.Id,utils.SuccessMessage})
			} else {
				c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotUpdatedMessage(workerTitle)})
			}
		}
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

//Set Worker Active = false
func UnActiveWorker(c *gin.Context) {
	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)

	existItem := models.GetWorkerById(id)
	if existItem == nil {
		c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityNotExistMessage(workerTitle)})
	} else {
		if models.UnActiveWorker(*existItem) {
			c.JSON(http.StatusOK, ApiMessage{utils.SuccessMessage})
		} else {
			c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotDeletedMessage(workerTitle)})
		}
	}
}

func DeleteWorker(c *gin.Context) {
	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)

	existItem := models.GetWorkerById(id)
	if existItem == nil {
		c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityNotExistMessage(workerTitle)})
	} else {
		if models.DeleteWorker(*existItem) {
			c.JSON(http.StatusOK, ApiMessage{utils.SuccessMessage})
		} else {
			c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotDeletedMessage(workerTitle)})
		}
	}
}

func GetWorkers(c *gin.Context) {
	townParam := c.Param("town")
	town, err := strconv.ParseInt(townParam, 10, 64)
	utils.CheckErr(err)

	c.JSON(http.StatusOK, models.GetWorkerByTown(town))
}

func GetWorkersWithDate(c *gin.Context) {
	townParam := c.Param("town")
	town, err := strconv.ParseInt(townParam, 10, 64)
	utils.CheckErr(err)

	dateParam := c.Param("date")
	t1, err := time.Parse(time.RFC3339,dateParam)
	utils.CheckErr(err)

	dateEndParam := c.Param("end_date")
	t2, err := time.Parse(time.RFC3339,dateEndParam)
	utils.CheckErr(err)

	tws := models.GetTaskWorkersBetweenDates(t1,t2)
	workers:=models.GetWorkerByTown(town)

	var result []models.Worker

	if len(tws)>0{
		for _, w := range workers {
			for _,tw:= range tws{
				if w.Id == tw.WorkerId {
					fmt.Println(tw.WorkerId,tw.WorkerId)
					fmt.Println(tw.StartDate,tw.EndDate)
					w.TaskIds = append(w.TaskIds,tw.TaskId)
				}
			}

			result = append(result,w)
		}
	}else {
		result = workers
	}

	var sortResult []models.Worker

	for _,w:=range result{
		lastTW := models.GetLastDateTaskWorkerByWorkerId(w.Id)
		if lastTW != nil {
			w.LastTaskEndDate = lastTW.EndDate
		}

		sortResult=append(sortResult,w)
	}

	sort.SliceStable(sortResult, func(i, j int) bool {
		return sortResult[j].LastTaskEndDate.Before(sortResult[i].LastTaskEndDate)
	})



	c.JSON(http.StatusOK, sortResult)
}


func ImportWorkerFile(c *gin.Context)  {

	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	excelFileName:="api/upload/"+"File_"+time.Now().Format("20060102150405")+"_"+file.Filename

	if err := c.SaveUploadedFile(file, excelFileName); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	w:=models.GetWorkerById(id)
	w.PassportUrl = excelFileName
	if models.UpdateWorker(*w){
		c.JSON(http.StatusOK, ApiMessage{fmt.Sprintf("File %s uploaded successfully", file.Filename)})
		return
	}

	c.JSON(http.StatusBadRequest, ApiMessage{"something wrong"})
}


func TemporaryImportExcel()  {

	excelFileName:="api/tmp/gruz2.xlsx"

	if readExcel(excelFileName){
		log.Println("End Read file ", excelFileName)
	}
}

func readExcel(excelFileName string) bool{
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		utils.CheckErr(err)
	}

	fmt.Println("Success open file")
	var ws []models.Worker

	//get Sheet 0
	if len(xlFile.Sheets) > 0 {
		fmt.Println("Sheets>0")
		sheet := xlFile.Sheets[0]

		if len(sheet.Rows) > 0 {
			fmt.Println("Rows Count: ", len(sheet.Rows))
			nameRow := sheet.Rows[0]
			for _, cell := range nameRow.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}

			for index, row := range sheet.Rows {
				if index > 0 {
					if len(row.Cells) <= 10 {
						fmt.Println("Cells: ", len(row.Cells))

						name := row.Cells[0]
						fmt.Println("0")
						phone := row.Cells[1]
						fmt.Println("1")
						address := row.Cells[2]
						fmt.Println("2")
						passport := row.Cells[3]
						fmt.Println("3")
						ready := row.Cells[4]
						fmt.Println("4")
						description := row.Cells[5]
						fmt.Println("5")
						cart := row.Cells[7]
						fmt.Println("7")

						passport2 := row.Cells[8]
						fmt.Println("8")

						pr := models.Worker{}
						pr.Name = name.String()
						pr.Phone = phone.String()
						pr.Address = address.String()
						pr.Age = ""
						pr.PassportType = passport.String()+" "+passport2.String()
						pr.PassportUrl = ""
						pr.CartNumber = cart.String()
						pr.Rating = 0
						pr.ActiveToday = false
						pr.Active = true
						pr.Description = description.String()
						pr.GroupId = 1
						pr.Skill = ""
						pr.Ready = ready.String()

						fmt.Println("Import W Phone: ", pr.Phone)
						fmt.Println("Import Len Phone: ", len(pr.Phone))
						if len(pr.Phone) > 0 {
							fmt.Println("Import W: ", pr.Phone)
							ws = append(ws, pr)
						}
					}
				}
			}
		}

		fmt.Println("Import W Count: ", len(ws))

		return models.ImportWorkers(ws)
	}

	return false
}

