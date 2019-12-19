package models

import (
	"fmt"
	"utils"
)

type TaskStatus struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
}

func CreateTaskStatus(w TaskStatus) int64 {

	var uid int64
	err := DB.QueryRow("INSERT INTO taskstatus(name,description) VALUES($1,$2) returning id", w.Name, w.Description).Scan(&uid)
	utils.CheckErr(err)
	return uid
}

func UpdateTaskStatus(w TaskStatus) bool {

	stmt, err := DB.Prepare("update taskstatus set (name,description)=($1,$2) where id=$3")
	utils.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(w.Name, w.Description, w.Id)
	utils.CheckErr(err)

	affect, err := res.RowsAffected()
	utils.CheckErr(err)

	fmt.Println(affect, "rows changed")
	if affect > 0 {
		return true
	}
	return false
}

func DeleteTaskStatus(w TaskStatus) bool {

	stmt, err := DB.Prepare("delete from taskstatus where id=$1")
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

func GetTaskStatuses() []TaskStatus {

	rows, err := DB.Query("SELECT id,name,description FROM taskstatus LIMIT 1000")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []TaskStatus
	for rows.Next() {
		w := TaskStatus{}
		err = rows.Scan(&w.Id, &w.Name, &w.Description)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

func GetTaskStatusById(id int64) *TaskStatus {

	rows, err := DB.Query("SELECT id,name,description FROM taskstatus WHERE id=$1 LIMIT 1", id)
	utils.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		w := TaskStatus{}
		err = rows.Scan(&w.Id, &w.Name, &w.Description)
		utils.CheckErr(err)
		return &w
	}
	return nil
}

func GetTaskStatusByName(p string) *TaskStatus {

	rows, err := DB.Query("SELECT id,name,description FROM taskstatus WHERE name=$1 LIMIT 1", p)
	utils.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		w := TaskStatus{}
		err = rows.Scan(&w.Id, &w.Name, &w.Description)
		utils.CheckErr(err)
		return &w
	}
	return nil
}