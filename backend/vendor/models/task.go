package models

import (
	"fmt"
	"utils"
	"time"
)

type Task struct {
	Id                 int64     `json:"id"`
	Name               string    `json:"name"`
	StartAddress       string    `json:"start_address"`
	EndAddress         string    `json:"end_address"`
	StartDate          time.Time `json:"start_date"`
	HumanStartDate     string    `json:"human_start_date" time_format:"02-01-2006 15:04:05"`
	EndDate            time.Time `json:"end_date"`
	HumanEndDate       string    `json:"human_end_date" time_format:"02-01-2006 15:04:05"`
	PaymentType        bool      `json:"payment_type"`
	PaymentFirst       float64   `json:"payment_first"`
	PaymentSecond      float64   `json:"payment_second"`
	WorkersCount       int64     `json:"workers_count"`
	TotalHours         float64   `json:"total_hours"`
	TotalPrice         float64   `json:"total_price"`
	TotalWorkerPrice   float64   `json:"total_worker_price"`
	TotalDelta         float64   `json:"total_delta"`
	AdditionalCost     float64   `json:"additional_cost"`
	Description        string    `json:"description"`
	CargoDescription   string    `json:"cargo_description"`
	CreateDate         time.Time `json:"create_date"`
	UpdateDate         time.Time `json:"update_date"`
	StatusId           int64     `json:"status_id"`
	StatusName         string    `json:"status_name"`
	PaymentFirstCount  float64   `json:"payment_first_count"`
	PaymentSecondCount float64   `json:"payment_second_count"`
	AdditionalIncome   float64   `json:"additional_income"`

	CustomerId   int64        `json:"customer_id"`
	CustomerName string       `json:"customer_name"`
	TaskWorkers  []TaskWorker `json:"task_workers"`

	Author       int64  `json:"author"`
	Executor     int64  `json:"executor"`
	ExecutorName string `json:"executor_name"`

	Paid           bool    `json:"paid"`
	Contact        string  `json:"contact"`
	DeclaratedCost float64 `json:"declarated_cost"`
	TownId         int64   `json:"town_id"`
	TotalManager   float64 `json:"total_manager"`
}

func CreateTask(w Task) int64 {

	var uid int64
	w.CreateDate = time.Now()
	w.UpdateDate = time.Now()
	err := DB.QueryRow("INSERT INTO tasks(name,start_address,end_address,start_date,end_date,payment_type,payment_first,payment_second,workers_count,"+
		"total_hours,total_price,total_worker_price,total_delta,additional_cost,description,cargo_description,create_date,update_date,status_id,payment_first_count,payment_second_count,customer_id,additional_income,author,executor,paid,contact,declarated_cost,town_id,total_manager)"+
		" VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30) returning id", w.Name, w.StartAddress, w.EndAddress, w.StartDate, w.EndDate, w.PaymentType, w.PaymentFirst, w.PaymentSecond, w.WorkersCount,
		w.TotalHours, w.TotalPrice, w.TotalWorkerPrice, w.TotalDelta, w.AdditionalCost, w.Description, w.CargoDescription, w.CreateDate, w.UpdateDate, w.StatusId, w.PaymentFirstCount, w.PaymentSecondCount, w.CustomerId, w.AdditionalIncome, w.Author, w.Executor, w.Paid, w.Contact, w.DeclaratedCost, w.TownId, w.TotalManager).Scan(&uid)
	utils.CheckErr(err)
	return uid
}

func UpdateTask(w Task) bool {

	w.UpdateDate = time.Now()
	stmt, err := DB.Prepare("update tasks set (name,start_address,end_address,start_date,end_date,payment_type,payment_first,payment_second,workers_count," +
		"total_hours,total_price,total_worker_price,total_delta,additional_cost,description,cargo_description,create_date,update_date,status_id,payment_first_count,payment_second_count,customer_id,additional_income,author,executor,paid,contact,declarated_cost,town_id,total_manager)=" +
		"($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30) where id=$31")
	utils.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(w.Name, w.StartAddress, w.EndAddress, w.StartDate, w.EndDate, w.PaymentType, w.PaymentFirst, w.PaymentSecond, w.WorkersCount,
		w.TotalHours, w.TotalPrice, w.TotalWorkerPrice, w.TotalDelta, w.AdditionalCost, w.Description, w.CargoDescription, w.CreateDate, w.UpdateDate, w.StatusId, w.PaymentFirstCount, w.PaymentSecondCount, w.CustomerId, w.AdditionalIncome, w.Author, w.Executor, w.Paid, w.Contact, w.DeclaratedCost, w.TownId, w.TotalManager, w.Id)
	utils.CheckErr(err)

	affect, err := res.RowsAffected()
	utils.CheckErr(err)

	fmt.Println(affect, "rows changed")
	if affect > 0 {
		return true
	}
	return false
}

func DeleteTask(w Task) bool {

	stmt, err := DB.Prepare("delete from tasks where id=$1")
	utils.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(w.Id)
	utils.CheckErr(err)

	affect, err := res.RowsAffected()
	utils.CheckErr(err)

	fmt.Println(affect, "rows changed")

	if affect > 0 {
		return true
	}
	return false

}

func GetTasks() []Task {

	rows, err := DB.Query("SELECT id,name,start_address,end_address,start_date,end_date,payment_type,payment_first,payment_second,workers_count," +
		"total_hours,total_price,total_worker_price,total_delta,additional_cost,description,cargo_description,create_date,update_date,status_id,payment_first_count,payment_second_count,customer_id,additional_income,author,executor,paid,contact,declarated_cost,town_id,total_manager FROM tasks LIMIT 1000")
	utils.CheckErr(err)
	defer rows.Close()

	var ws []Task
	for rows.Next() {
		w := Task{}
		err = rows.Scan(&w.Id, &w.Name, &w.StartAddress, &w.EndAddress, &w.StartDate, &w.EndDate, &w.PaymentType, &w.PaymentFirst, &w.PaymentSecond, &w.WorkersCount,
			&w.TotalHours, &w.TotalPrice, &w.TotalWorkerPrice, &w.TotalDelta, &w.AdditionalCost, &w.Description, &w.CargoDescription, &w.CreateDate, &w.UpdateDate, &w.StatusId, &w.PaymentFirstCount, &w.PaymentSecondCount, &w.CustomerId, &w.AdditionalIncome, &w.Author, &w.Executor, &w.Paid, &w.Contact, &w.DeclaratedCost, &w.TownId, &w.TotalManager)
		utils.CheckErr(err)

		w.StatusName = GetTaskStatusById(w.StatusId).Name
		if w.CustomerId > 0 {
			customer := GetCustomerById(w.CustomerId)
			if customer != nil {
				w.CustomerName = customer.Name
			}
		}
		if w.Executor > 0 {
			executor := GetUserById(w.Executor)
			if executor != nil {
				w.ExecutorName = executor.Name
			}
		}

		w.TaskWorkers = GetTaskWorkerByTaskId(w.Id)
		ws = append(ws, w)
	}

	return ws
}

func GetTaskById(id int64) *Task {

	rows, err := DB.Query("SELECT id,name,start_address,end_address,start_date,end_date,payment_type,payment_first,payment_second,workers_count,"+
		"total_hours,total_price,total_worker_price,total_delta,additional_cost,description,cargo_description,create_date,update_date,status_id,payment_first_count,payment_second_count,customer_id,additional_income,author,executor,paid,contact,declarated_cost,town_id,total_manager FROM tasks WHERE id=$1 LIMIT 1", id)
	utils.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		w := Task{}
		err = rows.Scan(&w.Id, &w.Name, &w.StartAddress, &w.EndAddress, &w.StartDate, &w.EndDate, &w.PaymentType, &w.PaymentFirst, &w.PaymentSecond, &w.WorkersCount,
			&w.TotalHours, &w.TotalPrice, &w.TotalWorkerPrice, &w.TotalDelta, &w.AdditionalCost, &w.Description, &w.CargoDescription, &w.CreateDate, &w.UpdateDate, &w.StatusId, &w.PaymentFirstCount, &w.PaymentSecondCount, &w.CustomerId, &w.AdditionalIncome, &w.Author, &w.Executor, &w.Paid, &w.Contact, &w.DeclaratedCost, &w.TownId, &w.TotalManager)
		utils.CheckErr(err)
		return &w
	}

	return nil
}

func GetTasksByTown(p int64) []Task {

	rows, err := DB.Query("SELECT id,name,start_address,end_address,start_date,end_date,payment_type,payment_first,payment_second,workers_count,"+
		"total_hours,total_price,total_worker_price,total_delta,additional_cost,description,cargo_description,create_date,update_date,status_id,payment_first_count,payment_second_count,customer_id,additional_income,author,executor,paid,contact,declarated_cost,town_id,total_manager FROM tasks WHERE town_id=$1 LIMIT 1000", p)
	utils.CheckErr(err)
	defer rows.Close()

	var ws []Task
	for rows.Next() {
		w := Task{}
		err = rows.Scan(&w.Id, &w.Name, &w.StartAddress, &w.EndAddress, &w.StartDate, &w.EndDate, &w.PaymentType, &w.PaymentFirst, &w.PaymentSecond, &w.WorkersCount,
			&w.TotalHours, &w.TotalPrice, &w.TotalWorkerPrice, &w.TotalDelta, &w.AdditionalCost, &w.Description, &w.CargoDescription, &w.CreateDate, &w.UpdateDate, &w.StatusId, &w.PaymentFirstCount, &w.PaymentSecondCount, &w.CustomerId, &w.AdditionalIncome, &w.Author, &w.Executor, &w.Paid, &w.Contact, &w.DeclaratedCost, &w.TownId, &w.TotalManager)
		utils.CheckErr(err)

		ws = append(ws, w)
	}

	var result []Task
	for _,w:=range ws{
		w.StatusName = GetTaskStatusById(w.StatusId).Name
		if w.CustomerId > 0 {
			customer := GetCustomerById(w.CustomerId)
			if customer != nil {
				w.CustomerName = customer.Name
			}
		}
		if w.Executor > 0 {
			executor := GetUserById(w.Executor)
			if executor != nil {
				w.ExecutorName = executor.Name
			}
		}

		w.TaskWorkers = GetTaskWorkerByTaskId(w.Id)
		result = append(result, w)
	}

	return result
}

func GetTasksByTownAndExecutor(p int64, executor int64) []Task {

	rows, err := DB.Query("SELECT id,name,start_address,end_address,start_date,end_date,payment_type,payment_first,payment_second,workers_count,"+
		"total_hours,total_price,total_worker_price,total_delta,additional_cost,description,cargo_description,create_date,update_date,status_id,payment_first_count,payment_second_count,customer_id,additional_income,author,executor,paid,contact,declarated_cost,town_id,total_manager FROM tasks WHERE town_id=$1 AND executor=$2 LIMIT 1000", p, executor)
	utils.CheckErr(err)
	defer rows.Close()

	var ws []Task
	for rows.Next() {
		w := Task{}
		err = rows.Scan(&w.Id, &w.Name, &w.StartAddress, &w.EndAddress, &w.StartDate, &w.EndDate, &w.PaymentType, &w.PaymentFirst, &w.PaymentSecond, &w.WorkersCount,
			&w.TotalHours, &w.TotalPrice, &w.TotalWorkerPrice, &w.TotalDelta, &w.AdditionalCost, &w.Description, &w.CargoDescription, &w.CreateDate, &w.UpdateDate, &w.StatusId, &w.PaymentFirstCount, &w.PaymentSecondCount, &w.CustomerId, &w.AdditionalIncome, &w.Author, &w.Executor, &w.Paid, &w.Contact, &w.DeclaratedCost, &w.TownId, &w.TotalManager)
		utils.CheckErr(err)

		w.StatusName = GetTaskStatusById(w.StatusId).Name
		if w.CustomerId > 0 {
			customer := GetCustomerById(w.CustomerId)
			if customer != nil {
				w.CustomerName = customer.Name
			}
		}
		if w.Executor > 0 {
			executor := GetUserById(w.Executor)
			if executor != nil {
				w.ExecutorName = executor.Name
			}
		}

		w.TaskWorkers = GetTaskWorkerByTaskId(w.Id)
		ws = append(ws, w)
	}

	return ws
}

func GetTasksByExecutorAndDates(executor int64, start time.Time, end time.Time) []Task {

	lowDate := start.Format("2006-01-02")
	heightDate := end.Format("2006-01-02")

	rows, err := DB.Query("SELECT id,name,start_address,end_address,start_date,end_date,payment_type,payment_first,payment_second,workers_count,"+
		"total_hours,total_price,total_worker_price,total_delta,additional_cost,description,cargo_description,create_date,update_date,status_id,payment_first_count,payment_second_count,customer_id,additional_income,author,executor,paid,contact,declarated_cost,town_id,total_manager FROM tasks WHERE executor=$1 AND start_date BETWEEN $2 AND $3", executor, lowDate, heightDate)
	utils.CheckErr(err)
	defer rows.Close()

	var ws []Task
	for rows.Next() {
		w := Task{}
		err = rows.Scan(&w.Id, &w.Name, &w.StartAddress, &w.EndAddress, &w.StartDate, &w.EndDate, &w.PaymentType, &w.PaymentFirst, &w.PaymentSecond, &w.WorkersCount,
			&w.TotalHours, &w.TotalPrice, &w.TotalWorkerPrice, &w.TotalDelta, &w.AdditionalCost, &w.Description, &w.CargoDescription, &w.CreateDate, &w.UpdateDate, &w.StatusId, &w.PaymentFirstCount, &w.PaymentSecondCount, &w.CustomerId, &w.AdditionalIncome, &w.Author, &w.Executor, &w.Paid, &w.Contact, &w.DeclaratedCost, &w.TownId, &w.TotalManager)
		utils.CheckErr(err)

		w.StatusName = GetTaskStatusById(w.StatusId).Name
		if w.CustomerId > 0 {
			customer := GetCustomerById(w.CustomerId)
			if customer != nil {
				w.CustomerName = customer.Name
			}
		}
		if w.Executor > 0 {
			executor := GetUserById(w.Executor)
			if executor != nil {
				w.ExecutorName = executor.Name
			}
		}

		w.TaskWorkers = GetTaskWorkerByTaskId(w.Id)
		ws = append(ws, w)
	}

	return ws
}

func GetTasksByDates(start time.Time, end time.Time) []Task {

	lowDate := start.Format("2006-01-02")
	heightDate := end.Format("2006-01-02")

	rows, err := DB.Query("SELECT id,name,start_address,end_address,start_date,end_date,payment_type,payment_first,payment_second,workers_count,"+
		"total_hours,total_price,total_worker_price,total_delta,additional_cost,description,cargo_description,create_date,update_date,status_id,payment_first_count,payment_second_count,customer_id,additional_income,author,executor,paid,contact,declarated_cost,town_id,total_manager FROM tasks WHERE start_date BETWEEN $1 AND $2", lowDate, heightDate)
	utils.CheckErr(err)
	defer rows.Close()

	var ws []Task
	for rows.Next() {
		w := Task{}
		err = rows.Scan(&w.Id, &w.Name, &w.StartAddress, &w.EndAddress, &w.StartDate, &w.EndDate, &w.PaymentType, &w.PaymentFirst, &w.PaymentSecond, &w.WorkersCount,
			&w.TotalHours, &w.TotalPrice, &w.TotalWorkerPrice, &w.TotalDelta, &w.AdditionalCost, &w.Description, &w.CargoDescription, &w.CreateDate, &w.UpdateDate, &w.StatusId, &w.PaymentFirstCount, &w.PaymentSecondCount, &w.CustomerId, &w.AdditionalIncome, &w.Author, &w.Executor, &w.Paid, &w.Contact, &w.DeclaratedCost, &w.TownId, &w.TotalManager)
		utils.CheckErr(err)

		w.StatusName = GetTaskStatusById(w.StatusId).Name
		if w.CustomerId > 0 {
			customer := GetCustomerById(w.CustomerId)
			if customer != nil {
				w.CustomerName = customer.Name
			}
		}
		if w.Executor > 0 {
			executor := GetUserById(w.Executor)
			if executor != nil {
				w.ExecutorName = executor.Name
			}
		}

		w.TaskWorkers = GetTaskWorkerByTaskId(w.Id)
		ws = append(ws, w)
	}

	return ws
}
