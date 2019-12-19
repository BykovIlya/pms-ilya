package main

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"fmt"
	"log"
	"models"
	"os"
	"path/filepath"
	"routes"

	"encoding/json"
	"roles"
	"time"
)

//TODO: Warning Dev or Prod Mode
var isDev = false

var isBuild = "dev"

type Dev struct {
	DBPort string `json:"port"`
}
type Prod struct {
	DBPort string `json:"port"`
}
type Configuration struct {
	Dev  `json:"dev"`
	Prod `json:"prod"`
}

var identityKey = "id"

func main() {

	if isBuild == "dev" {
		isDev = true
		fmt.Println("Start Developing Mode")
	} else if isBuild == "prod" {
		isDev = false
		fmt.Println("Start Prodaction Mode")
	}

	if models.InitDB(ParceConfig(isDev)) {
		fmt.Println("Database successfull initialisation")
		models.Migrate(models.DB)
		createFirstUsers()
		createCatalogs()
		defer models.DB.Close()
	} else {
		log.Panic("Database not initialisation")
	}

	//Create empty dirs for Files
	CreateDirsForFiles()

	//Temporary Import
	//routes.TemporaryImportExcel()

	//Init Roles and Permissions
	if roles.Init() {
		fmt.Println("Roles and Permissions successfull initialisation")
	} else {
		log.Panic("Roles and Permissions not initialisation")
	}

	router := gin.Default()
	router.Use(CORSMiddleware())
	router.Static("/api/upload", "./api/upload")
	router.Static("/api/tmp", "./api/tmp")

	//The jwt middleware
	userAuthMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "PMS secret zone",
		Key:             []byte("PMS secret key"),
		Timeout:         24 * 720 * time.Hour,
		MaxRefresh:      24 * 720 * time.Hour,
		Authenticator:   routes.AuthenticatorUser,
		Authorizator:    routes.AuthorizatorUser,
		Unauthorized:    routes.UnauthorizedUser,
		TokenLookup:     "header:Authorization",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		IdentityKey:     identityKey,
		PayloadFunc:     payloadFunc,
		IdentityHandler: identityHandler,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	//API
	api := router.Group("/api")
	api.POST("/login", userAuthMiddleware.LoginHandler)

	api.Use(userAuthMiddleware.MiddlewareFunc())
	{
		api.GET("users/info", routes.GetUserInfo)
		api.GET("users/id/:id", routes.GetUserById)
		api.GET("statistics/manager", routes.GetTasksStatisticsForManagers)
		api.GET("statistics/managersNames", routes.GetManagers)
		managers := api.Group("/managers")
		{
			managers.POST("", routes.CreateUser)
			managers.GET("/id/:id", routes.GetUserById)
			managers.PUT("/:id", routes.UpdateUser)
			managers.DELETE("/:id", routes.DeleteUser)
			managers.GET("/managersNames", routes.GetUsers)
		}
		status := api.Group("/taskstatus")
		{
			status.POST("", routes.CreateTaskStatus)
			status.GET("/id/:id", routes.ReadTaskStatusById)
			status.PUT("/:id", routes.UpdateTaskStatus)
			status.DELETE("/:id", routes.DeleteTaskStatus)
			status.GET("", routes.GetTaskStatuses)
		}

		town := api.Group("town/:town")
		{
			workers := town.Group("/workers")
			{
				workers.POST("", routes.CreateWorker)
				workers.GET("/id/:id", routes.ReadWorkerById)
				workers.PUT("/:id", routes.UpdateWorker)
				workers.DELETE("/:id", routes.UnActiveWorker)
				workers.GET("", routes.GetWorkers)
				workers.POST("/import/:id", routes.ImportWorkerFile)
				workers.GET("/date/:date/end_date/:end_date", routes.GetWorkersWithDate)

			}

			tasks := town.Group("/tasks")
			{
				tasks.POST("", routes.CreateTask)
				tasks.GET("/id/:id", routes.ReadTaskById)
				tasks.PUT("/:id", routes.UpdateTask)
				tasks.DELETE("/:id", routes.DeleteTask)
				tasks.GET("", routes.GetTasks)
			}

			customers := town.Group("/customers")
			{
				customers.POST("", routes.CreateCustomer)
				customers.GET("/id/:id", routes.ReadCustomerById)
				customers.PUT("/:id", routes.UpdateCustomer)
				customers.DELETE("/:id", routes.DeleteCustomer)
				customers.GET("", routes.GetCustomers)
			}

			payments := town.Group("/payments")
			{
				payments.GET("/detail", routes.GetDetailPayments)
				payments.GET("", routes.GetPayments)
				payments.POST("/task/:id", routes.UpdatePaidTaskWorker)
				payments.POST("/all", routes.UpdateAllPaidTaskWorker)
			}
		}
	}

	router.Run(":5000")
}

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(string); ok {
		return jwt.MapClaims{
			identityKey: v,
		}
	}
	return jwt.MapClaims{}
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return claims["id"].(string)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length,X-Requested-With, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func CreateDirsForFiles() {
	tmp := filepath.Join(".", "api/tmp")
	os.MkdirAll(tmp, os.ModePerm)

	upload := filepath.Join(".", "api/upload")
	os.MkdirAll(upload, os.ModePerm)
}

func ParceConfig(isDev bool) string {

	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration.Dev.DBPort)
	fmt.Println(configuration.Prod.DBPort)
	if isDev {
		return configuration.Dev.DBPort
	}
	return configuration.Prod.DBPort
}

func createFirstUsers() {
	admin := models.GetUserByUserName("admin")
	if admin == nil {
		newAdmin := models.NewUser{}
		newAdmin.Username = "admin"
		newAdmin.Password = routes.CryptUserPassword("admin")
		newAdmin.Role = "admin"
		newAdmin.Name = "Администратор"
		newAdmin.Email = "admin@a.a"
		newAdmin.Rate = 30
		uuid := models.CreateUser(newAdmin)
		if uuid == "admin" {
			fmt.Println("First admin Created")
		}
	}

	manager := models.GetUserByUserName("manager")
	if manager == nil {
		newAdmin := models.NewUser{}
		newAdmin.Username = "manager"
		newAdmin.Password = routes.CryptUserPassword("manager")
		newAdmin.Role = "manager"
		newAdmin.Name = "Менеджер"
		newAdmin.Email = "manager@a.a"
		newAdmin.Rate = 30
		uuid := models.CreateUser(newAdmin)
		if uuid == "manager" {
			fmt.Println("First manager Created")
		}
	}

	accountant := models.GetUserByUserName("acc")
	if accountant == nil {
		newAdmin := models.NewUser{}
		newAdmin.Username = "acc"
		newAdmin.Password = routes.CryptUserPassword("acc")
		newAdmin.Role = "accountant"
		newAdmin.Name = "Бухгалтер"
		newAdmin.Email = "acc@a.a"
		newAdmin.Rate = 30
		uuid := models.CreateUser(newAdmin)
		if uuid == "acc" {
			fmt.Println("First accountant Created")
		}
	}

	manager1 := models.GetUserByUserName("ekarkishchenko")
	if manager1 == nil {
		newAdmin := models.NewUser{}
		newAdmin.Username = "ekarkishchenko"
		newAdmin.Password = routes.CryptUserPassword("0a5a2h")
		newAdmin.Role = "manager"
		newAdmin.Name = "Екатерина Каркищенко"
		newAdmin.Email = "9675361511r@mail.ru"
		newAdmin.Rate = 30
		uuid := models.CreateUser(newAdmin)
		if uuid == "ekarkishchenko" {
			fmt.Println("ekarkishchenko Created")
		}
	}

	manager2 := models.GetUserByUserName("adautova")
	if manager2 == nil {
		newAdmin := models.NewUser{}
		newAdmin.Username = "adautova"
		newAdmin.Password = routes.CryptUserPassword("6w0ttk")
		newAdmin.Role = "manager"
		newAdmin.Name = "Айгуль Даутова"
		newAdmin.Email = "89633077572@mail.ru"
		newAdmin.Rate = 30
		uuid := models.CreateUser(newAdmin)
		if uuid == "adautova" {
			fmt.Println("adautova Created")
		}
	}

	manager3 := models.GetUserByUserName("ubarboshov")
	if manager3 == nil {
		newAdmin := models.NewUser{}
		newAdmin.Username = "ubarboshov"
		newAdmin.Password = routes.CryptUserPassword("tg3dm6")
		newAdmin.Role = "manager"
		newAdmin.Name = "Барбошов Юрий"
		newAdmin.Email = "ubarboshov@pro-gruzchiki.ru"
		newAdmin.Rate = 30
		uuid := models.CreateUser(newAdmin)
		if uuid == "ubarboshov" {
			fmt.Println("ubarboshov Created")
		}
	}

	manager4 := models.GetUserByUserName("ityrancev")
	if manager4 == nil {
		newAdmin := models.NewUser{}
		newAdmin.Username = "ityrancev"
		newAdmin.Password = routes.CryptUserPassword("rx73fv")
		newAdmin.Role = "manager"
		newAdmin.Name = "Туранцев Иван"
		newAdmin.Email = "malyarenko.70@list.ru"
		newAdmin.Rate = 20
		uuid := models.CreateUser(newAdmin)
		if uuid == "ityrancev" {
			fmt.Println("ityrancev Created")
		}
	}
}

func createCatalogs() {

	var n1 = "Новая"
	var d1 = "Не назначены грузчики"
	nStatus(n1, d1)
	var n2 = "Готова к работе"
	var d2 = "Грузчики назначены"
	nStatus(n2, d2)
	var n3 = "В работе"
	var d3 = "Грузчики на месте и начали работу"
	nStatus(n3, d3)
	var n4 = "Завешена"
	var d4 = "Оплачена грузчикам"
	nStatus(n4, d4)
	var n5 = "Ожидает оплаты заказчиком"
	var d5 = "Ожидает оплаты заказчиком"
	nStatus(n5, d5)
	var n6 = "Закрыта"
	var d6 = "Оплачена заказчиком"
	nStatus(n6, d6)
	var n7 = "Отменена"
	var d7 = "Указать причину отмены"
	nStatus(n7, d7)
}

func nStatus(n string, d string) {
	s1 := models.GetTaskStatusByName(n)
	if s1 == nil {
		s1 := models.TaskStatus{}
		s1.Name = n
		s1.Description = d
		uuid := models.CreateTaskStatus(s1)
		if uuid > 0 {
			fmt.Println("Task Status "+n+" Создан id = ", uuid)
		}
	}
}
