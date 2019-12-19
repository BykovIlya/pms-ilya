package models

import (
	"fmt"
	"utils"
	"time"
)

type TaskWorker struct {
	Id             int64     `json:"id"`
	WorkerId       int64     `json:"worker_id"`
	TaskId         int64     `json:"task_id"`
	Price          float64   `json:"price"`
	RoadPrice      float64   `json:"road_price"`
	StartDate      time.Time `json:"start_date"`
	HumanStartDate string    `json:"human_start_date" time_format:"02-01-2006 15:04:05"`
	EndDate        time.Time `json:"end_date"`
	HumanEndDate   string    `json:"human_end_date" time_format:"02-01-2006 15:04:05"`
	TotalPrice     float64     `json:"total_price"`
	Description    string    `json:"description"`
	Dinner         float64   `json:"dinner"`
	Prepaid        float64   `json:"prepaid"`
	Paid           bool      `json:"paid"`
}

func CreateTaskWorker(w TaskWorker) int64 {

	var uid int64
	err := DB.QueryRow("INSERT INTO task_workers(worker_id,task_id,price,road_price,start_date,end_date,total_price,description,dinner,prepaid,paid) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) returning id",
		w.WorkerId, w.TaskId, w.Price, w.RoadPrice, w.StartDate, w.EndDate, w.TotalPrice, w.Description, w.Dinner, w.Prepaid,w.Paid).Scan(&uid)
	utils.CheckErr(err)
	return uid
}

func UpdateTaskWorker(w TaskWorker) bool {

	stmt, err := DB.Prepare("update task_workers set (worker_id,task_id,price,road_price,start_date,end_date,total_price,description,dinner,prepaid,paid)=($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) where id=$12")
	utils.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(w.WorkerId, w.TaskId, w.Price, w.RoadPrice, w.StartDate, w.EndDate, w.TotalPrice, w.Description, w.Dinner, w.Prepaid,w.Paid, w.Id)
	utils.CheckErr(err)

	affect, err := res.RowsAffected()
	utils.CheckErr(err)

	fmt.Println(affect, "rows changed")
	if affect > 0 {
		return true
	}
	return false
}

func DeleteTaskWorker(id int64) bool {

	stmt, err := DB.Prepare("delete from task_workers where task_id=$1")
	utils.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(id)
	utils.CheckErr(err)

	affect, err := res.RowsAffected()
	utils.CheckErr(err)

	fmt.Println(affect, "rows changed")

	if affect > 0 {
		return true
	}
	return false

}

func GetTaskWorker() []TaskWorker {

	rows, err := DB.Query("SELECT id,worker_id,task_id,price,road_price,start_date,end_date,total_price,description,dinner,prepaid,paid FROM task_workers LIMIT 100")
	utils.CheckErr(err)
	defer rows.Close()

	var ws []TaskWorker
	for rows.Next() {
		w := TaskWorker{}
		err = rows.Scan(&w.Id, &w.WorkerId, &w.TaskId, &w.Price, &w.RoadPrice, &w.StartDate, &w.EndDate, &w.TotalPrice, &w.Description, &w.Dinner, &w.Prepaid,&w.Paid)
		utils.CheckErr(err)

		ws = append(ws, w)
	}

	return ws
}

func GetTaskWorkerById(id int64) *TaskWorker {

	rows, err := DB.Query("SELECT id,worker_id,task_id,price,road_price,start_date,end_date,total_price,description,dinner,prepaid,paid FROM task_workers WHERE id=$1 LIMIT 1", id)
	utils.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		w := TaskWorker{}
		err = rows.Scan(&w.Id, &w.WorkerId, &w.TaskId, &w.Price, &w.RoadPrice, &w.StartDate, &w.EndDate, &w.TotalPrice, &w.Description, &w.Dinner, &w.Prepaid,&w.Paid)
		utils.CheckErr(err)

		return &w
	}

	return nil
}

func GetTaskWorkerByTaskId(taskId int64) []TaskWorker {

	rows, err := DB.Query("SELECT id,worker_id,task_id,price,road_price,start_date,end_date,total_price,description,dinner,prepaid,paid FROM task_workers WHERE task_id=$1", taskId)
	utils.CheckErr(err)
	defer rows.Close()

	var ws []TaskWorker
	for rows.Next() {
		w := TaskWorker{}
		err = rows.Scan(&w.Id, &w.WorkerId, &w.TaskId, &w.Price, &w.RoadPrice, &w.StartDate, &w.EndDate, &w.TotalPrice, &w.Description, &w.Dinner, &w.Prepaid,&w.Paid)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

func GetTaskWorkerByWorkerId(workerId int64) []TaskWorker {

	rows, err := DB.Query("SELECT id,worker_id,task_id,price,road_price,start_date,end_date,total_price,description,dinner,prepaid,paid FROM task_workers WHERE worker_id=$1 LIMIT 1", workerId)
	utils.CheckErr(err)
	defer rows.Close()

	var ws []TaskWorker
	for rows.Next() {
		w := TaskWorker{}
		err = rows.Scan(&w.Id, &w.WorkerId, &w.TaskId, &w.Price, &w.RoadPrice, &w.StartDate, &w.EndDate, &w.TotalPrice, &w.Description, &w.Dinner, &w.Prepaid,&w.Paid)
		utils.CheckErr(err)

		ws = append(ws, w)
	}
	return ws
}

func GetTaskWorkerByWorkerIdAndDate(date time.Time) []TaskWorker {

	lowDate := date.Format("2006-01-02")
	heightDate := date.AddDate(0, 0, 1).Format("2006-01-02")

	rows, err := DB.Query("SELECT id,worker_id,task_id,price,road_price,start_date,end_date,total_price,description,dinner,prepaid,paid FROM task_workers WHERE start_date BETWEEN $1 AND $2", lowDate, heightDate)
	utils.CheckErr(err)
	defer rows.Close()

	var ws []TaskWorker
	for rows.Next() {
		w := TaskWorker{}
		err = rows.Scan(&w.Id, &w.WorkerId, &w.TaskId, &w.Price, &w.RoadPrice, &w.StartDate, &w.EndDate, &w.TotalPrice, &w.Description, &w.Dinner, &w.Prepaid,&w.Paid)
		utils.CheckErr(err)

		ws = append(ws, w)
	}
	return ws
}

func GetTaskWorkersBetweenDates(start time.Time,end time.Time) []TaskWorker {

	//lowDate := start.Format("2006-01-02 15:04:05-07")
	//heightDate := end.Format("2006-01-02 15:04:05-07")
	lowDate := start.Format("2006-01-02")
	//+1 день чтобы включить в выдачу последний день, так как время считается 00:00
	heightDate := end.AddDate(0, 0, 1).Format("2006-01-02")

	fmt.Println(lowDate,heightDate)

	//rows, err := DB.Query("SELECT id,worker_id,task_id,price,road_price,start_date,end_date,total_price,description,dinner,prepaid,paid FROM task_workers WHERE (start_date, end_date) OVERLAPS ($1::TIMESTAMP, $2::TIMESTAMP)", lowDate, heightDate)
	rows, err := DB.Query("SELECT id,worker_id,task_id,price,road_price,start_date,end_date,total_price,description,dinner,prepaid,paid FROM task_workers WHERE (start_date BETWEEN $1 AND $2)", lowDate, heightDate)
	utils.CheckErr(err)
	defer rows.Close()

	var ws []TaskWorker
	for rows.Next() {
		w := TaskWorker{}
		err = rows.Scan(&w.Id, &w.WorkerId, &w.TaskId, &w.Price, &w.RoadPrice, &w.StartDate, &w.EndDate, &w.TotalPrice, &w.Description, &w.Dinner, &w.Prepaid,&w.Paid)
		utils.CheckErr(err)

		ws = append(ws, w)
	}
	return ws
}

func GetTaskWorkersBetweenDatesByWorkerId(workerId int64,start time.Time,end time.Time) []TaskWorker {

	lowDate := start.Format("2006-01-02")
	//+1 день чтобы включить в выдачу последний день, так как время считается 00:00
	heightDate := end.AddDate(0, 0, 1).Format("2006-01-02")

	rows, err := DB.Query("SELECT id,worker_id,task_id,price,road_price,start_date,end_date,total_price,description,dinner,prepaid,paid FROM task_workers WHERE (start_date BETWEEN $1 AND $2) AND worker_id=$3", lowDate, heightDate,workerId)
	utils.CheckErr(err)
	defer rows.Close()

	var ws []TaskWorker
	for rows.Next() {
		w := TaskWorker{}
		err = rows.Scan(&w.Id, &w.WorkerId, &w.TaskId, &w.Price, &w.RoadPrice, &w.StartDate, &w.EndDate, &w.TotalPrice, &w.Description, &w.Dinner, &w.Prepaid,&w.Paid)
		utils.CheckErr(err)

		ws = append(ws, w)
	}
	return ws
}

func UpdatePaymentTaskWorker(w TaskWorker, flag bool) bool {

	stmt, err := DB.Prepare("update task_workers set (paid)=($1) where id=$2")
	utils.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(flag,w.Id)
	utils.CheckErr(err)

	affect, err := res.RowsAffected()
	utils.CheckErr(err)

	fmt.Println(affect, "rows changed")
	if affect > 0 {
		return true
	}
	return false
}

func GetLastDateTaskWorkerByWorkerId(workerId int64) *TaskWorker {

	rows, err := DB.Query("SELECT id,worker_id,task_id,price,road_price,start_date,end_date,total_price,description,dinner,prepaid,paid FROM task_workers WHERE worker_id=$1 ORDER BY end_date DESC LIMIT 1", workerId)
	utils.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		w := TaskWorker{}
		err = rows.Scan(&w.Id, &w.WorkerId, &w.TaskId, &w.Price, &w.RoadPrice, &w.StartDate, &w.EndDate, &w.TotalPrice, &w.Description, &w.Dinner, &w.Prepaid,&w.Paid)
		utils.CheckErr(err)

		return &w
	}

	return nil

}