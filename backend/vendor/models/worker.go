package models

import (
	"fmt"
	"utils"
	"github.com/lib/pq"
	"log"
	"time"
)

type Worker struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	Phone        string  `json:"phone"`
	Address      string  `json:"address"`
	Age          string  `json:"age"`
	PassportType string  `json:"passport_type"`
	PassportUrl  string  `json:"passport_url"`
	CartNumber   string  `json:"cart_number"`
	Rating       float64 `json:"rating"`
	ActiveToday  bool    `json:"active_today"`
	Active       bool    `json:"active"`
	Description  string  `json:"description"`
	GroupId      int64   `json:"group_id"`
	Skill        string   `json:"skill"`
	Ready        string   `json:"ready"`
	TownId       int64   `json:"town_id"`
	TaskIds      []int64   `json:"task_ids"`
	LastTaskEndDate time.Time `json:"last_task_end_date"`
}

func CreateWorker(w Worker) int64 {

	var uid int64
	err := DB.QueryRow("INSERT INTO workers(name,phone,address,age,passport_type,passport_url,cart_number,rating,active_today,active,description,group_id,skill,ready,town_id)"+
		" VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15) returning id",
		w.Name, w.Phone, w.Address, w.Age, w.PassportType, w.PassportUrl, w.CartNumber, w.Rating, w.ActiveToday, w.Active, w.Description, w.GroupId, w.Skill,w.Ready,w.TownId).Scan(&uid)
	utils.CheckErr(err)
	return uid
}

func UpdateWorker(w Worker) bool {

	stmt, err := DB.Prepare("update workers set (name,phone,address,age,passport_type,passport_url,cart_number,rating,active_today,active,description,group_id,skill,ready,town_id)=($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15) where id=$16")
	utils.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(w.Name, w.Phone, w.Address, w.Age, w.PassportType, w.PassportUrl, w.CartNumber, w.Rating, w.ActiveToday, w.Active, w.Description, w.GroupId, w.Skill,w.Ready,w.TownId, w.Id)
	utils.CheckErr(err)

	affect, err := res.RowsAffected()
	utils.CheckErr(err)

	fmt.Println(affect, "rows changed")
	if affect > 0 {
		return true
	}
	return false
}

func UnActiveWorker(w Worker) bool {

	stmt, err := DB.Prepare("update workers set (active)=(false) where id=$1")
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

func DeleteWorker(w Worker) bool {

	stmt, err := DB.Prepare("delete from workers where id=$1")
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

func GetWorkers() []Worker {

	rows, err := DB.Query("SELECT id,name,phone,address,age,passport_type,passport_url,cart_number,rating,active_today,active,description,group_id,skill,ready,town_id FROM workers LIMIT 1000")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Worker
	for rows.Next() {
		w := Worker{}
		err = rows.Scan(&w.Id, &w.Name, &w.Phone, &w.Address, &w.Age, &w.PassportType, &w.PassportUrl, &w.CartNumber, &w.Rating, &w.ActiveToday, &w.Active, &w.Description, &w.GroupId, &w.Skill,&w.Ready,&w.TownId)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

func GetWorkerById(id int64) *Worker {

	rows, err := DB.Query("SELECT id,name,phone,address,age,passport_type,passport_url,cart_number,rating,active_today,active,description,group_id,skill,ready,town_id FROM workers WHERE id=$1 LIMIT 1", id)
	utils.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		w := Worker{}
		err = rows.Scan(&w.Id, &w.Name, &w.Phone, &w.Address, &w.Age, &w.PassportType, &w.PassportUrl, &w.CartNumber, &w.Rating, &w.ActiveToday, &w.Active, &w.Description, &w.GroupId, &w.Skill,&w.Ready,&w.TownId)
		utils.CheckErr(err)
		return &w
	}
	return nil
}

func GetWorkerByPhone(p string) *Worker {

	rows, err := DB.Query("SELECT id,name,phone,address,age,passport_type,passport_url,cart_number,rating,active_today,active,description,group_id,skill,ready,town_id FROM workers WHERE phone=$1 LIMIT 1", p)
	utils.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		w := Worker{}
		err = rows.Scan(&w.Id, &w.Name, &w.Phone, &w.Address, &w.Age, &w.PassportType, &w.PassportUrl, &w.CartNumber, &w.Rating, &w.ActiveToday, &w.Active, &w.Description, &w.GroupId, &w.Skill,&w.Ready,&w.TownId)
		utils.CheckErr(err)
		return &w
	}
	return nil
}

func GetWorkerByTown(p int64) []Worker {

	rows, err := DB.Query("SELECT id,name,phone,address,age,passport_type,passport_url,cart_number,rating,active_today,active,description,group_id,skill,ready,town_id FROM workers WHERE town_id=$1 AND active=true", p)
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Worker
	for rows.Next() {
		w := Worker{}
		err = rows.Scan(&w.Id, &w.Name, &w.Phone, &w.Address, &w.Age, &w.PassportType, &w.PassportUrl, &w.CartNumber, &w.Rating, &w.ActiveToday, &w.Active, &w.Description, &w.GroupId, &w.Skill,&w.Ready,&w.TownId)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

func GetWorkerByTownAndDate(p int64,date time.Time) []Worker {

	rows, err := DB.Query("SELECT id,name,phone,address,age,passport_type,passport_url,cart_number,rating,active_today,active,description,group_id,skill,ready,town_id FROM workers WHERE town_id=$1", p)
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Worker
	for rows.Next() {
		w := Worker{}
		err = rows.Scan(&w.Id, &w.Name, &w.Phone, &w.Address, &w.Age, &w.PassportType, &w.PassportUrl, &w.CartNumber, &w.Rating, &w.ActiveToday, &w.Active, &w.Description, &w.GroupId, &w.Skill,&w.Ready,&w.TownId)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

func ImportWorkers(ws []Worker) bool {
	//if RemoveAllProduct(){
		txn, err := DB.Begin()
		if err != nil {
			fmt.Println("Error Import ",  err)
			log.Fatal(err)
			return false
		}

		stmt, err := txn.Prepare(pq.CopyIn("workers", "name", "phone", "address", "age", "passport_type","passport_url","cart_number","rating","active_today","active","description","group_id","skill","ready","town_id"))
		if err != nil {
			fmt.Println("Error Import ",  err)
			log.Fatal(err)
			return false
		}

		for _, p := range ws {
			_, err = stmt.Exec(p.Name, p.Phone,p.Address,p.Age,p.PassportType,p.PassportUrl,p.CartNumber,p.Rating,p.ActiveToday,p.Active,p.Description,p.GroupId,p.Skill,p.Ready,p.TownId)
			if err != nil {
				fmt.Println("Error Import ",  err)
				log.Fatal(err)
			}
		}

		_, err = stmt.Exec()
		if err != nil {
			fmt.Println("Error Import ",  err)
			log.Fatal(err)
			return false
		}

		err = stmt.Close()
		if err != nil {
			fmt.Println("Error Import ",  err)
			log.Fatal(err)
			return false
		}

		err = txn.Commit()
		if err != nil {
			fmt.Println("Error Import ",  err)
			log.Fatal(err)
			return false
		}

		return true
//	}

	//return false

}