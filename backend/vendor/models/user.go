package models

import (
	"fmt"
	"utils"
)

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type NewUser struct {
	Username string  `json:"username" validate:"required,gte=3"`
	Password string  `json:"password" validate:"required,gte=6"`
	Name     string  `json:"name" validate:"required,gte=3"`
	Email    string  `json:"email" validate:"required,email"`
	Role     string  `json:"role" validate:"required"`
	Rate     float64 `json:"rate"`
	RateTax     float64 `json:"rate_tax"`
}

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Rate     float64 `json:"rate"`
	RateTax     float64 `json:"rate_tax"`
}

//Check before create user
func GetUserByUserName(username string) *User {

	rows, err := DB.Query("SELECT username FROM users WHERE username=$1 LIMIT 1", username)
	utils.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Username)
		utils.CheckErr(err)
		return &user
	}

	return nil
}

func GetUserAuthByUserName(username string) *UserAuth {

	rows, err := DB.Query("SELECT username,password,role FROM users WHERE username=$1 LIMIT 1", username)
	utils.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		user := UserAuth{}
		err = rows.Scan(&user.Username, &user.Password, &user.Role)
		utils.CheckErr(err)
		return &user
	}

	return nil
}

func GetUserIdByUserName(username string) int64 {

	rows, err := DB.Query("SELECT id FROM users WHERE username=$1 LIMIT 1", username)
	utils.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var userId int64
		err = rows.Scan(&userId)
		utils.CheckErr(err)
		return userId
	}

	return 0
}

func CreateUser(u NewUser) string {

	var uid string
	err := DB.QueryRow("INSERT INTO users(username,password,name,email,role,rate,rate_tax) VALUES($1,$2,$3,$4,$5,$6,$7) returning username", u.Username, u.Password, u.Name, u.Email, u.Role,u.Rate,u.RateTax).Scan(&uid)
	utils.CheckErr(err)

	return uid
}

func UpdateUser(userAuth NewUser) bool {

	stmt, err := DB.Prepare("update users set (password)=($1) where username=$2")
	utils.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(userAuth.Password, userAuth.Username)
	utils.CheckErr(err)

	affect, err := res.RowsAffected()
	utils.CheckErr(err)

	fmt.Println(affect, "rows changed")
	if affect > 0 {
		return true
	}
	return false
}

func UpdateUserFull(user NewUser) bool {
	stmt, err := DB.Prepare("update users set (name, email, role, rate, rate_tax)=($2,$3,$4,$5,$6) where username=$1")
	utils.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(user.Username, user.Name, user.Email, user.Role, user.Rate, user.RateTax)
	utils.CheckErr(err)

	affect, err := res.RowsAffected()
	utils.CheckErr(err)

	fmt.Println(affect, "rows changed")
	if affect > 0 {
		return true
	}
	return false
}

func DeleteUser(user User) bool {

	stmt, err := DB.Prepare("delete from users where username=$1")
	utils.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(user.Username)
	utils.CheckErr(err)

	affect, err := res.RowsAffected()
	utils.CheckErr(err)

	fmt.Println(affect, "rows changed")

	if affect > 0 {
		return true
	}
	return false

}

func GetUsers() []User {

	rows, err := DB.Query("SELECT id,username,name,email,role,rate,rate_tax FROM users LIMIT 100")
	utils.CheckErr(err)
	defer rows.Close()

	var users []User
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Role,&user.Rate,&user.RateTax)
		utils.CheckErr(err)

		users = append(users, user)
	}

	return users
}

func GetUserById(id int64) *User {

	rows, err := DB.Query("SELECT id,username,name,email,role,rate,rate_tax FROM users WHERE id=$1 LIMIT 1", id)
	utils.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Role,&user.Rate,&user.RateTax)
		utils.CheckErr(err)

		return &user
	}

	return nil
}

func GetUserByUsername(username string) *User {

	rows, err := DB.Query("SELECT id,username,name,email,role,rate,rate_tax FROM users WHERE username=$1 LIMIT 1", username)
	utils.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Role,&user.Rate, &user.RateTax)
		utils.CheckErr(err)

		return &user
	}

	return nil
}

func GetUsersByRole(role string) []User {
	rows, err := DB.Query("SELECT id,username,name,email,role,rate,rate_tax FROM users WHERE role=$1 LIMIT 100", role)
	utils.CheckErr(err)
	defer rows.Close()

	var users []User
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Role,&user.Rate, &user.RateTax)
		utils.CheckErr(err)

		users = append(users, user)
	}

	return users
}

type BufferUser struct{
	Username string  `json:"username" validate:"required,gte=3"`
	Password string  `json:"password" validate:"required,gte=6"`
	Name     string  `json:"name" validate:"required,gte=3"`
	Email    string  `json:"email" validate:"required,email"`
	Role     string  `json:"role" validate:"required"`
	Rate     string `json:"rate"`
	RateTax  string `json:"rate_tax"`
}

func ToNewUser(buf BufferUser) NewUser {
	var user NewUser
	user.Username = buf.Username
	user.Password = buf.Password
	user.Name = buf.Name
	user.Email = buf.Email
	user.Role = buf.Role
	user.Rate = utils.StringToFloat(buf.Rate)
	user.RateTax = utils.StringToFloat(buf.RateTax)
	return user
}
