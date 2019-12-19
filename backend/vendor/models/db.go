package models

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"

	"utils"
	"fmt"
)

var DB *sql.DB

func InitDB(port string) bool {
	connStr := "host=localhost port="+port+" user=postgres password=postgres dbname=pms sslmode=disable fallback_application_name=go_pms"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
		return false
	}

	DB = db

	return true
}

func Migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS users(
        id BIGSERIAL PRIMARY KEY NOT NULL,
        username VARCHAR(255) NOT NULL,
        password VARCHAR(255) NOT NULL,
        name VARCHAR(255),
        email VARCHAR(255),
        role VARCHAR(255),
        rate REAL DEFAULT 0.0,
        rate_tax REAL DEFAULT 0.0
    );
  	CREATE TABLE IF NOT EXISTS worker_group(
        id BIGSERIAL PRIMARY KEY NOT NULL,
      	name VARCHAR(255)
  	);
	CREATE TABLE IF NOT EXISTS workers(
        id BIGSERIAL PRIMARY KEY NOT NULL,
        name VARCHAR(255),
        phone VARCHAR(255) NOT NULL,
        address VARCHAR(255),
        age VARCHAR(255),
        passport_type VARCHAR(255),
        passport_url VARCHAR(255),
        cart_number VARCHAR(255),
        rating REAL DEFAULT 0.0,
        active_today BOOLEAN DEFAULT TRUE,
        active BOOLEAN DEFAULT TRUE,
        description VARCHAR(255),
        group_id BIGSERIAL,
        skill VARCHAR(255) DEFAULT '',
        ready VARCHAR(255) DEFAULT '',
        town_id BIGINT DEFAULT 0
    );
	CREATE TABLE IF NOT EXISTS tasks(
        id BIGSERIAL PRIMARY KEY NOT NULL,
        name VARCHAR(255),
        start_address VARCHAR(255),
        end_address VARCHAR(255),
        start_date timestamp with time zone,
        end_date timestamp with time zone,
        payment_type BOOLEAN,
        payment_first REAL,
        payment_second REAL,
        workers_count INTEGER,
        total_hours REAL,
        total_price REAL,
        total_worker_price REAL,
        total_delta REAL,
        additional_cost REAL,
        additional_income REAL,
        description VARCHAR(255),
        cargo_description VARCHAR(255),
        create_date timestamp with time zone,
        update_date timestamp with time zone,
        status_id BIGSERIAL,
        payment_first_count INTEGER,
        payment_second_count INTEGER,
        customer_id BIGSERIAL,
        author BIGSERIAL,
        executor BIGSERIAL,
        paid BOOLEAN DEFAULT FALSE,
        contact VARCHAR(255),
        declarated_cost REAL,
        town_id BIGINT DEFAULT 0,
        total_manager REAl
        
    );
    CREATE TABLE IF NOT EXISTS towns(
        id BIGSERIAL PRIMARY KEY NOT NULL,
        name VARCHAR(255)
    );
    CREATE TABLE IF NOT EXISTS taskstatus(
        id BIGSERIAL PRIMARY KEY NOT NULL,
        name VARCHAR(255),
        description VARCHAR(255)
    );
    CREATE TABLE IF NOT EXISTS customers(
        id BIGSERIAL PRIMARY KEY NOT NULL,
        name VARCHAR(255),
        phone VARCHAR(255) NOT NULL,
        address VARCHAR(255),
        email VARCHAR(255),
        company_name VARCHAR (255),
        active BOOLEAN DEFAULT TRUE,
        description VARCHAR(255),
        urflag BOOLEAN DEFAULT FALSE,
        town_id BIGSERIAL,
        rate REAL,
        sale REAL,
        referee VARCHAR(255)
    );
    CREATE TABLE IF NOT EXISTS task_workers(
        id BIGSERIAL PRIMARY KEY NOT NULL,
        task_id BIGSERIAL REFERENCES tasks ON DELETE CASCADE,
        worker_id BIGSERIAL REFERENCES workers ON DELETE RESTRICT,
        price REAL,
        road_price REAL,
        start_date timestamp with time zone,
        end_date timestamp with time zone,
        total_price REAL,
        description VARCHAR(255),
        dinner REAL DEFAULT 0,
        prepaid REAL DEFAULT 0,
        paid BOOLEAN DEFAULT FALSE
    );
    `

	_, err := db.Exec(sql)

	//var err error
	//for i := 0; i < 10; i++ {
	//	_, err = db.Exec(sql)
	//
	//	if err == nil {
	//		break
	//	}
	//	time.Sleep(5 * time.Second)
	//}

	if err != nil{
		utils.CheckErr(err)
		return
	}

	//migrateTotalManager(db)
	//migrateUserRateTax(db)
	//migrateCustomerRateAndSale(db)

	//TODO: Миграция не добавлена на прод


	fmt.Println("Migrations successfull ended")

}

//Common Methods
func GetTotal(entity string) int {
	var count int
	sql := "SELECT COUNT(id) FROM " + entity
	row := DB.QueryRow(sql)
	err := row.Scan(&count)
	utils.CheckErr(err)

	return count
}

func migrateFields(db *sql.DB){

	sql := `
	ALTER TABLE tasks ADD COLUMN customer_id REAL;
	`

	//sql := `
  	//ALTER TABLE tasks ADD CONSTRAINT task_status_fkey FOREIGN KEY (status_id) REFERENCES taskstatus ON DELETE RESTRICT;
  	//ALTER TABLE tasks ADD CONSTRAINT task_customer_fkey FOREIGN KEY (customer_id) REFERENCES customers ON DELETE RESTRICT;
    //`

	//sql := `
  	//ALTER TABLE workers DROP COLUMN town_id;
    //`
	_, err := db.Exec(sql)

	if err != nil{
		utils.CheckErr(err)
		return
	}

	fmt.Println("Migration migrateFileds successfull ended")

}



func migrateTotalManager(db *sql.DB)  {
	sql := `
	ALTER TABLE tasks ADD COLUMN total_manager REAL DEFAULT 0;
	`

	_, err := db.Exec(sql)
	if err != nil{
		fmt.Println("Migration migrateTotalManager canceled with error")
		utils.CheckErr(err)
		return
	}
	fmt.Println("Migration migrateTotalManager successfull")
}

func migrateUserRateTax(db *sql.DB)  {
	sql := `
	ALTER TABLE users ADD COLUMN rate_tax REAL DEFAULT 0.0;
	`

	_, err := db.Exec(sql)
	if err != nil{
		fmt.Println("Migration migrateUserRateTax canceled with error")
		utils.CheckErr(err)
		return
	}
	fmt.Println("Migration migrateUserRateTax successfull")
}

func migrateCustomerRateAndSale(db *sql.DB)  {
	sql := `
	ALTER TABLE customers ADD COLUMN rate REAL DEFAULT 0.0;
	ALTER TABLE customers ADD COLUMN sale REAL DEFAULT 0.0;
	`

	_, err := db.Exec(sql)
	if err != nil{
		fmt.Println("Migration migrateCustomerRateAndSale canceled with error")
		utils.CheckErr(err)
		return
	}
	fmt.Println("Migration migrateCustomerRateAndSale successfull")
}