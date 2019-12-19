package models

import (
	"fmt"
	"utils"
)

type Customer struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Email		string `json:"email"`
	Company_name string `json:"company_name"`
	Active      bool   `json:"active"`
	TownId      int64  `json:"town_id"`
	UrFlag      bool   `json:"urflag"`
	Rate      float64   `json:"rate"`
	Sale      float64   `json:"sale"`
	Referee   string    `json:"referee"`
}

func CreateCustomer(w Customer) int64 {

	var uid int64
	err := DB.QueryRow("INSERT INTO customers(name,phone,address,email,company_name,active,description,town_id,urflag,rate,sale,referee)"+
		" VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) returning id",
		w.Name, w.Phone, w.Address, w.Email, w.Company_name, w.Active, w.Description, w.TownId, w.UrFlag,w.Rate,w.Sale, w.Referee).Scan(&uid)
	utils.CheckErr(err)
	return uid
}

func UpdateCustomer(w Customer) bool {

	stmt, err := DB.Prepare("update customers set (name,phone,address,email,company_name,active,description,town_id,urflag,rate,sale,referee)=($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) where id=$13")
	utils.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(w.Name, w.Phone, w.Address, w.Email, w.Company_name, w.Active, w.Description, w.TownId, w.UrFlag,w.Rate,w.Sale, w.Id)
	utils.CheckErr(err)

	affect, err := res.RowsAffected()
	utils.CheckErr(err)

	fmt.Println(affect, "rows changed")
	if affect > 0 {
		return true
	}
	return false
}

func DeleteCustomer(w Customer) bool {

	stmt, err := DB.Prepare("delete from customers where id=$1")
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

func GetCustomers() []Customer {

	rows, err := DB.Query("SELECT id,name,phone,address,email,company_name,active,description,town_id,urflag,rate,sale,referee FROM customers LIMIT 1000")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Customer
	for rows.Next() {
		w := Customer{}
		err = rows.Scan(&w.Id, &w.Name, &w.Phone, &w.Address, &w.Email, &w.Company_name, &w.Active, &w.Description, &w.TownId, &w.UrFlag,&w.Rate,&w.Sale, &w.Referee)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

func GetCustomerById(id int64) *Customer {

	rows, err := DB.Query("SELECT id,name,phone,address,email,company_name,active,description,town_id,urflag,rate,sale,referee FROM customers WHERE id=$1 LIMIT 1", id)
	utils.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		w := Customer{}
		err = rows.Scan(&w.Id, &w.Name, &w.Phone, &w.Address, &w.Email, &w.Company_name, &w.Active, &w.Description, &w.TownId, &w.UrFlag,&w.Rate,&w.Sale,&w.Referee)
		utils.CheckErr(err)
		return &w
	}
	return nil
}

func GetCustomerByPhone(p string) *Customer {
	rows, err := DB.Query("SELECT id,name,phone,address,email,company_name,active,description,town_id,urflag,rate,sale,referee FROM customers WHERE phone=$1 LIMIT 1", p)

	utils.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		w := Customer{}
		err = rows.Scan(&w.Id, &w.Name, &w.Phone, &w.Address, &w.Email, &w.Company_name, &w.Active, &w.Description, &w.TownId, &w.UrFlag,&w.Rate,&w.Sale,&w.Referee)
		utils.CheckErr(err)
		return &w
	}
	return nil
}

func GetCustomerByTown(p int64) []Customer {

	rows, err := DB.Query("SELECT id,name,phone,address,email,company_name,active,description,town_id,urflag,rate,sale,referee FROM customers WHERE town_id=$1", p)
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Customer
	for rows.Next() {
		w := Customer{}
		err = rows.Scan(&w.Id, &w.Name, &w.Phone, &w.Address, &w.Email, &w.Company_name, &w.Active, &w.Description, &w.TownId, &w.UrFlag,&w.Rate,&w.Sale, &w.Referee)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}
